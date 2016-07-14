package main

import (
	"log"

	"github.com/montanaflynn/is-pokemon-go-online"
)

func main() {
	status, err := pogo.CheckStatus()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(status)
}
