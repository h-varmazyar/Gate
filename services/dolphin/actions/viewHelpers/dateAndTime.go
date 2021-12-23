package viewHelpers

import (
	"github.com/mrNobody95/Gate/services/dolphin/internal/pkg/app"
	"time"
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
	TimeStampFormat = app.TemplateFunc{
		Key: "TimeFormatter",
		Func: func(input int64) string {
			return time.Unix(input, 0).Format("06/01/02, 15:04")
		},
	}
)
