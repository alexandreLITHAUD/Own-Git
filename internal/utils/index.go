package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"own/internal/types"
)

// TODO Find an efficent way to write only the needed entries instead of the whole index
// TODO Add a function to remove an entry from the index without rewriting the whole index
// TODO Boyer–Moore string-search algorithm ???

func GetIndexFilePath() string {
	return filepath.Join(".", ".own-git", "index")
}

func IsIndex() bool {
	if !IsOwnFolder() {
		return false
	}

	path := GetIndexFilePath()
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}

func IsIndexEmpty() (bool, error) {
	if !IsIndex() {
		return true, nil
	}
	path := GetIndexFilePath()
	fileinfo, err := os.Stat(path)
	if err != nil {
		return false, err
	}

	if fileinfo.Size() == 0 {
		return true, nil
	}

	return false, nil
}

func ParseIndex() ([]types.IndexEntry, error) {
	isempty, err := IsIndexEmpty()
	if err != nil {
		return nil, err
	}

	if isempty {
		return make([]types.IndexEntry, 0), nil
	}

	path := GetIndexFilePath()
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

func MergeIndexEntries(currentIndex []types.IndexEntry, newEntries []types.IndexEntry) []types.IndexEntry {
	// TODO Add a function to remove an entry from the index without rewriting the whole index

	indexMap := make(map[string]types.IndexEntry)

	for _, entry := range currentIndex {
		indexMap[entry.Path] = entry
	}
	for _, entry := range newEntries {
		indexMap[entry.Path] = entry // new entries override old
	}

	merged := make([]types.IndexEntry, 0, len(indexMap))
	for _, entry := range indexMap {
		merged = append(merged, entry)
	}
	return merged
}

func WriteEntryToIndex(entries []types.IndexEntry) error {

	currentIndex, err := ParseIndex()
	if err != nil {
		return err
	}

	mergedIndexEntries := MergeIndexEntries(currentIndex, entries)

	jsonData, err := json.Marshal(mergedIndexEntries)
	if err != nil {
		return err
	}

	path := GetIndexFilePath()
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func RemoveEntryFromIndex(path string) error {
	currentIndex, err := ParseIndex()
	if err != nil {
		return err
	}

	// Remove the entry from the index
	for i, entry := range currentIndex {
		if entry.Path == path {
			currentIndex = append(currentIndex[:i], currentIndex[i+1:]...)
			break
		}
	}

	jsonData, err := json.Marshal(currentIndex)
	if err != nil {
		return err
	}

	path = GetIndexFilePath()
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func FilePathtoIndexEntry(path string) (types.IndexEntry, error) {

	var mode string = "100644"

	isexec, err := IsExecutable(path)
	if isexec && err == nil {
		mode = "100755"
	}

	hash, err := GetFileSHA1(path)
	if err != nil {
		return types.IndexEntry{}, fmt.Errorf("error getting file SHA1: %w", err)
	}

	return types.CreateIndexEntry(path, mode, hash), nil
}
