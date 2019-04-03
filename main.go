package main

import (
	"net/http"
	"os"

	"github.com/diegohce/logger"
)

var (
	//Global logger artifact
	log *logger.Logger
)


func main() {

	configFromEnv()

	log = logger.New("kpm-proxy - ")

	if len(config.dirs) == 0 {
		log.Error().Println("Oh! No dirs to scan... :'(")
		os.Exit(1)
	}

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	http.HandleFunc("/kpmproxy/read", wsocketHandler)
	http.HandleFunc("/kpmproxy/write", writeRequest)

	http.ListenAndServe(config.tcpBind, nil)
}
