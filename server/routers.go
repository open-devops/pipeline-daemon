package server

import (
	"fmt"
	"github.com/gorilla/mux"
	ctl "github.com/open-devops/pipeline-daemon/server/controller"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
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
		handler = utl.Logger(handler, route.Name)

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
		ctl.GetPipelineStatus,
	},

	Route{
		"OperatePipeline",
		"POST",
		"/v1/pipelines/{pipelineId}/{operation}",
		ctl.OperatePipeline,
	},

	Route{
		"OperateCapability",
		"POST",
		"/v1/pipelines/{pipelineId}/{operation}/{capability}",
		ctl.OperatePipeline,
	},

	Route{
		"AddProvision",
		"POST",
		"/v1/pipelines/{pipelineId}/provision",
		ctl.AddProvision,
	},

	Route{
		"DeleteProvision",
		"DELETE",
		"/v1/pipelines/{pipelineId}/provision",
		ctl.DeleteProvision,
	},

	Route{
		"GetProvision",
		"GET",
		"/v1/pipelines/{pipelineId}/provision",
		ctl.GetProvision,
	},
}
