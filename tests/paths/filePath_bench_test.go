package tests

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/alexandreLITHAUD/Own-Git/internal/paths"
)

func BenchmarkGetAllFiles(b *testing.B) {
	tempDir := b.TempDir()
	paths.SetBasePath(tempDir)

	// Create a deep directory structure with many files
	for i := 0; i < 100; i++ {
		dir := filepath.Join(tempDir, "dir", "nested", "level", "depth", filepath.Base(filepath.Join("sub", "dir", string(rune('a'+(i%26))))))
		_ = os.MkdirAll(dir, 0755)
		for j := 0; j < 10; j++ {
			file := filepath.Join(dir, "file"+string(rune('a'+(j%26)))+".txt")
			_ = os.WriteFile(file, []byte("benchmark content"), 0644)
		}
	}
	_ = os.MkdirAll(filepath.Join(tempDir, ".git"), 0755)
	_ = os.WriteFile(filepath.Join(tempDir, ".own-git", "shouldNotAppear.txt"), []byte("skip me"), 0644)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := paths.GetAllFiles(tempDir)
		if err != nil {
			b.Fatalf("Unexpected error: %v", err)
		}
	}
}
