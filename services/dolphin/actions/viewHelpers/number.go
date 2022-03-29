package viewHelpers

import "github.com/h-varmazyar/Gate/services/dolphin/internal/pkg/app"

var (
	Sum = app.TemplateFunc{
		Key: "Sum",
		Func: func(inputs ...int) int {
			sum := 0
			for _, input := range inputs {
				sum += input
			}
			return sum
		},
	}
)
