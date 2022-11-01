package restmock

import (
	"github.com/gorilla/mux"
	"net/http"
)

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

type LogMiddleware struct {
	logger *HttpLogger
}

func NewLogMiddleware(logger *HttpLogger) *LogMiddleware {
	return &LogMiddleware{logger: logger}
}

func (m *LogMiddleware) Func() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			lrw := NewLoggingResponseWriter(w)
			next.ServeHTTP(lrw, r)
			m.logger.logRequest(r, lrw.statusCode)
		})
	}
}
