package paths

import "path/filepath"

var basePathOverride string

// SetBasePath overrides the base path of the index file, used in the
// GetIndexFilePath function. This is useful in tests to set the base path
// to a temporary directory.
func SetBasePath(path string) {
	basePathOverride = path
}

// GetOwnGitFolderPath returns the path of the .own-git folder. If SetBasePath has been used, the
// base path is the one set by SetBasePath; otherwise, it is the current working directory.
func GetOwnGitFolderPath() string {
	if basePathOverride != "" {
		return filepath.Join(basePathOverride, ".own-git")
	}
	return filepath.Join(".", ".own-git")
}

// GetIndexFilePath returns the path of the index file. If SetBasePath has been used,
// the base path is the one set by SetBasePath; otherwise, it is the current
// working directory.
func GetIndexFilePath() string {
	if basePathOverride != "" {
		return filepath.Join(basePathOverride, ".own-git", "index")
	}
	return filepath.Join(".", ".own-git", "index")
}

// GetObjectFilePath returns the path of an object file, given its hash.
// If SetBasePath has been used, the base path is the one set by SetBasePath;
// otherwise, it is the current working directory. The object file is stored
// in subfolders of the objects folder, with the first two characters of the
// hash used as the name of the subfolder.
func GetObjectFilePath(hash string) string {
	if basePathOverride != "" {
		return filepath.Join(basePathOverride, ".own-git", "objects", hash[:2], hash[2:])
	}
	return filepath.Join(".", ".own-git", "objects", hash[:2], hash[2:])
}

// GetObjectFolderPath returns the path of the folder containing object files.
// If SetBasePath has been used, the base path is the one set by SetBasePath;
// otherwise, it is the current working directory.
func GetObjectFolderPath() string {
	if basePathOverride != "" {
		return filepath.Join(basePathOverride, ".own-git", "objects")
	}
	return filepath.Join(".", ".own-git", "objects")
}

// GetRefsFolderPath returns the path of the folder containing ref files.
// If SetBasePath has been used, the base path is the one set by SetBasePath;
// otherwise, it is the current working directory.
func GetRefsFolderPath() string {
	if basePathOverride != "" {
		return filepath.Join(basePathOverride, ".own-git", "refs")
	}
	return filepath.Join(".", ".own-git", "refs")
}

func GetHeadFilePath() string {
	if basePathOverride != "" {
		return filepath.Join(basePathOverride, ".own-git", "HEAD")
	}
	return filepath.Join(".", ".own-git", "HEAD")
}

// GetConfigFilePath returns the path of the config file. If SetBasePath has been used, the
// base path is the one set by SetBasePath; otherwise, it is the current working directory.
func GetConfigFilePath() string {
	if basePathOverride != "" {
		return filepath.Join(basePathOverride, ".own-git", "config")
	}
	return filepath.Join(".", ".own-git", "config")
}
