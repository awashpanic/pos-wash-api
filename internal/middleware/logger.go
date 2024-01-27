package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	custom_jwt "github.com/ffajarpratama/pos-wash-api/pkg/jwt"
	"github.com/ffajarpratama/pos-wash-api/pkg/util"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type userLog struct {
	ID   string
	Role string
}

type customResponseWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func (r customResponseWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r *customResponseWriter) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

func Logger(log *logrus.Logger) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			ww := &customResponseWriter{body: &bytes.Buffer{}, ResponseWriter: w}

			bodyJSON, err := io.ReadAll(r.Body)
			if err != nil {
				log.Printf("[error-io-reader] \n%v\n", err)
			}

			var payload map[string]interface{}
			if len(bodyJSON) > 0 {
				json.Unmarshal(bodyJSON, &payload)
			}

			token, _ := util.GetTokenFromHeader(r)
			userClaims := parseWithoutVerified(token)
			defer func() {
				fields := map[string]interface{}{
					"@path": fmt.Sprintf("[%d][%s] %s", ww.statusCode, r.Method, r.RequestURI),
					"@time": util.TimeNow().Format(time.DateTime),
					"body":  payload,
				}

				if token != "" && userClaims != nil {
					fields["auth"] = map[string]interface{}{
						"user": fmt.Sprintf("id: %s, role: %s", userClaims.ID, userClaims.Role),
					}
				}

				err = r.Body.Close()
				if err != nil {
					log.Printf("[error-body-close] \n%v\n", err)
				}

				if !isExcludeRouter(r.RequestURI) {
					log.WithFields(fields).Info()
				}
			}()

			// create new body for handlers
			newBody := io.NopCloser(bytes.NewBuffer(bodyJSON))
			r.Body = newBody

			h.ServeHTTP(ww, r)
		}

		return http.HandlerFunc(fn)
	}
}

func isExcludeRouter(path string) bool {
	if strings.Contains(path, "media") {
		return true
	}

	if strings.Contains(path, "import") {
		return true
	}

	// list of excluded routes
	excluded := make(map[string]bool)

	if _, ok := excluded[path]; ok {
		return true
	}

	return false
}

func parseWithoutVerified(tokenString string) *userLog {
	resJwt, _, err := new(jwt.Parser).ParseUnverified(tokenString, &custom_jwt.CustomUserClaims{})
	if err != nil {
		return nil
	}

	userClaims, ok := resJwt.Claims.(*custom_jwt.CustomUserClaims)
	if ok && userClaims.ID != "" {
		return &userLog{
			ID:   userClaims.ID,
			Role: userClaims.Role,
		}
	}

	return nil
}

// // getTraceID returns a request ID from the given context if one is present.
// // Returns the empty string if a request ID cannot be found.
// func getTraceID(ctx context.Context) string {
// 	if ctx == nil {
// 		return ""
// 	}

// 	if reqID, ok := ctx.Value(chi_middleware.RequestIDKey).(string); ok {
// 		return reqID
// 	}

// 	return ""
// }
