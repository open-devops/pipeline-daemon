package model

import (
	"github.com/open-devops/pipeline-daemon/server/types"
	"github.com/open-devops/pipeline-daemon/server/types/status"
)

func FetchPipelineStatus(pipelineId string) *types.PipelineStatus {
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

	return status
}
