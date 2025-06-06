package utils

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
)

func IsOwnFolder() bool {
	path := paths.GetOwnGitFolderPath()

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateOwnFolder(initialBranchName string, configFile string) error {

	if IsOwnFolder() {
		return fmt.Errorf("error: .own-git folder already exists")
	}
	path := paths.GetOwnGitFolderPath()

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

	configBytes := []byte("")

	if configFile != "" {
		if valid := IsConfFileValid(configFile); !valid {
			return fmt.Errorf("error parsing config file")
		}

		configBytes, err = os.ReadFile(configFile)
		if err != nil {
			return fmt.Errorf("error reading config file: %v", err)
		}
	}

	err = os.WriteFile(filepath.Join(path, "HEAD"), []byte("ref: refs/heads/"+initialBranchName+"\n"), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating HEAD file: %v", err)
	}
	err = os.WriteFile(filepath.Join(path, "config"), configBytes, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating config file: %v", err)
	}
	err = os.WriteFile(filepath.Join(path, "index"), []byte(""), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating index file: %v", err)
	}

	return nil
}
