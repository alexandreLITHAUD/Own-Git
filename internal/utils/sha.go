package utils

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func IsExecutable(filepath string) (bool, error) {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return false, err
	}

	// Check if the file has execute permission for the user
	return fileInfo.Mode()&0100 != 0, nil
}

// getFileSHA1 computes the SHA1 of a file.
//
// The function opens the file at the given path and computes its SHA1 hash. If
// the file does not exist or an error occurs while reading the file, the
// function returns an error.
//
// The returned string is the hexadecimal representation of the file's SHA1
// hash.
func GetFileSHA1(filepath string) (string, error) {
	// Open the file
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Printf("error closing file: %v\n", err)
		}
	}()

	// Create new SHA1 hash
	hash := sha1.New()

	// Copy file contents to the hash
	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	// Get the hash sum as a byte slice
	hashInBytes := hash.Sum(nil)

	// Convert to hex string
	hashString := hex.EncodeToString(hashInBytes)

	return hashString, nil
}
