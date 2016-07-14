package pogo

import (
	"errors"
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
)

func dialTimeout(network, addr string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, addr, ConnectionTimeout)
	conn.SetDeadline(time.Now().Add(ResponseTimeout))
	return conn, err
}

func CheckStatus() (string, error) {
	client := http.Client{
		Transport: &http.Transport{
			Dial: dialTimeout,
		},
	}

	resp, err := client.Get(serverURL)
	if err != nil {
		return serverName + " is not online", err
	}

	if resp.StatusCode != 200 {
		return serverName + " is not working", errors.New(resp.Status)
	}

	return serverName + " is online", nil

}
