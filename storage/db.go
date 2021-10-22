package storage

import (
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
)

var dbQueue chan models.Candle

func init() {
	dbQueue = make(chan models.Candle, 10000)
	go func() {
		for candle := range dbQueue {
			if err := candle.CreateOrUpdate(); err != nil {
				log.WithError(err).Error("saving new candle failed")
			}
		}
	}()
}
