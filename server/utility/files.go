package utility

import (
	"io"
	"os"
)

// exists returns whether the given file or directory exists or not
func Exists(path string) (bool, error) {
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
func CopyFolder(source string, dest string) (err error) {
	// Fetch folder contents list
	directory, _ := os.Open(source)
	objects, err := directory.Readdir(-1)

	// Create destination folder if not existed
	if existed, err := Exists(dest); existed == false && err == nil {
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
			err = CopyFolder(sourceFilePointer, destinationFilePointer)
			if err != nil {
				return err
			}
		} else {
			err = CopyFile(sourceFilePointer, destinationFilePointer)
			if err != nil {
				return err
			}
		}

	}
	return
}

func CopyFile(source string, dest string) (err error) {
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
