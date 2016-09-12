package model

import (
	"github.com/open-devops/pipeline-daemon/server/types"
	sta "github.com/open-devops/pipeline-daemon/server/types/status"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os"
	"os/exec"
	"fmt"
)

func HandlePipelineAction(pipelineInfo *types.PipelineInfo,
	action string,
	capability string) (*types.PipelineStatus, error) {

	// Engine program path
	engineParentPath := utl.GetEngineParentPath(pipelineInfo)

	// Engine program name
	engineProgramPath := utl.GetEngineProgramPath(pipelineInfo)

	// Change current working folder to engine path
	if err := os.Chdir(engineParentPath); err != nil {
		return sta.StatusAs(pipelineInfo, sta.Unknown), err
	}

	// Engine program parameter
	ActionMapper := map[string]string{
		"start":   "up",
		"restart": "restart",
		"stop":    "stop",
	}
	args := []string{ActionMapper[action], capability}
	fmt.Println("action = " + action + ";args[0] = " + args[0] + ";args[1] = " + args[1] )

	// Dedicate the status fetch job to pipeline engine
	if _, err := exec.Command(engineProgramPath, args...).Output(); err != nil {
		return sta.StatusAs(pipelineInfo, sta.Unknown), err
	} else {
		return FetchPipelineStatus(pipelineInfo)
	}
}
