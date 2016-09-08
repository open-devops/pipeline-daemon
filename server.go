package main

import (
	sw "github.com/open-devops/pipeline-daemon/server"
	"log"
	"net/http"
)

func main() {
	log.Printf("Pipeline Daemon Server started!")

	router := sw.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
