package muxext

import (
	"github.com/gorilla/mux"
	"net/http"
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
* Date: 16.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

func NewRouter(strictSlash bool) *mux.Router {
	return mux.NewRouter().StrictSlash(strictSlash)
}

func PathParam(req *http.Request, name string) string {
	return mux.Vars(req)[name]
}
