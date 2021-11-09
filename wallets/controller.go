package wallets

import (
	"errors"
	"fmt"
	"github.com/mrNobody95/Gate/brokerages"
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type Controller struct {
	wallets  map[string]*models.Wallet
	Requests brokerages.BrokerageRequests
}

func NewController(requests brokerages.BrokerageRequests) *Controller {
	c := new(Controller)
	c.wallets = make(map[string]*models.Wallet)
	c.Requests = requests

	c.checkWalletsStatus()
	return c
}

func (c *Controller) checkWalletsStatus() {
	go func() {
		for {
			response := c.Requests.WalletList()
			if response.Error != nil {
				log.WithError(response.Error).Error("wallet list status failed")
			} else {
				for _, wallet := range response.Wallets {
					c.wallets[wallet.Asset.Symbol] = &wallet
				}
			}
			time.Sleep(time.Minute)
		}
	}()
}

func (c *Controller) GetWallet(symbol string) (*models.Wallet, error) {
	wallet, ok := c.wallets[symbol]
	if !ok {
		return nil, errors.New("wallet not found")
	}
	return wallet, nil
}

func (c *Controller) UsdtValue() float64 {
	totalUsdt := float64(0)
	for _, wallet := range c.wallets {
		resp := c.Requests.MarketStatistics(brokerages.MarketInfoParams{MarketName: fmt.Sprintf("%sUSDT", strings.ToUpper(wallet.Asset.Symbol))})
		if resp.Error != nil {
			continue
		}
		totalUsdt += resp.Candle.Close * wallet.TotalBalance
	}
	return totalUsdt
}
