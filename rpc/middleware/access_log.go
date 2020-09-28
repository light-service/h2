package middleware

import (
	"context"
	"encoding/json"
	"github.com/light-service/h2/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"os"
	"time"
)

func NewAccessLogInterceptor(logger log.Interface) grpc.UnaryServerInterceptor {
	fieldLogger := log.AdaptFieldLogger(logger)

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		since := time.Now()
		resp, err = handler(ctx, req)

		remoteIP := ""
		hostname, _ := os.Hostname()
		reqJson, _ := json.Marshal(req)
		elapsed := time.Since(since)
		if p, ok := peer.FromContext(ctx); ok {
			remoteIP = p.Addr.String()
		}
		fieldLogger.WithFields(map[string]interface{}{
			"host":          hostname,
			"full_method":   info.FullMethod,
			"remote_ip":     remoteIP,
			"request":       string(reqJson),
			"code":          status.Code(err),
			"error":         err,
			"latency":       int(elapsed),
			"latency_human": elapsed.String(),
		}).Info()
		return resp, err
	}
}
