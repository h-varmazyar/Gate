package httpext

import (
	"bytes"
	"encoding/json"
	"github.com/h-varmazyar/Gate/pkg/errors"
	"google.golang.org/grpc/codes"
	"net/http"
)

func BindModel(req *http.Request, model interface{}) error {
	//if header := req.Header.Get("Content-Type"); header != string(MimeJson) {
	//	return errors.New("content-Type header is not application/json")
	//}
	//
	//dec := json.NewDecoder(req.Body)
	//dec.DisallowUnknownFields()
	//
	//err := dec.Decode(&model)
	//if err != nil {
	//	var syntaxError *json.SyntaxError
	//	var unmarshalTypeError *json.UnmarshalTypeError
	//
	//	switch {
	//	case errors.As(err, &syntaxError):
	//		return fmt.Errorf("request body contains badly-formed JSON (at position %d)", syntaxError.Offset)
	//
	//	case errors.Is(err, io.ErrUnexpectedEOF):
	//		return errors.New("request body contains badly-formed JSON")
	//
	//	case errors.As(err, &unmarshalTypeError):
	//		return fmt.Errorf("request body contains an invalid value for the %q field (at position %d)", unmarshalTypeError.Field, unmarshalTypeError.Offset)
	//
	//	case strings.HasPrefix(err.Error(), "json: unknown field "):
	//		fieldName := strings.TrimPrefix(err.Error(), "json: unknown field ")
	//		return fmt.Errorf("request body contains unknown field %s", fieldName)
	//
	//	case errors.Is(err, io.EOF):
	//		return errors.New("request body must not be empty")
	//
	//	case err.Error() == "http: request body too large":
	//		return errors.New("request body must not be larger than 1MB")
	//
	//	default:
	//		return err
	//	}
	//}
	//
	//err = dec.Decode(&struct{}{})
	//if err != io.EOF {
	//	return errors.New("request body must only contain a single JSON object")
	//}
	//
	//return nil
	var err error
	defer func() {
		_ = req.Body.Close()
		if err != nil {
			err = errors.New(req.Context(), codes.InvalidArgument).
				AddDetails(err.Error())
		}
	}()
	var obj map[string]interface{}
	if err = json.NewDecoder(req.Body).Decode(&obj); err != nil {
		return err
	}
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err = encoder.Encode(obj); err != nil {
		return err
	}
	return json.NewDecoder(&buf).Decode(model)
}
