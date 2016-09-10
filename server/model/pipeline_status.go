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
	enginePath := utl.GetEnginePath(pipelineInfo)

	// Engine program name
	engineName := utl.GetEngineName()

	// Engine program parameter
	args := []string{"status"}

	// Pipeline status information
	var status *types.PipelineStatus

	fmt.Println("engine = " + enginePath + "/" + engineName)

	// Fetch pipeline status information using engine program
	if err := os.Chdir(enginePath); err != err {
		return sta.StatusAs(pipelineInfo, sta.Unknown)
	}

	// Dedicate the status fetch job to pipeline engine
	if out, err := exec.Command(engineName, args...).Output(); err != nil {
		return sta.StatusAs(pipelineInfo, sta.Unknown)
	} else {
		fmt.Println(string(out))
		return sta.StatusAs(pipelineInfo, sta.Up)
	}

	return status
}
