package model

import (
	"github.com/open-devops/pipeline-daemon/server/types"
	sta "github.com/open-devops/pipeline-daemon/server/types/status"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os"
	"os/exec"
)

const (
	ACTION_START   = "up"
	ACTION_RESTART = "restart"
	ACTION_STOP    = "stop"
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
	actionMapping := map[string]string{
		"start":   ACTION_START,
		"restart": ACTION_RESTART,
		"stop":    ACTION_STOP,
	}

	args := []string{actionMapping[action], capability}

	// Dedicate the status fetch job to pipeline engine
	if _, err := exec.Command(engineProgramPath, args...).Output(); err != nil {
		return sta.StatusAs(pipelineInfo, sta.Unknown), err
	} else {
		switch action {
		default:
			return sta.StatusAs(pipelineInfo, sta.Unknown), nil
		case "start", "restart":
			return sta.StatusAs(pipelineInfo, sta.Up), nil
		case "stop":
			return sta.StatusAs(pipelineInfo, sta.Down), nil
		}
	}
}
