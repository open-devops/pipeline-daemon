package model

import (
	"fmt"
	"github.com/open-devops/pipeline-daemon/server/types"
	"github.com/open-devops/pipeline-daemon/server/types/status"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os/exec"
)

func FetchPipelineStatus(pipelineInfo *types.PipelineInfo) *types.PipelineStatus {
	// Engine shell program path
	engineShell := utl.GetEnginePath(pipelineInfo)

	// Engine shell parameter
	args := []string{"status"}

	var status *types.PipelineStatus

	fmt.Println("engineShell = " + engineShell)

	if out, err := exec.Command(engineShell, args...).Output(); err != nil {
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
