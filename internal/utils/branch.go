package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetBranchName() (string, error) {

	if !IsOwnFolder() {
		return "", fmt.Errorf("error: .own-git folder does not exist")
	}
	path := filepath.Join(".", ".own-git", "HEAD")
	file, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error reading file: %v", err)
	}

	regexp := regexp.MustCompile(`ref: refs/heads/(.*)`)
	matches := regexp.FindStringSubmatch(string(file))
	if len(matches) < 2 {
		return "", fmt.Errorf("error parsing branch name")
	}

	branchName := strings.TrimSpace(matches[1])
	return branchName, nil
}
