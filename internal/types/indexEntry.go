package types

type IndexEntry struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Hash string `json:"hash"`
}

func CreateIndexEntry(path string, mode string, hash string) IndexEntry {
	return IndexEntry{
		Path: path,
		Mode: mode,
		Hash: hash,
	}
}
