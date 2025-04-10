package types

type IndexEntry struct {
	Path string `json:"path"`
	Mode string `json:"mode"`
	Hash string `json:"hash"`
}
