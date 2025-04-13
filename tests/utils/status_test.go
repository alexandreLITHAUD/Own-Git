package tests

import (
	"encoding/json"
	"os"
	"own/internal/paths"
	"own/internal/types"
	"own/internal/utils"
	"path/filepath"
	"testing"
)

func TestGetFileStatusString(t *testing.T) {
	tests := []struct {
		status   uint8
		expected string
	}{
		{utils.Added, "added"},
		{utils.Removed, "removed"},
		{utils.Modified, "modified"},
		{utils.Renamed, "renamed"},
		{utils.Untracked, "untracked"},
		{utils.Ignored, "ignored"},
		{utils.Unknown, "unknown"},
	}

	for _, test := range tests {
		result := utils.GetFileStatusString(test.status)
		if result != test.expected {
			t.Errorf("Expected %s for status %v, got %s", test.expected, test.status, result)
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
	var result uint8
	result, err = utils.GetFileStatus("test.txt")
	if err == nil {
		t.Fatalf("Expected error")
	}
	if result != utils.Unknown {
		t.Errorf("Expected status %v, got %v", utils.Unknown, result)
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
	if result != utils.Untracked {
		t.Errorf("Expected status %v, got %v", utils.Untracked, result)
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
	if result != utils.Added {
		t.Errorf("Expected status %v, got %v", utils.Added, result)
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
	if result != utils.Removed {
		t.Errorf("Expected status %v, got %v", utils.Removed, result)
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
	if result != utils.Modified {
		t.Errorf("Expected status %v, got %v", utils.Modified, result)
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
	if result != utils.Renamed {
		t.Errorf("Expected status %v, got %v", utils.Renamed, result)
	}
}
