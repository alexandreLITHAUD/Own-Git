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

// GetIndexFilePath returns the path of the index file. This could be improve to check
// if the .own-git folder exists in parents.
//
// The path is the path of the current working directory plus the name of the
// .own-git folder plus the name of the index file.
func GetIndexFilePath() string {
	return filepath.Join(".", ".own-git", "index")
}

// IsIndex checks if the index exists.
//
// The function first checks if the .own-git folder exists. If it does, it then
// checks if the index file exists. If it does, the function returns true;
// otherwise, it returns false.
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

// IsIndexEmpty checks if the index is empty.
//
// The function first checks if the index exists. If it does not, the function
// returns true and a nil error. If the index does exist, the function then
// checks if the file size is 0. If it is, the function returns true and a nil
// error; otherwise, the function returns false and a nil error.
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

// ParseIndex reads the index file and parses it into a slice of IndexEntry.
//
// If the index does not exist, the function returns nil and an
// error. If the index exists but is empty, the function returns an empty slice
// and a nil error. If the index exists and is not empty, the function reads the
// file, unmarshals the JSON data into a slice of IndexEntry, and returns the
// slice and a nil error. If an error occurs while reading the file or
// unmarshalling the JSON data, the function returns nil and the error.
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

// MergeIndexEntries merges two slices of IndexEntry.
//
// This function takes two slices of IndexEntry: the current index and new entries.
// It creates a map of the index entries using their paths as keys. New entries
// override existing entries with the same path. The function returns a slice of
// merged index entries, effectively combining the current index with the new entries.
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

// WriteEntryToIndex writes a slice of IndexEntry to the index.
//
// This function takes a slice of IndexEntry and writes it to the index file.
// It first parses the current index, merges the new entries with the current index,
// and then writes the merged index to disk.
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

// RemoveEntryFromIndex removes an entry from the index by its file path.
//
// The function first parses the current index to get a list of IndexEntry.
// It then iterates over the entries and removes the entry that matches the
// provided path. After the entry is removed, it marshals the updated list
// back to JSON and writes it to the index file. If an error occurs at any
// point, the function returns the error.
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

// FilePathtoIndexEntry takes a file path and returns an IndexEntry.
//
// It first determines the file mode, either 100644 or 100755 depending on
// if the file is executable. It then computes the SHA1 of the file and
// creates an IndexEntry with the file path, mode, and hash. If an error
// occurs while computing the SHA1, the function returns an error.
func FilePathtoIndexEntry(path string) (types.IndexEntry, error) {

	var mode string = "100644"

	isexec, err := IsExecutable(path)
	// CHECK ERROR ??
	if isexec && err == nil {
		mode = "100755"
	}

	hash, err := GetFileSHA1(path)
	if err != nil {
		return types.IndexEntry{}, fmt.Errorf("error getting file SHA1: %w", err)
	}

	return types.CreateIndexEntry(path, mode, hash), nil
}
