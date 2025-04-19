package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
	"github.com/alexandreLITHAUD/Own-Git/internal/utils"
)

func TestGetAllFiles(t *testing.T) {
	// Test with a valid path
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)
	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	// Create some files and directories
	subDir := filepath.Join(tempDir, "subdir")
	nestedDir := filepath.Join(subDir, "nested")
	skipDir := filepath.Join(tempDir, ".git")

	_ = os.MkdirAll(nestedDir, 0755)
	_ = os.MkdirAll(skipDir, 0755)

	file1 := filepath.Join(tempDir, "file1.txt")
	file2 := filepath.Join(subDir, "file2.txt")
	file3 := filepath.Join(nestedDir, "file3.txt")
	skipFile := filepath.Join(skipDir, "skip.txt")

	os.WriteFile(file1, []byte("file1"), 0644)
	os.WriteFile(file2, []byte("file2"), 0644)
	os.WriteFile(file3, []byte("file3"), 0644)
	os.WriteFile(skipFile, []byte("skip"), 0644)

	files, err := paths.GetAllFiles(tempDir)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expected := []string{file1, file2, file3}
	if len(files) != len(expected) {
		t.Errorf("Expected %d files, got %d", len(expected), len(files))
	}

	for _, exp := range expected {
		found := false
		for _, f := range files {
			if f == exp {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected file %s not found in result", exp)
		}
	}
}
