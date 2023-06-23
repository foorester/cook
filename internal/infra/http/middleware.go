package http

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/foorester/cook/internal/sys/log"
)

type (
	ContextKey string
)

const (
	BookCtxKey = "book"
)

func BookContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bookID := "a4e52de2-352b-4a61-964b-7efb7c137538" // WIP: Extract it from path

		ctx := context.WithValue(r.Context(), BookCtxKey, bookID)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// Request logger

const (
	tsFormat = "2006/01/02 15:04:05"
)

type (
	ReqLogger struct {
		log log.Logger
	}
)

func (rl *ReqLogger) Log() log.Logger {
	return rl.log
}

func NewReqLogger(log log.Logger) *ReqLogger {
	return &ReqLogger{log: log}
}

func NewReqLoggerMiddleware(log log.Logger) func(next http.Handler) http.Handler {
	return middleware.RequestLogger(NewReqLogger(log))
}

func (rl *ReqLogger) NewLogEntry(r *http.Request) middleware.LogEntry {
	fields := map[string]string{}

	if reqID := middleware.GetReqID(r.Context()); reqID != "" {
		fields["req-id"] = reqID
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	fields["scheme"] = scheme
	fields["proto"] = r.Proto
	fields["method"] = r.Method
	fields["addr"] = r.RemoteAddr
	fields["agent"] = r.UserAgent()
	fields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)
	fields["ts"] = time.Now().UTC().Format(tsFormat)

	sb := strings.Builder{}
	for k, v := range fields {
		sb.WriteString(fmt.Sprintf("%s: %s, ", k, v))
	}

	return NewLogEntry(rl.Log(), &sb)
}

type (
	LogEntry struct {
		log   log.Logger
		entry *strings.Builder
	}
)

func NewLogEntry(log log.Logger, sb *strings.Builder) *LogEntry {
	return &LogEntry{
		log:   log,
		entry: sb,
	}
}

func (le *LogEntry) Log() log.Logger {
	return le.log
}

func (le *LogEntry) Write(status, bytes int, header http.Header, elapsed time.Duration, extra interface{}) {
	le.entry.WriteString(fmt.Sprintf("%s: %d, ", "status", status))
	le.entry.WriteString(fmt.Sprintf("%s: %d, ", "bytes", bytes))
	le.entry.WriteString(fmt.Sprintf("%s: %fms", "elapsed", float64(elapsed.Nanoseconds())/1000000.0))
	le.Log().Debugf("%s", le.entry.String())
}

func (le *LogEntry) Panic(v interface{}, stack []byte) {
	le.entry.WriteString(fmt.Sprintf("%s: %s, ", "stack", string(stack)))
	le.entry.WriteString(fmt.Sprintf("%s: %s, ", "panic", fmt.Sprintf("%+v", v)))
	le.Log().Debugf("%s", le.entry.String())
}
