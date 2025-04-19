package paths

import (
	"io/fs"
	"path/filepath"
)

// GetAllFiles takes a root directory as input and returns a slice of all files
// found below that root. The returned slice is sorted by the path of the file.
// If an error occurs while traversing the directory tree, the error is returned.
//
// The returned paths are absolute paths relative to the root.
func GetAllFiles(root string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		// Skip the .own-git folder and its contents
		if d.IsDir() && (d.Name() == ".own-git" || d.Name() == ".git") {
			return fs.SkipDir
		}

		if !d.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}
