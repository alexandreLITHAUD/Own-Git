package types

type WorktreeEntry struct {
	Name     string           `json:"name"`     // "main.go", "docs/", etc.
	Mode     string           `json:"mode"`     // "100644" for files, "040000" for directories
	Type     string           `json:"type"`     // "blob" or "tree"
	Path     string           `json:"path"`     // Full path relative to repo root
	Content  []byte           `json:"content"`  // Only for blobs
	Children []*WorktreeEntry `json:"children"` // Only for trees
}
