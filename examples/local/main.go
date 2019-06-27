package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gocumulus/wsk"
)

func main() {
	r := strings.NewReader(`{"action_name":"/matthew@example.com_dev/standalone","activation_id":"1234","api_key":"5678:abcd","deadline":"1561642521187","namespace":"matthew@example.com_dev","value":{}}`)
	w := os.Stdout
	wsk.HandleRWFunc(r, w, myHandler)
	fmt.Println()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	todo := struct {
		Title string `json:"title"`
	}{
		Title: "First task",
	}
	wsk.RespondWithJSON(w, http.StatusOK, todo)
}
