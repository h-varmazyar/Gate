package indicators

type Indicator interface {
	Calculate()
	Update()
}

type Source string

const (
	SourceCustom Source = "custom"
	SourceOHLC4  Source = "ohlc4"
	SourceClose  Source = "close"
	SourceOpen   Source = "open"
	SourceHigh   Source = "high"
	SourceHLC3   Source = "hlc3"
	SourceLow    Source = "low"
	SourceHL2    Source = "hl2"
)
