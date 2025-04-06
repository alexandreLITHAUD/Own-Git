package sctypes

type WorktreeEntry struct {
	Name     string           // "main.go", "docs/", etc.
	Mode     string           // "100644" for files, "040000" for directories
	Type     string           // "blob" or "tree"
	Path     string           // Full path relative to repo root
	Content  []byte           // Only for blobs
	Children []*WorktreeEntry // Only for trees
}
