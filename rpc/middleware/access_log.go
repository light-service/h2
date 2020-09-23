package middleware

import (
	"context"
	"encoding/json"
	"github.com/light-service/h2/log"
	"google.golang.org/grpc"
	"time"
)

func NewAccessLogInterceptor(logger log.Interface) grpc.UnaryServerInterceptor {
	fieldLogger := log.AdaptFieldLogger(logger)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		since := time.Now()
		defer func() {
			elapsed := time.Since(since)
			reqJson, _ := json.Marshal(req)
			fieldLogger.WithFields(map[string]interface{}{
				"full_method": info.FullMethod,
				"request":     reqJson,
				"error":       err,
				"elapsed":     elapsed,
			}).Info()
		}()

		resp, err = handler(ctx, req)
		return resp, err
	}
}
