package tests

import (
	"testing"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
	"github.com/alexandreLITHAUD/Own-Git/internal/types"
	"github.com/alexandreLITHAUD/Own-Git/internal/utils"
)

func BenchmarkParseIndex(b *testing.B) {
	tempDir := b.TempDir()
	paths.SetBasePath(tempDir)
	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		b.Fatalf("CreateOwnFolder failed: %v", err)
	}

	// Prepare index content
	entries := []types.IndexEntry{
		types.CreateIndexEntry("file1.txt", "100644", "abc123"),
		types.CreateIndexEntry("file2.txt", "100644", "def456"),
	}
	err = utils.WriteEntryToIndex(entries)
	if err != nil {
		b.Fatalf("WriteEntryToIndex failed: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := utils.ParseIndex()
		if err != nil {
			b.Fatalf("ParseIndex failed: %v", err)
		}
	}
}

func BenchmarkWriteEntryToIndex(b *testing.B) {
	tempDir := b.TempDir()
	paths.SetBasePath(tempDir)
	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		b.Fatalf("CreateOwnFolder failed: %v", err)
	}

	entries := []types.IndexEntry{
		types.CreateIndexEntry("file1.txt", "100644", "abc123"),
		types.CreateIndexEntry("file2.txt", "100644", "def456"),
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := utils.WriteEntryToIndex(entries)
		if err != nil {
			b.Fatalf("WriteEntryToIndex failed: %v", err)
		}
	}
}

func BenchmarkRemoveEntryFromIndex(b *testing.B) {
	tempDir := b.TempDir()
	paths.SetBasePath(tempDir)
	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		b.Fatalf("CreateOwnFolder failed: %v", err)
	}

	entries := []types.IndexEntry{
		types.CreateIndexEntry("file1.txt", "100644", "abc123"),
		types.CreateIndexEntry("file2.txt", "100644", "def456"),
	}
	err = utils.WriteEntryToIndex(entries)
	if err != nil {
		b.Fatalf("WriteEntryToIndex failed: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err := utils.RemoveEntryFromIndex("file1.txt")
		if err != nil {
			b.Fatalf("RemoveEntryFromIndex failed: %v", err)
		}
		// Re-add the entry for the next iteration
		_ = utils.WriteEntryToIndex(entries)
	}
}
