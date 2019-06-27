package wsk

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type webParams struct {
	Method string `json:"__ow_method"`
	Path   string `json:"__ow_path"`
}

// HandleFunc handles an OpenWhisk web action.
func HandleFunc(handler http.HandlerFunc) {
	fout := os.NewFile(3, "pipe")
	HandleRWFunc(os.Stdin, fout, handler)
	fout.Close()
}

// HandleRWFunc handles an OpenWhisk web action using the given reader and
// writer for input and output.
func HandleRWFunc(r io.Reader, w io.Writer, handler http.HandlerFunc) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		data := scanner.Bytes()
		var envVar envVariables
		err := json.Unmarshal(data, &envVar)
		if err != nil {
			log.Printf("error getting environment variables: %s", err)
			fmt.Fprintf(w, "error getting environment variables: %s", err)
			continue
		}
		envVar.setVariables()
		var wp webParams
		err = json.Unmarshal(data, &wp)
		if err != nil {
			log.Printf("error getting web HTTP parameters: %s", err)
			fmt.Fprintf(w, "error getting web HTTP parameters: %s", err)
			continue
		}
		req, err := http.NewRequest(wp.Method, "", nil)
		if err != nil {
			log.Printf("error creating new request: %s", err)
			fmt.Fprintf(w, "error creating new request: %s", err)
			continue
		}
		rw := newResponseWriter(w)
		handler(rw, req)
	}
}
