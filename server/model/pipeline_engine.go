package model

import (
	"fmt"
	"github.com/open-devops/pipeline-daemon/server/types"
	sta "github.com/open-devops/pipeline-daemon/server/types/status"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os"
	"os/exec"
)

func StartPipeline(pipelineInfo *types.PipelineInfo, capability string) (*types.PipelineStatus, error) {
	// Engine program path
	engineParentPath := utl.GetEngineParentPath(pipelineInfo)

	// Engine program name
	engineProgramPath := utl.GetEngineProgramPath(pipelineInfo)

	// Change current working folder to engine path
	if err := os.Chdir(engineParentPath); err != nil {
		return sta.StatusAs(pipelineInfo, sta.Unknown), err
	}

	// Engine program parameter
	args := []string{"up", "-d", capability}

	// Dedicate the status fetch job to pipeline engine
	if out, err := exec.Command(engineProgramPath, args...).Output(); err != nil {
		return sta.StatusAs(pipelineInfo, sta.Unknown), err
	} else {
		fmt.Println(string(out))
		return FetchPipelineStatus(pipelineInfo)
	}
}
