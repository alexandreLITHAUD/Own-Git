package scutils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateOwnFolder(initialBranchName string, configFile string) error {

	var path string = filepath.Join(".", ".own-git")

	if _, err := os.Stat(path); os.IsExist(err) {
		return nil
	}

	err := os.MkdirAll(path, os.ModePerm)

	if err != nil {
		return fmt.Errorf("error creating folder: %v", err)
	}

	objectFolderPath := filepath.Join(path, "objects")
	err = os.MkdirAll(objectFolderPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating objects folder: %v", err)
	}
	refsFolderPath := filepath.Join(path, "refs")
	err = os.MkdirAll(refsFolderPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating refs folder: %v", err)
	}

	var configBytes []byte = []byte("")

	if configFile != "" {
		if valid := IsConfFileValid(configFile); valid {
			return fmt.Errorf("error parsing config file: %v", err)
		}

		configBytes, err = os.ReadFile(configFile)
		if err != nil {
			return fmt.Errorf("error reading config file: %v", err)
		}
	}

	os.WriteFile(filepath.Join(path, "HEAD"), []byte("ref: refs/heads/"+initialBranchName+"\n"), os.ModePerm)
	os.WriteFile(filepath.Join(path, "config"), configBytes, os.ModePerm)
	os.WriteFile(filepath.Join(path, "index"), []byte(""), os.ModePerm)

	return nil
}
