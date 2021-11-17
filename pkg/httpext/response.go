package httpext

import (
	"encoding/json"
	"github.com/mrNobody95/Gate/pkg/errors"
	log "github.com/sirupsen/logrus"
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
* Date: 13.11.21
* Github: https://github.com/mrNobody95
* Email: hossein.varmazyar@yahoo.com
**/

func SendError(res http.ResponseWriter, req *http.Request, err error) {
	errModel := errors.Cast(req.Context(), err)
	SendModel(res, req, errModel.HttpStatus(), errModel)
}

func SendModel(res http.ResponseWriter, req *http.Request, code int, model interface{}) {
	jsonModel, err := json.Marshal(model)
	if err != nil {
		log.WithError(err).Error("cannot marshal model")
	}
	SendData(res, req, code, string(MimeJson), jsonModel)
}

func SendData(res http.ResponseWriter, req *http.Request, code int, mime string, data []byte) {
	res.Header().Set(ContentTypeHeader, mime)
	res.Header().Set(CharsetHeader, "utf-8")
	res.WriteHeader(code)
	_, err := res.Write(data)
	if err != nil {
		log.WithError(err).Error("write response data failed")
	}
}

func SendCode(res http.ResponseWriter, _ *http.Request, code int) {
	res.WriteHeader(code)
}
