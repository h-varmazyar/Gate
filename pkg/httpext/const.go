package httpext

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
* Date: 13.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

type MIME string

const (
	MimeJson MIME = "application/json"
)

const (
	ContentTypeHeader   = "Content-Type"
	CharsetHeader       = "Accept-Charset"
	AcceptHeader        = "Accept"
	AuthorizationHeader = "Authorization"
)
