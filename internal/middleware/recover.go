package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/ffajarpratama/pos-wash-api/internal/http/response"
	"github.com/ffajarpratama/pos-wash-api/pkg/constant"
)

func Recoverer(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 2048)
				n := runtime.Stack(buf, false)
				buf = buf[:n]

				log.Println(red(err))
				fmt.Println(string(buf))

				w.Header().Set("Content-Type", "application/json; charset=utf-8")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(response.JsonResponse{
					Error: &response.ErrorResponse{
						Code:    constant.DefaultUnhandledError,
						Status:  http.StatusInternalServerError,
						Message: constant.HTTPStatusText(http.StatusInternalServerError),
					},
				})
			}
		}()

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
