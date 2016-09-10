package model

import (
	"fmt"
	"github.com/open-devops/pipeline-daemon/server/types"
	"github.com/open-devops/pipeline-daemon/server/types/status"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os/exec"
)

func FetchPipelineStatus(pipelineInfo *types.PipelineInfo) *types.PipelineStatus {
	// Engine program path
	engine := utl.GetEnginePath(pipelineInfo)

	// Engine program parameter
	args := []string{"status"}

	// Pipeline status information
	var status *types.PipelineStatus

	fmt.Println("engine = " + engine)

	// Fetch pipeline status information using engine program
	if out, err := exec.Command(engine, args...).Output(); err != nil {
		status = &types.PipelineStatus{
			PipelineId:            pipelineInfo.PipelineId,
			RequirementManagement: Status.Unknown,
			SoftwareControlManage: Status.Unknown,
			ContinuousIntegration: Status.Unknown,
			CodeQualityInspection: Status.Unknown,
			RepositoryForArtifact: Status.Unknown,
			RepositoryOfContainer: Status.Unknown,
			PipelineDashboard:     Status.Unknown,
			ContainerManagement:   Status.Unknown,
		}
	} else {
		fmt.Println(string(out))
		status = &types.PipelineStatus{
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
	}

	return status
}
