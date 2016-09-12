package model

import (
	"github.com/open-devops/pipeline-daemon/server/types"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os"
	"fmt"
	"os/exec"
)

func CreateProvision(pipelineInfo *types.PipelineInfo) error {
	// Engine program path
	engineParentPath := utl.GetEngineParentPath(pipelineInfo)

	// Remove existing environment
    existed, err := exists(engineParentPath);
	if (err != nil) {
		return err
	}
	if existed {
		if err := os.RemoveAll(engineParentPath); err != nil {
			fmt.Println(err)
			return err
		}
	}

	// Create new environment
	if err := os.MkdirAll(engineParentPath, os.ModePerm); err != nil {
		return err
	} else {
		if err := os.Chdir(engineParentPath); err != nil {
			return err
		}
	}

	// Make engine program ready
	cmd := "cp"
	src := utl.GetEngineTemplatePath()
	args := []string {"-R", src ,engineParentPath}
	if err := exec.Command(cmd, args...).Run() ; err != nil {
		return err
	}

	return nil
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}
