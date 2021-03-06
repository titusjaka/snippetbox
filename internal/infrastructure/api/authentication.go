package api

import (
	"context"
	"net/http"
	"strings"

	kithttp "github.com/go-kit/kit/transport/http"

	"github.com/titusjaka/go-sample/internal/infrastructure/log"
)

type authenticatorKey int

const (
	// AuthorizationHeaderKey is the value used to place
	// the authorization header content in context.Context.
	AuthorizationHeaderKey authenticatorKey = iota
)

// AuthorizationHeader tries to retrieve the token string from the "Authorization"
// request header and adds it to the context.
func AuthorizationHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r.WithContext(
			context.WithValue(
				r.Context(),
				AuthorizationHeaderKey,
				tokenFromHeader(r),
			)),
		)
	})
}

// InternalCommunication performs bearer authentication with provided token
// for internal service communication purposes.
func InternalCommunication(token string, logger log.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logError := func() {
				logger.Info("Unauthorized request",
					log.Field("method", r.Method),
					log.Field("host", r.Host),
					log.Field("remote_addr", r.RemoteAddr),
					log.Field("uri", r.URL.RequestURI()),
					log.Field("status", http.StatusUnauthorized),
				)
			}
			if v, ok := r.Context().Value(AuthorizationHeaderKey).(string); !ok || v != token {
				logError()
				_ = kithttp.EncodeJSONResponse(context.Background(), w, ErrUnauthorized())
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func tokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")

	if len(bearer) > 7 && strings.EqualFold(bearer[0:6], "bearer") {
		return bearer[7:]
	}

	return ""
}
