package viewHelpers

import (
	"github.com/h-varmazyar/Gate/services/Dolphin/internal/pkg/app"
	"strconv"
)

var (
	ThousandFormat = app.TemplateFunc{
		Key: "ThousandFormatter",
		Func: func(input int64) string {
			in := strconv.FormatInt(input, 10)
			numOfDigits := len(in)
			if input < 0 {
				numOfDigits-- // First character is the - sign (not a digit)
			}
			numOfCommas := (numOfDigits - 1) / 3

			out := make([]byte, len(in)+numOfCommas)
			if input < 0 {
				in, out[0] = in[1:], '-'
			}

			for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
				out[j] = in[i]
				if i == 0 {
					return string(out)
				}
				if k++; k == 3 {
					j, k = j-1, 0
					out[j] = ','
				}
			}
		},
	}
)
