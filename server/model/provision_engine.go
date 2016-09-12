package model

import (
	"fmt"
	"github.com/open-devops/pipeline-daemon/server/types"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"io"
	"os"
	"os/exec"
)

func CreateProvision(pipelineInfo *types.PipelineInfo) error {
	// Engine program path
	engineParentPath := utl.GetEngineParentPath(pipelineInfo)

	// Engine program name
	engineProgramPath := utl.GetEngineProgramPath(pipelineInfo)

	// Remove existing environment
	existed, err := exists(engineParentPath)
	if err != nil {
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
	if err := copy_folder(utl.GetEngineTemplatePath(), engineParentPath); err != nil {
		fmt.Println(err)
		return err
	}

	// Make engine program executable
	if err := os.Chmod(engineProgramPath, os.ModePerm); err != nil {
		fmt.Println(err)
		return err
	}

	// Start provision process
	args := []string{"init"}
	if out, err := exec.Command(engineProgramPath, args...).Output(); err != nil {
		fmt.Println(string(out))
		fmt.Println(err)
		return err
	} else {
		fmt.Println(out)
		return nil
	}

	return nil
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}

// Copy folder recursively
func copy_folder(source string, dest string) (err error) {
	// Fetch folder contents list
	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)

	// Create destination folder if not existed
	if existed, err := exists(dest); existed == false && err == nil {
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	// Copy folder recursively
	for _, obj := range objects {
		// source file full path
		sourceFilePointer := source + "/" + obj.Name()
		// target file full path
		destinationFilePointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = copy_folder(sourceFilePointer, destinationFilePointer)
			if err != nil {
				return err
			}
		} else {
			err = copy_file(sourceFilePointer, destinationFilePointer)
			if err != nil {
				return err
			}
		}

	}
	return
}

func copy_file(source string, dest string) (err error) {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourceFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err == nil {
		sourceInfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceInfo.Mode())
		}
	}

	return
}
