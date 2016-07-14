package main

import (
	"log"
	"net/http"

	"github.com/montanaflynn/is-pokemon-go-online"
)

func handler(w http.ResponseWriter, req *http.Request) {
	status, err := pogo.CheckStatus()
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte(status))
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4444", nil))
}
