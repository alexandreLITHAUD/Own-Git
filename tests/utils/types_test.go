package tests

import (
	"own/internal/types"
	"testing"
)

func TestCreateIndexEntry(t *testing.T) {
	path := "test/path"
	mode := "100644"
	hash := "1234567890abcdef1234567890abcdef12345678"

	entry := types.CreateIndexEntry(path, mode, hash)

	if entry.Path != path {
		t.Errorf("Expected path %s, got %s", path, entry.Path)
	}
	if entry.Mode != mode {
		t.Errorf("Expected mode %s, got %s", mode, entry.Mode)
	}
	if entry.Hash != hash {
		t.Errorf("Expected hash %s, got %s", hash, entry.Hash)
	}
}

func TestCreateBlobWorktreeEntry(t *testing.T) {
	path := "test/path"
	content := []byte("test content")

	entry := types.CreateBlobWorktreeEntry(path, content)

	if entry.Name != path {
		t.Errorf("Expected name %s, got %s", path, entry.Name)
	}
	if entry.Mode != "100644" {
		t.Errorf("Expected mode %s, got %s", "100644", entry.Mode)
	}
	if entry.Type != "blob" {
		t.Errorf("Expected type %s, got %s", "blob", entry.Type)
	}
	if entry.Path != path {
		t.Errorf("Expected path %s, got %s", path, entry.Path)
	}
	if string(entry.Content) != string(content) {
		t.Errorf("Expected content %s, got %s", string(content), string(entry.Content))
	}
}

func TestCreateTreeWorktreeEntry(t *testing.T) {
	path := "test/path"
	children := []*types.WorktreeEntry{
		{Name: "child1", Mode: "100644", Type: "blob", Path: "child1", Content: []byte("child1 content")},
		{Name: "child2", Mode: "100644", Type: "blob", Path: "child2", Content: []byte("child2 content")},
	}

	entry := types.CreateTreeWorktreeEntry(path, children)

	if entry.Name != path {
		t.Errorf("Expected name %s, got %s", path, entry.Name)
	}
	if entry.Mode != "040000" {
		t.Errorf("Expected mode %s, got %s", "040000", entry.Mode)
	}
	if entry.Type != "tree" {
		t.Errorf("Expected type %s, got %s", "tree", entry.Type)
	}
	if entry.Path != path {
		t.Errorf("Expected path %s, got %s", path, entry.Path)
	}
	if len(entry.Children) != 2 {
		t.Errorf("Expected 2 children, got %d", len(entry.Children))
	}
	if entry.Children[0].Name != "child1" {
		t.Errorf("Expected child name %s, got %s", "child1", entry.Children[0].Name)
	}
	if entry.Children[1].Name != "child2" {
		t.Errorf("Expected child name %s, got %s", "child2", entry.Children[1].Name)
	}
	if string(entry.Children[0].Content) != "child1 content" {
		t.Errorf("Expected child content %s, got %s", "child1 content", string(entry.Children[0].Content))
	}
	if string(entry.Children[1].Content) != "child2 content" {
		t.Errorf("Expected child content %s, got %s", "child2 content", string(entry.Children[1].Content))
	}
	if entry.Children[0].Mode != "100644" {
		t.Errorf("Expected child mode %s, got %s", "100644", entry.Children[0].Mode)
	}
	if entry.Children[1].Mode != "100644" {
		t.Errorf("Expected child mode %s, got %s", "100644", entry.Children[1].Mode)
	}
	if entry.Children[0].Type != "blob" {
		t.Errorf("Expected child type %s, got %s", "blob", entry.Children[0].Type)
	}
	if entry.Children[1].Type != "blob" {
		t.Errorf("Expected child type %s, got %s", "blob", entry.Children[1].Type)
	}
	if entry.Children[0].Path != "child1" {
		t.Errorf("Expected child path %s, got %s", "child1", entry.Children[0].Path)
	}
	if entry.Children[1].Path != "child2" {
		t.Errorf("Expected child path %s, got %s", "child2", entry.Children[1].Path)
	}
}
