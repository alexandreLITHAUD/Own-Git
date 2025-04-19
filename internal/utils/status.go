package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
	"github.com/alexandreLITHAUD/Own-Git/internal/types"
)

func GetFileStatusString(status types.FileStatus) (string, types.Color) {
	switch status {
	case types.Added:
		return "added", types.Green
	case types.Removed:
		return "removed", types.Cyan
	case types.Modified:
		return "modified", types.Yellow
	case types.Renamed:
		return "renamed", types.Blue
	case types.Untracked:
		return "untracked", types.Red
	case types.Ignored:
		return "ignored", types.Purple
	default:
		return "unknown", types.NoColor
	}
}

func GetObjectFile(hash string) (types.WorktreeEntry, error) {

	path := paths.GetObjectFilePath(hash)
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
func GetFileStatus(path string) (types.FileStatusStruct, error) {

	var isInIndex bool = false
	var isInObject bool = true
	var fileStatusStruct types.FileStatusStruct = types.FileStatusStruct{
		Path:   path,
		Status: types.Unknown,
	}

	if !IsIndex() {
		return fileStatusStruct, fmt.Errorf("index does not exist")
	}

	currentIndexEntries, err := ParseIndex()
	if err != nil {
		return fileStatusStruct, err
	}

	fileEntry, err := FilePathtoIndexEntry(path)
	if err != nil {
		return fileStatusStruct, err
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
			fileStatusStruct.Status = types.Renamed
			return fileStatusStruct, nil
		} else {
			fileStatusStruct.Status = types.Modified
			return fileStatusStruct, nil
		}
	}
	if isInIndex && !isInObject {
		fileStatusStruct.Status = types.Added
		return fileStatusStruct, nil
	}
	if !isInIndex && isInObject {
		fileStatusStruct.Status = types.Removed
		return fileStatusStruct, nil
	}
	if !isInIndex && !isInObject {
		fileStatusStruct.Status = types.Untracked
		return fileStatusStruct, nil
	}
	return fileStatusStruct, nil
}
