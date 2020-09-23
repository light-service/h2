package middleware

import (
	"context"
	"encoding/json"
	"github.com/light-service/h2/log"
	"google.golang.org/grpc"
	"time"
)

func NewAccessLogInterceptor(logger log.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		since := time.Now()
		defer func() {
			elapsed := time.Since(since)
			reqJson, _ := json.Marshal(req)
			logger.Infof("%s: req=%s, err=%s, elapsed=%.2f", info.FullMethod, reqJson, err, float64(elapsed)/float64(time.Second))
		}()

		resp, err = handler(ctx, req)
		return resp, err
	}
}
