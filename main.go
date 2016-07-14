package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"time"
)

const (
	serverName = "Pokémon Go"
	serverURL  = "https://pgorelease.nianticlabs.com/plfe/"
)

var (
	ConnectionTimeout = 5 * time.Second
	ResponseTimeout   = 5 * time.Second
	VerboseMode       = false
)

func dialTimeout(network, addr string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, addr, ConnectionTimeout)
	conn.SetDeadline(time.Now().Add(ResponseTimeout))
	return conn, err
}

func checkStatus(resource, url string) string {
	client := http.Client{
		Transport: &http.Transport{
			Dial: dialTimeout,
		},
	}

	resp, err := client.Get(url)
	if err != nil {
		if VerboseMode == true {
			log.Println(err)
		}
		return resource + " is not online"
	}

	if resp.StatusCode != 200 {
		if VerboseMode == true {
			log.Println(errors.New("Returned " + resp.Status))
		}
		return resource + " is not working"
	}

	return resource + " is online"

}

func main() {
	status := checkStatus(serverName, serverURL)
	log.Println(status)
}
