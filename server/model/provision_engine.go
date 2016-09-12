package model

import (
	"github.com/open-devops/pipeline-daemon/server/types"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os"
	"fmt"
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
			return err
		} else {
			fmt.Println(engineParentPath + " Removed!")
		}

	}

	// Create new environment
	if err := os.Mkdir(engineParentPath, os.ModePerm); err != err {
		return err
	} else {
		fmt.Println(engineParentPath + " Created!")
		return nil
	}
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}
