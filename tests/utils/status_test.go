package tests

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
	"github.com/alexandreLITHAUD/Own-Git/internal/types"
	"github.com/alexandreLITHAUD/Own-Git/internal/utils"
)

func TestGetFileStatusString(t *testing.T) {
	tests := []struct {
		status        types.FileStatus
		expected      string
		colorExpected types.Color
	}{
		{types.Added, "added", types.Green},
		{types.Removed, "removed", types.Cyan},
		{types.Modified, "modified", types.Yellow},
		{types.Renamed, "renamed", types.Blue},
		{types.Untracked, "untracked", types.Red},
		{types.Ignored, "ignored", types.Purple},
		{types.Unknown, "unknown", types.NoColor},
	}

	for _, test := range tests {
		result, color := utils.GetFileStatusString(test.status)
		if result != test.expected {
			t.Errorf("Expected %s for status %v, got %s", test.expected, test.status, result)
		}
		if color != test.colorExpected {
			t.Errorf("Expected color %s for status %v, got %s", test.colorExpected, test.status, color)
		}
	}
}

func TestGetObjectFile(t *testing.T) {

	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)
	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	hash := "1234567890abcdef1234567890abcdef12345678"

	_, err = utils.GetObjectFile(hash)

	if err == nil {
		t.Errorf("Expected error, got nil")
	}

	worktree := types.WorktreeEntry{
		Name:    "test.txt",
		Mode:    "100644",
		Type:    "blob",
		Path:    "test.txt",
		Content: []byte("test content"),
	}

	worktreeJSON, err := json.Marshal(worktree)
	if err != nil {
		t.Fatalf("Failed to marshal worktree entry: %v", err)
	}

	err = os.Mkdir(filepath.Join(paths.GetObjectFolderPath(), hash[:2]), 0755)
	if err != nil {
		t.Fatalf("Failed to create object folder: %v", err)
	}
	err = os.WriteFile(paths.GetObjectFilePath(hash), worktreeJSON, 0644)
	if err != nil {
		t.Fatalf("Failed to write object file: %v", err)
	}

	worktree, err = utils.GetObjectFile(hash)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if worktree.Name != "test.txt" {
		t.Errorf("Expected name test.txt, got %s", worktree.Name)
	}
	if worktree.Mode != "100644" {
		t.Errorf("Expected mode 100644, got %s", worktree.Mode)
	}
	if worktree.Type != "blob" {
		t.Errorf("Expected type blob, got %s", worktree.Type)
	}
	if worktree.Path != "test.txt" {
		t.Errorf("Expected path test.txt, got %s", worktree.Path)
	}
	if string(worktree.Content) != "test content" {
		t.Errorf("Expected content 'test content', got %s", string(worktree.Content))
	}
	if len(worktree.Children) != 0 {
		t.Errorf("Expected no children, got %d", len(worktree.Children))
	}
}

func TestGetFileStatus(t *testing.T) {

	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)
	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	// TEST FILE NOT EXISTING TO TEST UNKNOWN
	var result types.FileStatusStruct
	result, err = utils.GetFileStatus("test.txt")
	if err == nil {
		t.Fatalf("Expected error")
	}
	if result.Status != types.Unknown {
		t.Errorf("Expected status %v, got %v", types.Unknown, result)
	}

	// CREATE FILE TO TEST UNTRACKED
	testFilePath := filepath.Join(tempDir, "test.txt")
	err = os.WriteFile(testFilePath, []byte("test content"), 0644)
	if err != nil {
		t.Fatalf("Failed to write file: %v", err)
	}
	result, err = utils.GetFileStatus(testFilePath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Status != types.Untracked {
		t.Errorf("Expected status %v, got %v", types.Untracked, result)
	}

	// CREATE INDEX TO TEST ADDED
	var indexEntryArray []types.IndexEntry
	var indexEntry types.IndexEntry
	indexEntry, err = utils.FilePathtoIndexEntry(testFilePath)
	if err != nil {
		t.Fatalf("Failed to convert file path to index entry: %v", err)
	}
	indexEntryArray = append(indexEntryArray, indexEntry)
	err = utils.WriteEntryToIndex(indexEntryArray)
	if err != nil {
		t.Fatalf("Failed to write entry to index: %v", err)
	}
	result, err = utils.GetFileStatus(testFilePath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Status != types.Added {
		t.Errorf("Expected status %v, got %v", types.Added, result)
	}

	// REMOVE FILE TO ENTRY
	err = utils.RemoveEntryFromIndex(testFilePath)
	if err != nil {
		t.Fatalf("Failed to remove entry from index: %v", err)
	}

	// ADD FILE TO OBJECT AND REMOVE INDEX TO TEST REMOVE
	var hash string
	hash, err = utils.GetFileSHA1(testFilePath)
	if err != nil {
		t.Fatalf("Unexpected error, got %v", err)
	}
	err = os.Mkdir(filepath.Join(paths.GetObjectFolderPath(), hash[:2]), 0755)
	if err != nil {
		t.Fatalf("Failed to create object folder: %v", err)
	}
	worktreeJSON, err := json.Marshal(types.CreateBlobWorktreeEntry(testFilePath, []byte("test content")))
	if err != nil {
		t.Fatalf("Failed to marshal worktree entry: %v", err)
	}
	err = os.WriteFile(paths.GetObjectFilePath(hash), worktreeJSON, 0644)
	if err != nil {
		t.Fatalf("Failed to write object file: %v", err)
	}
	result, err = utils.GetFileStatus(testFilePath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Status != types.Removed {
		t.Errorf("Expected status %v, got %v", types.Removed, result)
	}

	// ADD FILE TO INDEX AGAIN TO TEST MODIFIED
	err = utils.WriteEntryToIndex(indexEntryArray)
	if err != nil {
		t.Fatalf("Failed to write entry to index: %v", err)
	}
	result, err = utils.GetFileStatus(testFilePath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Status != types.Modified {
		t.Errorf("Expected status %v, got %v", types.Modified, result)
	}

	// RENAME FILE TO TEST RENAMED
	newFilePath := filepath.Join(tempDir, "test_renamed.txt")
	err = os.Rename(testFilePath, newFilePath)
	if err != nil {
		t.Fatalf("Failed to rename file: %v", err)
	}
	err = utils.RemoveEntryFromIndex(testFilePath)
	if err != nil {
		t.Fatalf("Failed to remove entry from index: %v", err)
	}
	indexEntry, err = utils.FilePathtoIndexEntry(newFilePath)
	if err != nil {
		t.Fatalf("Failed to convert file path to index entry: %v", err)
	}
	var newIndexArray []types.IndexEntry
	newIndexArray = append(newIndexArray, indexEntry)
	err = utils.WriteEntryToIndex(newIndexArray)
	if err != nil {
		t.Fatalf("Failed to write entry to index: %v", err)
	}
	result, err = utils.GetFileStatus(newFilePath)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if result.Status != types.Renamed {
		t.Errorf("Expected status %v, got %v", types.Renamed, result)
	}
}
