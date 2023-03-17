package handler

import (
	"net/http"

	"github.com/wuhc1010/go-zero/core/logx"
	"github.com/wuhc1010/go-zero/core/syncx"
	"github.com/wuhc1010/go-zero/rest/internal"
)

// MaxConnsHandler returns a middleware that limit the concurrent connections.
func MaxConnsHandler(n int) func(http.Handler) http.Handler {
	if n <= 0 {
		return func(next http.Handler) http.Handler {
			return next
		}
	}

	return func(next http.Handler) http.Handler {
		latch := syncx.NewLimit(n)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if latch.TryBorrow() {
				defer func() {
					if err := latch.Return(); err != nil {
						logx.Error(err)
					}
				}()

				next.ServeHTTP(w, r)
			} else {
				internal.Errorf(r, "concurrent connections over %d, rejected with code %d",
					n, http.StatusServiceUnavailable)
				w.WriteHeader(http.StatusServiceUnavailable)
			}
		})
	}
}
