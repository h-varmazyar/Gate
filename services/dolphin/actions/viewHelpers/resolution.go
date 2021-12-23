package viewHelpers

import (
	brokerageApi "github.com/mrNobody95/Gate/services/brokerage/api"
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
)

var (
	ResolutionLabel = app.TemplateFunc{
		Key: "ResolutionLabel",
		Func: func(resolution *brokerageApi.Resolution) string {
			return resolution.Label
		},
	}
)
