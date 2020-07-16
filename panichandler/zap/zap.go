package zap

import (
	"context"
	"fmt"
	"runtime"

	"github.com/higebu/go-grpc-interceptor/zap/zapctx"
	"go.uber.org/zap"
)

func LogPanicWithStackTrace(ctx context.Context, r interface{}) {
	logger, ok := zapctx.FromContext(ctx)
	if !ok {
		return
	}
	logger.Error("recovered from panic",
		zap.String("panic", fmt.Sprintf("%v", r)),
		zapStack(),
	)
}

func zapStack() zap.Field {
	callers := []string{}
	for i := 0; true; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fn := runtime.FuncForPC(pc)
		callers = append(callers, fmt.Sprintf("%s(%d): %s", file, line, fn.Name()))
	}
	return zap.Any("stacktrace", callers)
}
