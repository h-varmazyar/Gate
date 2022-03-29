package models

import "time"

/**
* Dear programmer:
* When I wrote this code, only god And I know how it worked.
* Now, only god knows it!
*
* Therefore, if you are trying to optimize this code And it fails(most surely),
* please increase this counter as a warning for the next person:
*
* total_hours_wasted_here = 0 !!!
*
* Best regards, mr-nobody
* Date: 04.12.21
* Github: https://github.com/h-varmazyar
* Email: hossein.varmazyar@yahoo.com
**/

type Candle struct {
	Open         float64
	High         float64
	Low          float64
	Close        float64
	Volume       float64
	Amount       float64
	Time         time.Time
	MarketID     string
	ResolutionID string
	Indicators
}
