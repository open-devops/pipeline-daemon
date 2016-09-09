package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/model"
	"github.com/open-devops/pipeline-daemon/server/types"
	"github.com/open-devops/pipeline-daemon/server/types/status"
	"net/http"
)

func StartPipeline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Get pipeline ID from path parameters
	pipelineId := mux.Vars(r)["pipelineId"]

	// Get pipeline fundamental info
	pipelineInfo := model.FetchPipelineInfo(pipelineId)

	// Invalid Pipeline ID supplied
	if len(pipelineInfo.PipelineName) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Return the pipeline's provisioning & running status
	status := &types.PipelineStatus{
		PipelineId:            pipelineInfo.PipelineId,
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
	w.Write(response)
}
