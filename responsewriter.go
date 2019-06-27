package wsk

import (
	"encoding/json"
	"io"
	"net/http"
)

type responseWriter struct {
	wroteHeader bool
	header      http.Header
	w           io.Writer
	statusCode  int
}

func (rw *responseWriter) Header() http.Header {
	return rw.header
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	if !rw.wroteHeader {
		rw.WriteHeader(http.StatusOK)
	}
	d := struct {
		StatusCode int             `json:"statusCode"`
		Header     http.Header     `json:"headers"`
		Data       json.RawMessage `json:"body"`
	}{
		StatusCode: rw.statusCode,
		Header:     rw.header,
		Data:       data,
	}
	output, err := json.Marshal(d)
	if err != nil {
		// FIXME(mdr): Log error.
	}
	return rw.w.Write(output)
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.wroteHeader = true
	rw.statusCode = statusCode
}

func newResponseWriter(w io.Writer) *responseWriter {
	rw := responseWriter{
		wroteHeader: false,
		w:           w,
		header:      make(http.Header),
	}
	return &rw
}
