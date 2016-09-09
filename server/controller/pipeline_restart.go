package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/types"
	"github.com/open-devops/pipeline-daemon/server/types/status"
	"net/http"
)

func RestartPipeline(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	pipelineId := mux.Vars(r)["pipelineId"]

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
	w.Write(response)
}
