package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Pipeline Daemon Server")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/v1/",
		Index,
	},

	Route{
		"GetPipelineStatus",
		"GET",
		"/v1/pipelines/{pipelineId}/status",
		GetPipelineStatus,
	},

	Route{
		"RestartPipeline",
		"POST",
		"/v1/pipelines/{pipelineId}/restart",
		RestartPipeline,
	},

	Route{
		"StartPipeline",
		"POST",
		"/v1/pipelines/{pipelineId}/start",
		StartPipeline,
	},

	Route{
		"StopPipeline",
		"POST",
		"/v1/pipelines/{pipelineId}/stop",
		StopPipeline,
	},

	Route{
		"AddProvision",
		"POST",
		"/v1/pipelines/{pipelineId}/provision",
		AddProvision,
	},

	Route{
		"DeleteProvision",
		"DELETE",
		"/v1/pipelines/{pipelineId}/provision",
		DeleteProvision,
	},

	Route{
		"GetProvision",
		"GET",
		"/v1/pipelines/{pipelineId}/provision",
		GetProvision,
	},
}
