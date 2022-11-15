package httpext

import (
	"encoding/json"
	"github.com/h-varmazyar/Gate/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

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
