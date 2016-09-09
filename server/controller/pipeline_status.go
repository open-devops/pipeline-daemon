package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/model"
	"github.com/open-devops/pipeline-daemon/server/types"
	"github.com/open-devops/pipeline-daemon/server/types/status"
	"net/http"
)

func GetPipelineStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	// Get pipeline ID from path parameters
	pipelineId := mux.Vars(r)["pipelineId"]

	// Get pipeline fundamental info
	pipelineInfo := model.FetchPipelineInfo(pipelineId)

	debugInfo, _ := json.Marshal(pipelineInfo)
	fmt.Println(string(debugInfo))

	// Invalid Pipeline ID supplied
	if len(pipelineInfo.PipelineName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Return the pipeline's provisioning & running status
	status := &types.PipelineStatus{
		PipelineId:            pipelineId,
		RequirementManagement: Status.Up,
		SoftwareControlManage: Status.Up,
		ContinuousIntegration: Status.Up,
		CodeQualityInspection: Status.Up,
		RepositoryForArtifact: Status.Up,
		RepositoryOfContainer: Status.Up,
		PipelineDashboard:     Status.Up,
		ContainerManagement:   Status.Up,
	}

	response, _ := json.Marshal(status)

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
