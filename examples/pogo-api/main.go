package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/montanaflynn/is-pokemon-go-online"
)

var port string

func init() {
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = ":8080"
	}
}

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
	port = fmt.Sprintf(":%s", strings.Trim(port, ":"))
	log.Fatal(http.ListenAndServe(port, nil))
}
