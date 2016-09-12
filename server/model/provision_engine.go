package model

import (
	"github.com/open-devops/pipeline-daemon/server/types"
	utl "github.com/open-devops/pipeline-daemon/server/utility"
	"os"
	"fmt"
	"io"
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
	if err := copy_folder(utl.GetEngineTemplatePath(), engineParentPath); err != nil {
		return err;
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

func copy_folder(source string, dest string) (err error) {
	// Get folder contents
	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)

	// Copy folder contents
	for _, obj := range objects {

		sourceFilePointer := source + "/" + obj.Name()

		destinationFilePointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			err = copy_folder(sourceFilePointer, destinationFilePointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			err = copy_file(sourceFilePointer, destinationFilePointer)
			if err != nil {
				fmt.Println(err)
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

