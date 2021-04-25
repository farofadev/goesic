package responses

import (
	"io"
	"net/http"
	"strings"

	"compress/gzip"

	"github.com/julienschmidt/httprouter"
)

type GzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w GzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func GzipHandler(fn httprouter.Handle) httprouter.Handle {
	return func(responseWriter http.ResponseWriter, request *http.Request, params httprouter.Params) {
		if !strings.Contains(request.Header.Get("Accept-Encoding"), "gzip") {
			fn(responseWriter, request, params)
			return
		}

		responseWriter.Header().Set("Content-Encoding", "gzip")
		responseWriter.Header().Set("Vary", "Accept-Encoding")

		gz, _ := gzip.NewWriterLevel(responseWriter, gzip.BestSpeed)

		defer gz.Close()

		gzResponseWriter := GzipResponseWriter{Writer: gz, ResponseWriter: responseWriter}

		fn(gzResponseWriter, request, params)
	}
}
