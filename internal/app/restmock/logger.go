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
	zap_config := zap.NewProductionConfig()
	zap_config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	logger, err := zap_config.Build()
	if err != nil{
		fmt.Println("failed to start logger")
	}
	sugar := logger.Sugar()
	return &HttpLogger{sugar}
}

func (h *HttpLogger) logRequest(r *http.Request, status_code int){
	h.logger.Infow("Request processing finished",
		"method", r.Method,
		"host", r.Host,
		"request", r.RequestURI,
		"status_code", status_code,
		"user_agent", r.UserAgent(),
		"remote_address", r.RemoteAddr,
		"referer", r.Referer(),
	)
}
