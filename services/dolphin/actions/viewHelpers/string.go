package viewHelpers

import (
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
	"strings"
)

var (
	EqURL = app.TemplateFunc{
		Key: "EqURL",
		Func: func(url, checker string) bool {
			return strings.TrimSuffix(strings.TrimSpace(url), "/") == strings.TrimSuffix(strings.TrimSpace(checker), "/")
		},
	}
)
