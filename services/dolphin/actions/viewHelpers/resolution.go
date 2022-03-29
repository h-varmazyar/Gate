package viewHelpers

import (
	brokerageApi "github.com/h-varmazyar/Gate/services/brokerage/api"
	"github.com/h-varmazyar/Gate/services/dolphin/internal/pkg/app"
)

var (
	ResolutionLabel = app.TemplateFunc{
		Key: "ResolutionLabel",
		Func: func(resolution *brokerageApi.Resolution) string {
			return resolution.Label
		},
	}
)
