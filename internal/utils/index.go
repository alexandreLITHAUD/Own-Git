package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"own/internal/types"
)

func isIndex() bool {
	if !IsOwnFolder() {
		return false
	}

	path := filepath.Join(".", ".own-git", "index")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsIndexEmpty() (bool, error) {
	if !isIndex() {
		return true, nil
	}
	path := filepath.Join(".", ".own-git", "index")
	fileinfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if fileinfo.Size() == 0 {
		return true, nil
	}

	return false, nil
}

func parseIndex() ([]types.IndexEntry, error) {
	isempty, err := IsIndexEmpty()
	if err != nil {
		return nil, err
	}

	if isempty {
		return make([]types.IndexEntry, 0), nil
	}

	path := filepath.Join(".", ".own-git", "index")
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var entries []types.IndexEntry
	err = json.Unmarshal(file, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
