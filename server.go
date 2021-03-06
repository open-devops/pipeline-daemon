package main

import (
	sw "github.com/open-devops/pipeline-daemon/server"
	"log"
	"net/http"
	"os"
)

const (
	DEFAULT_DAEMON_PORT = "8080"
)

func main() {
	log.Println("Pipeline Daemon Server started!")

	router := sw.NewRouter()

	daemonBindingPort := os.Getenv("PIPELINE_DAEMON_PORT")
	if len(daemonBindingPort) == 0 {
		log.Println("Used default daemon binding port!")
		daemonBindingPort = DEFAULT_DAEMON_PORT
	}

	log.Fatal(http.ListenAndServe(":"+daemonBindingPort, router))
}
