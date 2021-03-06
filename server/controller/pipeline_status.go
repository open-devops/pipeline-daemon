package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/model"
	"net/http"
)

func GetPipelineStatus(w http.ResponseWriter, r *http.Request) {
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get pipeline ID from path parameters
	pipelineId := mux.Vars(r)["pipelineId"]

	// Get pipeline fundamental info
	pipelineInfo := model.FetchPipelineInfo(pipelineId)

	// Invalid Pipeline ID supplied
	if len(pipelineInfo.PipelineName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the pipeline's provisioning & running status
	if status, err := model.FetchPipelineStatus(pipelineInfo); err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	} else {
		response, _ := json.Marshal(status)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
