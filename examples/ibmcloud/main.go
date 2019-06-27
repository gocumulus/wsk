package main

import (
	"net/http"

	"github.com/gocumulus/wsk"
)

func main() {
	wsk.HandleFunc(myHandler)
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	todo := struct {
		Title string `json:"title"`
	}{
		Title: "First task",
	}
	wsk.RespondWithJSON(w, http.StatusOK, todo)
}
