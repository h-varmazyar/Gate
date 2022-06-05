package viewHelpers

import (
	"github.com/h-varmazyar/Gate/services/Dolphin/internal/pkg/app"
	"time"
)

var (
	TimeStampFormat = app.TemplateFunc{
		Key: "TimeFormatter",
		Func: func(input int64) string {
			return time.Unix(input, 0).Format("06/01/02, 15:04")
		},
	}
)
