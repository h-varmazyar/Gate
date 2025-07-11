package marketUpdate

import (
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-co-op/gocron/v2"
	chipmunkApi "github.com/h-varmazyar/Gate/services/chipmunk/api/proto"
	coreApi "github.com/h-varmazyar/Gate/services/core/api/proto"
	"github.com/h-varmazyar/Gate/services/gather/configs"
	"github.com/h-varmazyar/Gate/services/gather/internal/domain"
	"github.com/h-varmazyar/Gate/services/gather/internal/models"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/candleTicker"
	"github.com/h-varmazyar/Gate/services/gather/internal/workers/lastCandle"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"os"
	"path/filepath"
	"time"
)

type coreAdapter interface {
	MarketList(ctx context.Context) ([]*chipmunkApi.Market, error)
	MarketInfo(ctx context.Context, marketKey string) (*coreApi.MarketInfo, error)
}

type coinexAdapter interface {
	MarketList(ctx context.Context) (domain.CoinexMarkets, error)
}

type marketsRepo interface {
	All(ctx context.Context) ([]models.Market, error)
	Delete(ctx context.Context, marketID uint) error
	MarketCount(ctx context.Context, assetID uint) (int64, error)
	Create(ctx context.Context, market models.Market) (models.Market, error)
}

type candlesRepo interface {
	AllMarketCandles(ctx context.Context, marketID uint, offset int) ([]models.Candle, error)
	DeleteMarketCandles(ctx context.Context, marketID uint) error
}

type assetsRepo interface {
	ReturnBySymbol(ctx context.Context, symbol string) (*models.Asset, error)
	Create(ctx context.Context, asset *models.Asset) (*models.Asset, error)
	Delete(ctx context.Context, assetID uint) error
}

type resolutionsRepo interface {
	All(ctx context.Context) ([]models.Resolution, error)
	ReturnByID(ctx context.Context, id uint) (models.Resolution, error)
}

type Worker struct {
	logger             *log.Logger
	cfg                configs.WorkerMarketUpdate
	assetsRepo         assetsRepo
	candlesRepo        candlesRepo
	marketsRepo        marketsRepo
	resolutionsRepo    resolutionsRepo
	coreAdapter        coreAdapter
	coinexAdapter      coinexAdapter
	job                gocron.Job
	candleTickerWorker *candleTicker.Worker
	lastCandleWorker   *lastCandle.Worker
}

func NewWorker(
	logger *log.Logger,
	cfg configs.WorkerMarketUpdate,
	assetsRepo assetsRepo,
	candlesRepo candlesRepo,
	marketsRepo marketsRepo,
	resolutionsRepo resolutionsRepo,
	coreAdapter coreAdapter,
	coinexAdapter coinexAdapter,
	candleTickerWorker *candleTicker.Worker,
	lastCandleWorker *lastCandle.Worker,
) *Worker {
	return &Worker{
		logger:             logger,
		cfg:                cfg,
		assetsRepo:         assetsRepo,
		candlesRepo:        candlesRepo,
		marketsRepo:        marketsRepo,
		resolutionsRepo:    resolutionsRepo,
		coreAdapter:        coreAdapter,
		coinexAdapter:      coinexAdapter,
		candleTickerWorker: candleTickerWorker,
		lastCandleWorker:   lastCandleWorker,
	}
}

func (w *Worker) Start(s gocron.Scheduler) error {
	if !w.cfg.Running {
		return nil
	}
	if w.cfg.NeedWarmup {
		w.immediateUpdate()
	}
	j, err := s.NewJob(
		gocron.CronJob(w.cfg.RunningTime, false),
		gocron.NewTask(w.update),
	)
	if err != nil {
		return err
	}

	w.job = j

	return nil
}

func (w *Worker) immediateUpdate() {
	w.update()
}

func (w *Worker) update() {
	w.logger.Info("updating markets")
	ctx := context.Background()
	localMarkets, err := w.marketsRepo.All(ctx)
	if err != nil {
		w.logger.WithError(err).Errorf("failed to fetch markets")
		return
	}

	w.logger.Infof("local market len: %v", len(localMarkets))

	remoteMarkets := make([]models.Market, 0)
	markets, err := w.coinexAdapter.MarketList(ctx)
	if err != nil {
		w.logger.WithError(err).Errorf("failed to list remote markets")
		return
	}

	for _, market := range markets.List {
		localMarket := models.Market{
			IsPremarketTradingAvailable: market.IsPremarketTradingAvailable,
			IsAmmAvailable:              market.IsAmmAvailable,
			IsMarginAvailable:           market.IsMarginAvailable,

			BaseCurrencyPrecision:  market.BaseCurrencyPrecision,
			QuoteCurrencyPrecision: market.QuoteCurrencyPrecision,
			TakerFeeRate:           market.TakerFeeRate,
			MakerFeeRate:           market.MakerFeeRate,
			QuoteCurrency: &models.Asset{
				Symbol: market.QuoteCurrency,
			},
			MinAmount: market.MinAmount,
			BaseCurrency: &models.Asset{
				Symbol: market.BaseCurrency,
			},
			Name:      market.Market,
			IssueDate: market.IssueDate,
			Status:    models.MarketStatusEnable,
		}
		remoteMarkets = append(remoteMarkets, localMarket)
	}

	w.logger.Infof("remote market len: %v", len(remoteMarkets))

	w.deleteRemovableMarkets(ctx, localMarkets, remoteMarkets)
	w.saveNewMarkets(ctx, localMarkets, remoteMarkets)
}

// loop over markets and in each iteration:
// 1: create new asset if not available
// 2: create market in database
// 3: attach market to data collector
func (w *Worker) saveNewMarkets(ctx context.Context, localMarkets, remoteMarkets []models.Market) {
	dbMarketNames := make(map[string]bool)
	for _, market := range localMarkets {
		dbMarketNames[market.Name] = true
	}

	var newMarkets []models.Market
	for _, market := range remoteMarkets {
		if !dbMarketNames[market.Name] {
			newMarkets = append(newMarkets, market)
		}
	}

	w.logger.Info("new markets len: ", len(newMarkets))

	for _, market := range newMarkets {
		time.Sleep(time.Second)
		var err error
		market.BaseCurrency, err = w.createAsset(ctx, market.BaseCurrency.Name, market.BaseCurrency.Symbol)
		if err != nil {
			w.logger.WithError(err).Errorf("failed to save source %v", market.BaseCurrency.Symbol)
			continue
		}

		market.QuoteCurrency, err = w.createAsset(ctx, market.QuoteCurrency.Name, market.QuoteCurrency.Symbol)
		if err != nil {
			w.logger.WithError(err).Errorf("failed to save destination %v", market.QuoteCurrency.Symbol)
			continue
		}

		market.Status = models.MarketStatusEnable

		//marketInfo, err := w.coreAdapter.MarketInfo(ctx, market.Source.Name)
		//if err != nil {
		//	w.logger.WithError(err).Errorf("failed to get market info for %v", market.Name)
		//	continue
		//}
		//market.IssueDate = time.Unix(marketInfo.IssueDate, 0)

		newMarket, err := w.marketsRepo.Create(ctx, market)
		if err != nil {
			w.logger.WithError(err).Errorf("failed to create market %v", market.Name)
			continue
		}

		if err = w.attachMarket(ctx, newMarket); err != nil {
			w.logger.WithError(err).Errorf("failed to attach market: %v", market.ID)
		}
	}
}

func (w *Worker) createAsset(ctx context.Context, name, symbol string) (*models.Asset, error) {
	asset, err := w.assetsRepo.ReturnBySymbol(ctx, symbol)
	if err == nil {
		return asset, nil
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return w.assetsRepo.Create(ctx, &models.Asset{
			Name:   name,
			Symbol: symbol,
		})
	}
	return nil, err
}

func (w *Worker) attachMarket(ctx context.Context, market models.Market) error {
	if err := w.lastCandleWorker.AttachMarket(ctx, market); err != nil {
		return err
	}
	if err := w.candleTickerWorker.AttachMarket(ctx, market); err != nil {
		return err
	}
	return nil
}

// loop over markets and in each iteration:
// 1: detach removable market from data collector
// 2: archive market data(include candles)
// 3: remove market from database
// 4: remove market assets if no other market runs on those assets
// 5: remove all market candles
func (w *Worker) deleteRemovableMarkets(ctx context.Context, localMarkets, remoteMarkets []models.Market) {
	serverMarketNames := make(map[string]bool)
	for _, market := range remoteMarkets {
		serverMarketNames[market.Name] = true
	}

	var deletedMarkets []models.Market
	for _, market := range localMarkets {
		if !serverMarketNames[market.Name] {
			deletedMarkets = append(deletedMarkets, market)
		}
	}

	w.logger.Infof("deleting removable markets: %v", len(deletedMarkets))

	for _, market := range deletedMarkets {
		if err := w.detachRemovableMarket(ctx, market); err != nil {
			continue
		}

		if err := w.archiveMarketData(ctx, market); err != nil {
			continue
		}

		if err := w.marketsRepo.Delete(ctx, market.ID); err != nil {
			continue
		}

		if err := w.removeAssets(ctx, market); err != nil {
			continue
		}

		err := w.candlesRepo.DeleteMarketCandles(ctx, market.ID)
		if err != nil {
			continue
		}
	}
}

func (w *Worker) prepareCandlesMap(ctx context.Context, marketID uint) (map[uint][]models.Candle, error) {
	candlesMap := make(map[uint][]models.Candle)
	offset := 0

	resolutions, err := w.resolutionsRepo.All(ctx)
	if err != nil {
		return nil, err
	}

	for _, resolution := range resolutions {
		candlesMap[resolution.ID] = make([]models.Candle, 0)
	}

	for {
		tmp, err := w.candlesRepo.AllMarketCandles(ctx, marketID, offset)
		if err != nil {
			return nil, err
		}
		if len(tmp) == 0 {
			break
		}
		offset += len(tmp)
		for _, candle := range tmp {
			candlesMap[candle.ResolutionID] = append(candlesMap[candle.ResolutionID], candle)
		}
	}

	return candlesMap, nil
}

func (w *Worker) detachRemovableMarket(ctx context.Context, market models.Market) error {
	if err := w.lastCandleWorker.DetachMarket(ctx, market); err != nil {
		return err
	}
	if err := w.candleTickerWorker.DetachMarket(ctx, market); err != nil {
		return err
	}
	return nil
}

func (w *Worker) removeAssets(ctx context.Context, market models.Market) error {
	//check source asset
	{
		count, err := w.marketsRepo.MarketCount(ctx, market.BaseCurrencyID)
		if err != nil {
			return err
		}
		if count == 0 {
			return w.assetsRepo.Delete(ctx, market.BaseCurrencyID)
		}
	}

	//check destination asset
	{
		count, err := w.marketsRepo.MarketCount(ctx, market.QuoteCurrencyID)
		if err != nil {
			return err
		}
		if count == 0 {
			return w.assetsRepo.Delete(ctx, market.QuoteCurrencyID)
		}
	}
	return nil
}

func (w *Worker) archiveMarketData(ctx context.Context, market models.Market) error {
	candlesMap, err := w.prepareCandlesMap(ctx, market.ID)
	if err != nil {
		return err
	}

	folderPath := filepath.Join("./", market.Name)
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return err
	}

	for resolutionID, candles := range candlesMap {
		resolution, err := w.resolutionsRepo.ReturnByID(ctx, resolutionID)
		if err != nil {
			return err
		}
		filePath := filepath.Join(folderPath, resolution.Label+".csv")
		if err := w.saveCandlesToCSV(filePath, candles); err != nil {
			return err
		}

		key := filepath.Join(market.Name, resolution.Label+".csv")
		if err := w.uploadFileToS3(key, filePath); err != nil {
			return err
		}
	}

	return nil
}

func (w *Worker) uploadFileToS3(key, filePath string) error {
	s3Session, err := session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			w.cfg.S3AccessKey,
			w.cfg.S3SecretKey,
			"",
		),
		Endpoint: &w.cfg.S3Endpoint,
		Region:   &w.cfg.S3Region,
	})
	if err != nil {
		return err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	svc := s3.New(s3Session)
	if _, err = svc.PutObject(&s3.PutObjectInput{
		Bucket: &w.cfg.S3Bucket,
		Key:    aws.String(key),
		Body:   file,
	}); err != nil {
		return err
	}
	return os.Remove(filePath)
}

func (w *Worker) saveCandlesToCSV(fileName string, candles []models.Candle) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	if err := writer.Write([]string{"Open", "High", "Low", "Close", "Volume", "Amount", "Time"}); err != nil {
		return err
	}

	// Write data
	for _, c := range candles {
		row := []string{
			fmt.Sprintf("%f", c.Open),
			fmt.Sprintf("%f", c.High),
			fmt.Sprintf("%f", c.Low),
			fmt.Sprintf("%f", c.Close),
			fmt.Sprintf("%f", c.Volume),
			fmt.Sprintf("%f", c.Amount),
			fmt.Sprintf("%d", c.Time.Unix()),
		}
		if err := writer.Write(row); err != nil {
			return err
		}
	}

	return nil
}
