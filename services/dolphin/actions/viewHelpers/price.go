package viewHelpers

import (
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
	"strconv"
)

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
* Date: 11.12.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

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
