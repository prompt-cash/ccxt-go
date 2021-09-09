package utils

import "os"

// EnsureDirectoryExists ensures that a given directory path exists.
// This method can create a path with multiple non-existing parent directories
// in a single call.
func EnsureDirectoryExists(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(path, 0755)
		if err != nil {
			return err
		}
	}
	return err
}
