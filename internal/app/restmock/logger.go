package restmock

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
	"time"
)

type HttpLogger struct {
	logger *zap.SugaredLogger
}

func NewLogger() *HttpLogger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, err := zapConfig.Build()
	if err != nil {
		fmt.Println("failed to start logger")
	}
	sugar := logger.Sugar()
	return &HttpLogger{sugar}
}

func (h *HttpLogger) logRequest(r *http.Request, statusCode int) {
	h.logger.Infow("Request processing finished",
		"method", r.Method,
		"host", r.Host,
		"request", r.RequestURI,
		"protocol", r.Proto,
		"headers:", r.Header,
		"body", r.Body,
		"status_code", statusCode,
		"user_agent", r.UserAgent(),
		"remote_address", r.RemoteAddr,
		"referer", r.Referer(),
	)
}
