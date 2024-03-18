package logger

import (
	"context"
	"io"
	"log/slog"
)

type TraceIDFn func(ctx context.Context) string

type Logger struct {
	handler slog.Handler
	traceIDFn TraceIDFn
}

func (log *Logger) Info(ctx context.Context,msg string ,args ..any){
	log.
}


func new(w io.Writer, minLevel Level)