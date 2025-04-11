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

func IsIndex() bool {
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
	if !IsIndex() {
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

func ParseIndex() ([]types.IndexEntry, error) {
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

func MergeIndexEntries(currentIndex []types.IndexEntry, newEntries []types.IndexEntry) []types.IndexEntry {
	// TODO Add a function to remove an entry from the index without rewriting the whole index
	// TODO find a better way to do this

	mergedIndex := make([]types.IndexEntry, 0)

	for _, entry := range currentIndex {
		// Check if the entry already exists in the new entries
		exist := false
		for _, newEntry := range newEntries {
			if entry.Path == newEntry.Path {
				mergedIndex = append(mergedIndex, newEntry)
				exist = true
				break
			}
		}
		if !exist {
			mergedIndex = append(mergedIndex, entry)
		}
	}

	// Add new entries that are not in the current index
	for _, newEntry := range newEntries {
		exist := false
		for _, entry := range currentIndex {
			if newEntry.Path == entry.Path {
				exist = true
				break
			}
		}
		if !exist {
			mergedIndex = append(mergedIndex, newEntry)
		}
	}

	return mergedIndex
}

func WriteEntryToIndex(entries []types.IndexEntry) error {

	currentIndex, err := ParseIndex()
	if err != nil {
		return err
	}

	mergedIndexEntries := MergeIndexEntries(currentIndex, entries)

	jsonData, err := json.Marshal(mergedIndexEntries)
	if err != nil {
		return nil
	}

	path := filepath.Join(".", ".own-git", "index")
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
		return nil
	}

	path = filepath.Join(".", ".own-git", "index")
	err = os.WriteFile(path, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
