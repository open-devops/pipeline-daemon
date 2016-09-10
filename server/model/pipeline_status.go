package model

import (
	"fmt"
	"github.com/open-devops/pipeline-daemon/server/types"
	sta "github.com/open-devops/pipeline-daemon/server/types/status"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os"
	"os/exec"
)

func FetchPipelineStatus(pipelineInfo *types.PipelineInfo) *types.PipelineStatus {
	// Engine program path
	engineParentPath := utl.GetEngineParentPath(pipelineInfo)

	// Engine program name
	engineProgramPath := utl.GetEngineProgramPath(pipelineInfo)

	// Engine program parameter
	args := []string{"status",  "|grep " + utl.GetContainerPrefix(pipelineInfo.PipelineName)}

	// Pipeline status information
	var status *types.PipelineStatus

	// Fetch pipeline status information using engine program
	if err := os.Chdir(engineParentPath); err != err {
		return sta.StatusAs(pipelineInfo, sta.Unknown)
	}

	// Dedicate the status fetch job to pipeline engine
	if out, err := exec.Command(engineProgramPath, args...).Output(); err != nil {
		return sta.StatusAs(pipelineInfo, sta.Unknown)
	} else {
		fmt.Println(string(out))
		return sta.StatusAs(pipelineInfo, sta.Up)
	}

	return status
}
