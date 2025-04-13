package tests

import (
	"own/internal/paths"
	"own/internal/utils"
	"testing"
)

func BenchmarkCreateOwnFolder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tempDir := b.TempDir()
		paths.SetBasePath(tempDir)
		err := utils.CreateOwnFolder("main", "")
		if err != nil {
			b.Errorf("BenchmarkCreateOwnFolder failed: %v", err)
		}
	}
}
