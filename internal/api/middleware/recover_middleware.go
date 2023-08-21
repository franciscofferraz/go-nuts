package middleware

import (
	"fmt"
	"net/http"

	"github.com/franciscofferraz/go-struct/internal/customerrors"
)

func RecoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "close")
				customerrors.ServerErrorResponse(w, r, fmt.Errorf("%s", err))

			}
		}()

		next.ServeHTTP(w, r)
	})
}
