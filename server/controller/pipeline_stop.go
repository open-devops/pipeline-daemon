package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/types"
	"github.com/open-devops/pipeline-daemon/server/types/status"
	"net/http"
)

func StopPipeline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	pipelineId := mux.Vars(r)["pipelineId"]

	status := &types.PipelineStatus{
		PipelineId:            pipelineId,
		RequirementManagement: Status.Down,
		SoftwareControlManage: Status.Down,
		ContinuousIntegration: Status.Down,
		CodeQualityInspection: Status.Down,
		RepositoryForArtifact: Status.Down,
		RepositoryOfContainer: Status.Down,
		PipelineDashboard:     Status.Down,
		ContainerManagement:   Status.Down,
	}

	response, _ := json.Marshal(status)
	w.Write(response)
}
