package utils

import (
	"encoding/json"
	"fmt"
	"os"
	"own/internal/types"
	"path/filepath"
)

const (
	added     uint8 = iota // IF IN INDEX BUT NOT IN OBJECT
	removed                // IF NOT IN INDEX BUT IN OBJECT
	modified               // IF IN INDEX AND IN OBJECT WITH SAME NAME
	renamed                // IF IN INDEX AND IN OBJECT BUT WITH DIFFERENT NAME
	untracked              // IF NOT IN INDEX AND NOT IN OBJECT
	ignored                // IF NOT IN INDEX AND NOT IN OBJECT AND IGNORED
	unknown                // ERROR CASE
)

func GetFileStatusString(status uint8) string {
	switch status {
	case added:
		return "added"
	case removed:
		return "removed"
	case modified:
		return "modified"
	case renamed:
		return "renamed"
	case untracked:
		return "untracked"
	case ignored:
		return "ignored"
	default:
		return "unknown"
	}
}

func GetObjectFile(hash string) (types.WorktreeEntry, error) {

	path := filepath.Join(".", ".own-git", "objects", hash[:2], hash[2:])
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return types.WorktreeEntry{}, err
	}

	file, err := os.ReadFile(path)
	if err != nil {
		return types.WorktreeEntry{}, err
	}

	var fileEntry types.WorktreeEntry
	err = json.Unmarshal(file, &fileEntry)
	if err != nil {
		return types.WorktreeEntry{}, err
	}

	return fileEntry, nil
}

// GetFileStatus determines the status of a file in the index and object store.
//
// If the index does not exist, the function returns unknown and an error.
// The function then parses the index, creates an IndexEntry for the file using
// FilePathtoIndexEntry, and then compares the hash of the file entry with the
// hash of the object file in the object store using GetObjectFile.
//
// The function returns the status of the file as a uint8 based on the following
// cases:
//
// - added:      IF IN INDEX BUT NOT IN OBJECT
// - removed:    IF NOT IN INDEX BUT IN OBJECT
// - modified:   IF IN INDEX AND IN OBJECT WITH SAME NAME
// - renamed:    IF IN INDEX AND IN OBJECT BUT WITH DIFFERENT NAME
// - untracked:  IF NOT IN INDEX AND NOT IN OBJECT
// - ignored:    IF NOT IN INDEX AND NOT IN OBJECT AND IGNORED
// - unknown:    ERROR CASE
// TODO DEAL WITH IGNORED AND CONFLICTED IN THE FUTURE
func GetFileStatus(path string) (uint8, error) {

	var isInIndex bool = false
	var isInObject bool = true

	if !IsIndex() {
		return unknown, fmt.Errorf("index does not exist")
	}

	currentIndexEntries, err := ParseIndex()
	if err != nil {
		return unknown, err
	}

	fileEntry, err := FilePathtoIndexEntry(path)
	if err != nil {
		return unknown, err
	}

	for _, entry := range currentIndexEntries {
		if entry.Hash == fileEntry.Hash && entry.Path == fileEntry.Path {
			isInIndex = true
			break
		}
	}

	objectFile, err := GetObjectFile(fileEntry.Hash)
	if err != nil {
		isInObject = false
	}

	if isInIndex && isInObject {
		if fileEntry.Path != objectFile.Path {
			return renamed, nil
		} else {
			return modified, nil
		}
	}
	if isInIndex && !isInObject {
		return added, nil
	}
	if !isInIndex && isInObject {
		return removed, nil
	}
	if !isInIndex && !isInObject {
		return untracked, nil
	}
	return unknown, nil
}
