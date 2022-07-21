package logger

import (
	"github.com/jwalton/gchalk"
	"net/http"
	"runtime/debug"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

func LoggingMiddleware(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Info(log.Fields{
						"err":   err,
						"trace": string(debug.Stack()),
					})
				}
			}()

			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)
			logger.Infof("[%s][%s][%s][%s]",
				gchalk.BrightWhite(time.Since(start).String()),
				gchalk.BrightYellow(strconv.Itoa(wrapped.status)),
				gchalk.BrightRed(r.Method),
				gchalk.BrightCyan(r.URL.EscapedPath()),
			)
		}

		return http.HandlerFunc(fn)
	}
}
