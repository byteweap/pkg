package filex

import (
	"os"
)

// IsExist checks if a file or directory exists at the given path.
//
// It takes a string parameter `path` which represents the path to the file or directory.
// It returns a boolean value indicating whether the file or directory exists or not.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		if os.IsNotExist(err) {
			return false
		}
		return false
	}
	return true
}
