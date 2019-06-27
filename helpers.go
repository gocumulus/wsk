package wsk

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON responds with JSON using the given responsewriter, response
// code, and data.
func RespondWithJSON(w http.ResponseWriter, code int, data interface{}) {
	content, err := json.Marshal(data)
	if string(content) == "null" {
		content = []byte("[]")
	}
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(content)
}

// RespondWithError responds with an error in JSON using the given
// responsewriter, response code, and data.
func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithOptions responds to an HTTP request with allowed options.
func RespondWithOptions(w http.ResponseWriter, options string) {
	w.Header().Set("Allow", options)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(""))
}
