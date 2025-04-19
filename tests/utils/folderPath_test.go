package tests

import (
	"path/filepath"
	"testing"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
	"github.com/alexandreLITHAUD/Own-Git/internal/utils"
)

func TestGetOwnGitFolderPath(t *testing.T) {
	// Test with a valid path
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)

	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	expectedPath := filepath.Join(tempDir, ".own-git")
	if paths.GetOwnGitFolderPath() != expectedPath {
		t.Errorf("Expected .own-git folder path to be %s, got %s", expectedPath, paths.GetOwnGitFolderPath())
	}
}

func TestGetIndexFilePath(t *testing.T) {
	// Test with a valid path
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)

	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	expectedPath := filepath.Join(tempDir, ".own-git", "index")
	if paths.GetIndexFilePath() != expectedPath {
		t.Errorf("Expected index file path to be %s, got %s", expectedPath, paths.GetIndexFilePath())
	}
}

func TestGetObjectFilePath(t *testing.T) {
	// Test with a valid path
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)

	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	hash := "1234567890abcdef1234567890abcdef12345678"
	expectedPath := filepath.Join(tempDir, ".own-git", "objects", hash[:2], hash[2:])
	if paths.GetObjectFilePath(hash) != expectedPath {
		t.Errorf("Expected object file path to be %s, got %s", expectedPath, paths.GetObjectFilePath(hash))
	}
}

func TestGetObjectFolderPath(t *testing.T) {
	// Test with a valid path
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)

	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	expectedPath := filepath.Join(tempDir, ".own-git", "objects")
	if paths.GetObjectFolderPath() != expectedPath {
		t.Errorf("Expected object folder path to be %s, got %s", expectedPath, paths.GetObjectFolderPath())
	}
}

func TestGetRefsFolderPath(t *testing.T) {
	// Test with a valid path
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)

	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	expectedPath := filepath.Join(tempDir, ".own-git", "refs")
	if paths.GetRefsFolderPath() != expectedPath {
		t.Errorf("Expected refs folder path to be %s, got %s", expectedPath, paths.GetRefsFolderPath())
	}
}

func TestGetHeadFilePath(t *testing.T) {
	// Test with a valid path
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)

	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}
	expectedPath := filepath.Join(tempDir, ".own-git", "HEAD")
	if paths.GetHeadFilePath() != expectedPath {
		t.Errorf("Expected HEAD file path to be %s, got %s", expectedPath, paths.GetHeadFilePath())
	}
}

func TestGetConfigFilePath(t *testing.T) {
	// Test with a valid path
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)

	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}
	expectedPath := filepath.Join(tempDir, ".own-git", "config")
	if paths.GetConfigFilePath() != expectedPath {
		t.Errorf("Expected config file path to be %s, got %s", expectedPath, paths.GetConfigFilePath())
	}
}
