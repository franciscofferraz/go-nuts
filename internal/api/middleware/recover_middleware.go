package middleware

import (
	"fmt"
	"net/http"

	"github.com/franciscofferraz/go-struct/internal/errors"
)

func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				errors.ServerErrorResponse(w, r, fmt.Errorf("%s", err))

			}
		}()

		next.ServeHTTP(w, r)
	})
}
