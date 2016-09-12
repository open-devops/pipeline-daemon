package controller

import (
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/model"
	"net/http"
	"fmt"
)

func AddProvision(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get pipeline ID from path parameters
	pipelineId := mux.Vars(r)["pipelineId"]
	//capability := mux.Vars(r)["capability"]

	// Get pipeline fundamental info
	pipelineInfo := model.FetchPipelineInfo(pipelineId)

	// Invalid Pipeline ID supplied
	if len(pipelineInfo.PipelineName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create new provision environment for the pipeline
	if err := model.CreateProvision(pipelineInfo); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
