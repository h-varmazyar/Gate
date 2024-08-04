package posts

import (
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/app/posts/workers"
	"github.com/h-varmazyar/Gate/services/chipmunk/internal/pkg/db"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
)

type App struct {
}

type AppDependencies struct {
}

func NewApp(ctx context.Context, logger *log.Logger, configs Configs, db *db.DB, _ *AppDependencies) (*App, error) {
	logger.Infof("initiating posts app")
	collectorWorker := workers.NewCollector(ctx, logger, configs.WorkersConfigs, db)
	collectorWorker.Start(ctx)

	//workers.NewSentimentFiller(logger, configs.WorkersConfigs, db).Start(ctx)

	return &App{}, nil
}
