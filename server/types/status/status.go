package Status

import (
	"github.com/open-devops/pipeline-daemon/server/types"
)

const (
	Up      = "Up"
	Down    = "Down"
	Unknown = "N/A"
)

func StatusAs(pipelineInfo *types.PipelineInfo, status string) *types.PipelineStatus {
	return &types.PipelineStatus{
		PipelineId:            pipelineInfo.PipelineId,
		RequirementManagement: status,
		SoftwareControlManage: status,
		ContinuousIntegration: status,
		CodeQualityInspection: status,
		RepositoryForArtifact: status,
		RepositoryOfContainer: status,
		PipelineDashboard:     status,
		ContainerManagement:   status,
	}
}
