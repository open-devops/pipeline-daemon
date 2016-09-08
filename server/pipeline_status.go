package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/open-devops/pipeline-daemon/server/types"
	"net/http"
)

func GetPipelineStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	pipelineId := mux.Vars(r)["pipelineId"]

	status := &types.PipelineStatus{
		PipelineId:            pipelineId,
		RequirementManagement: "Up",
		SoftwareControlManage: "Up",
		ContinuousIntegration: "Up",
		CodeQualityInspection: "Up",
		RepositoryForArtifact: "Up",
		RepositoryOfContainer: "Up",
		PipelineDashboard:     "Up",
		ContainerManagement:   "Up",
	}

	response, _ := json.Marshal(status)
	w.Write(response)
}
