package httpext

import (
	"github.com/h-varmazyar/Gate/pkg/errors"
	"github.com/rs/cors"

	"google.golang.org/grpc/codes"
	"net/http"
	"strings"
)

var (
	DefaultCors = cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
		},
		AllowedHeaders: []string{"*"},
	})
)

func Origin(req *http.Request) (string, error) {
	if origin := req.Header.Get("Origin"); origin != "" {
		origin = strings.ReplaceAll(origin, "http://", "")
		origin = strings.ReplaceAll(origin, "https://", "")
		return strings.Trim(origin, "/"), nil
	}
	if host := req.Header.Get("Host"); host != "" {
		host = strings.ReplaceAll(host, "http://", "")
		host = strings.ReplaceAll(host, "https://", "")
		return strings.Trim(host, "/"), nil
	}
	return "", errors.New(req.Context(), codes.InvalidArgument).
		AddDetails("can find request origin")
}
