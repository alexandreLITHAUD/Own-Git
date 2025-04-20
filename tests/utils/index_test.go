package tests

import (
	"os"
	"sort"
	"testing"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
	"github.com/alexandreLITHAUD/Own-Git/internal/types"
	"github.com/alexandreLITHAUD/Own-Git/internal/utils"
)

func TestIsIndex(t *testing.T) {

	if utils.IsIndex() {
		t.Errorf("Expected IsIndex to return false when index doesn't exist")
	}

	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)
	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	if !utils.IsIndex() {
		t.Errorf("Expected IsIndex to return true when index exists")
	}
}

func TestIsIndexEmpty(t *testing.T) {

	isEmpty, err := utils.IsIndexEmpty()

	if !isEmpty {
		t.Errorf("Expected IsIndexEmpty to return true when index not created")
	}

	if err != nil {
		t.Fatalf("IsIndexEmpty failed: %v", err)
	}

	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)
	err = utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	isEmpty, err = utils.IsIndexEmpty()

	if err != nil {
		t.Fatalf("IsIndexEmpty failed: %v", err)
	}

	if !isEmpty {
		t.Errorf("Expected IsIndexEmpty to return true when index is empty")
	}

	err = os.WriteFile(paths.GetIndexFilePath(), []byte("test"), os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to write index file: %v", err)
	}
	isEmpty, err = utils.IsIndexEmpty()

	if err != nil {
		t.Fatalf("IsIndexEmpty failed: %v", err)
	}

	if isEmpty {
		t.Errorf("Expected IsIndexEmpty to return false when index is not empty")
	}
}

func TestParseIndex(t *testing.T) {
	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)
	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	entries, err := utils.ParseIndex()
	if err != nil {
		t.Fatalf("ParseIndex failed: %v", err)
	}

	if len(entries) != 0 {
		t.Errorf("Expected ParseIndex to return empty slice when index is empty")
	}

	// Create test index entries
	indexContent := `[{"path":"file1.txt","mode":"100644","hash":"abc123"},{"path":"file2.txt","mode":"100644","hash":"def456"}]`
	err = os.WriteFile(paths.GetIndexFilePath(), []byte(indexContent), os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to write index file: %v", err)
	}

	entries, err = utils.ParseIndex()
	if err != nil {
		t.Fatalf("ParseIndex failed: %v", err)
	}

	// Sort entries for comparison
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Path < entries[j].Path
	})

	if len(entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(entries))
	}

	if entries[0].Path != "file1.txt" || entries[1].Path != "file2.txt" {
		t.Errorf("Unexpected index entries: %+v", entries)
	}

	if entries[0].Mode != "100644" || entries[1].Mode != "100644" {
		t.Errorf("Unexpected index modes: %s, %s", entries[0].Mode, entries[1].Mode)
	}

	if entries[0].Hash != "abc123" || entries[1].Hash != "def456" {
		t.Errorf("Unexpected index hashes: %s, %s", entries[0].Hash, entries[1].Hash)
	}
}

func TestMergeIndexEntries(t *testing.T) {

	entriesEmpty1 := []types.IndexEntry{}
	entriesEmpty2 := []types.IndexEntry{}
	mergedEmpty := utils.MergeIndexEntries(entriesEmpty1, entriesEmpty2)
	if len(mergedEmpty) != 0 {
		t.Errorf("Expected merged entries to be empty, got %d", len(mergedEmpty))
	}

	entries1 := []types.IndexEntry{
		types.CreateIndexEntry("file1.txt", "100644", "abc123"),
		types.CreateIndexEntry("file2.txt", "100644", "def456"),
	}

	entries2 := []types.IndexEntry{
		types.CreateIndexEntry("file2.txt", "100644", "def789"),
		types.CreateIndexEntry("file3.txt", "100644", "ghi012"),
	}

	expectedMerged := []types.IndexEntry{
		types.CreateIndexEntry("file1.txt", "100644", "abc123"),
		types.CreateIndexEntry("file2.txt", "100644", "def789"),
		types.CreateIndexEntry("file3.txt", "100644", "ghi012"),
	}

	result := utils.MergeIndexEntries(entries1, entries2)

	if len(result) != len(expectedMerged) {
		t.Fatalf("Expected %d merged entries, got %d", len(expectedMerged), len(result))
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Path < result[j].Path
	})

	sort.Slice(expectedMerged, func(i, j int) bool {
		return expectedMerged[i].Path < expectedMerged[j].Path
	})

	for i := range result {
		if result[i] != expectedMerged[i] {
			t.Errorf("Expected merged entry %d to be %+v, got %+v", i, expectedMerged[i], result[i])
		}
	}
}

func TestWriteEntryToIndex(t *testing.T) {

	err := utils.WriteEntryToIndex([]types.IndexEntry{})
	if err != nil {
		t.Fatalf("WriteEntryToIndex failed: %v", err)
	}

	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)
	err = utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	entries := []types.IndexEntry{
		types.CreateIndexEntry("file1.txt", "100644", "abc123"),
		types.CreateIndexEntry("file2.txt", "100644", "def456"),
	}

	err = utils.WriteEntryToIndex(entries)
	if err != nil {
		t.Fatalf("WriteEntryToIndex failed: %v", err)
	}

	parsedEntries, err := utils.ParseIndex()
	if err != nil {
		t.Fatalf("ParseIndex failed: %v", err)
	}

	sort.Slice(parsedEntries, func(i, j int) bool {
		return parsedEntries[i].Path < parsedEntries[j].Path
	})

	if len(parsedEntries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(parsedEntries))
	}
	if parsedEntries[0].Path != "file1.txt" || parsedEntries[1].Path != "file2.txt" {
		t.Errorf("Unexpected index entries: %+v", parsedEntries)
	}
	if parsedEntries[0].Mode != "100644" || parsedEntries[1].Mode != "100644" {
		t.Errorf("Unexpected index modes: %s, %s", parsedEntries[0].Mode, parsedEntries[1].Mode)
	}
	if parsedEntries[0].Hash != "abc123" || parsedEntries[1].Hash != "def456" {
		t.Errorf("Unexpected index hashes: %s, %s", parsedEntries[0].Hash, parsedEntries[1].Hash)
	}
}

func TestRemoveEntryFromIndex(t *testing.T) {

	tempDir := t.TempDir()
	paths.SetBasePath(tempDir)
	err := utils.CreateOwnFolder("main", "")
	if err != nil {
		t.Fatalf("CreateOwnFolder failed: %v", err)
	}

	err = utils.RemoveEntryFromIndex("nonexisting.txt")
	if err != nil {
		t.Fatalf("RemoveEntryFromIndex failed: %v", err)
	}

	err = utils.WriteEntryToIndex([]types.IndexEntry{
		types.CreateIndexEntry("file1.txt", "100644", "abc123"),
		types.CreateIndexEntry("file2.txt", "100644", "def456"),
	})
	if err != nil {
		t.Fatalf("WriteEntryToIndex failed: %v", err)
	}

	err = utils.RemoveEntryFromIndex("file1.txt")
	if err != nil {
		t.Fatalf("RemoveEntryFromIndex failed: %v", err)
	}

	parsedEntries, err := utils.ParseIndex()
	if err != nil {
		t.Fatalf("ParseIndex failed: %v", err)
	}

	if len(parsedEntries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(parsedEntries))
	}
	if parsedEntries[0].Path != "file2.txt" {
		t.Errorf("Unexpected index entry: %+v", parsedEntries[0])
	}
	if parsedEntries[0].Mode != "100644" {
		t.Errorf("Unexpected index mode: %s", parsedEntries[0].Mode)
	}
	if parsedEntries[0].Hash != "def456" {
		t.Errorf("Unexpected index hash: %s", parsedEntries[0].Hash)
	}

}
