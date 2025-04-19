package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
	"github.com/alexandreLITHAUD/Own-Git/internal/utils"
)

func TestIsOwnFolder(t *testing.T) {
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)

	if utils.IsOwnFolder() {
		t.Errorf("Expected IsOwnFolder to return false when folder doesn't exist")
	}

	err := os.MkdirAll(filepath.Join(tempDir, ".own-git"), os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create test folder: %v", err)
	}

	if !utils.IsOwnFolder() {
		t.Errorf("Expected IsOwnFolder to return true when folder exists")
	}
}

func TestCreateOwnFolder(t *testing.T) {
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)

	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Errorf("Expected CreateOwnFolder to succeed, got error: %v", err)
	}

	if !utils.IsOwnFolder() {
		t.Errorf("Expected IsOwnFolder to return true after creation")
	}

	// Check if key files/folders exist
	expectedPaths := []string{
		filepath.Join(tempDir, ".own-git", "objects"),
		filepath.Join(tempDir, ".own-git", "refs"),
		filepath.Join(tempDir, ".own-git", "HEAD"),
		filepath.Join(tempDir, ".own-git", "config"),
		filepath.Join(tempDir, ".own-git", "index"),
	}

	for _, p := range expectedPaths {
		if _, err := os.Stat(p); os.IsNotExist(err) {
			t.Errorf("Expected path to exist: %s", p)
		}
	}
}
