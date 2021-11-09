package storage

import (
	"github.com/mrNobody95/Gate/models"
	log "github.com/sirupsen/logrus"
)

var DbQueue chan models.Candle

func init() {
	DbQueue = make(chan models.Candle, 100000)
	go func() {
		for candle := range DbQueue {
			if err := candle.CreateOrUpdate(); err != nil {
				log.WithError(err).Error("saving new candle failed")
			}
		}
	}()
}
