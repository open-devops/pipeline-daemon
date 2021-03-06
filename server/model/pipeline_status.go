package model

import (
	"fmt"
	"github.com/open-devops/pipeline-daemon/server/types"
	sta "github.com/open-devops/pipeline-daemon/server/types/status"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os"
	"os/exec"
)

func FetchPipelineStatus(pipelineInfo *types.PipelineInfo) (*types.PipelineStatus, error) {
	// Engine program parent folder
	engineParentPath := utl.GetEngineParentPath(pipelineInfo)

	// Engine program full path
	engineProgramPath := utl.GetEngineProgramPath(pipelineInfo)

	// Engine program parameter
	containerPrefix := utl.GetContainerPrefix(pipelineInfo.PipelineName)
	args := []string{"status",
		"|grep " + containerPrefix}

	// Change current working folder to engine path
	if err := os.Chdir(engineParentPath); err != nil {
		return sta.StatusAs(pipelineInfo, sta.Unknown), err
	}

	// Dedicate the status fetch job to pipeline engine
	if out, err := exec.Command(engineProgramPath, args...).Output(); err != nil {
		return sta.StatusAs(pipelineInfo, sta.Unknown), err
	} else {
		fmt.Println(string(out))
		return sta.StatusAs(pipelineInfo, sta.Up), nil
	}
}
