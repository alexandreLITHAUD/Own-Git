package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"own/internal/types"
)

// TODO Find an efficent way to write only the needed entries instead of the whole index
// TODO Add a function to remove an entry from the index without rewriting the whole index
// TODO Boyer–Moore string-search algorithm ???

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

func writeEntryToIndex(entries []types.IndexEntry) error {

	currentIndex, err := parseIndex()
	if err != nil {
		return err
	}

	// TODO
	return nil
}
