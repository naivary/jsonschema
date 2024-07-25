package fsutil

import (
	"io/fs"
	"os"
)

func MkdirAllIfNotExsting(path string, perm fs.FileMode) error {
	_, err := os.Stat(path)
	if !os.IsNotExist(err) {
		return nil
	}
	err = os.MkdirAll(path, perm)
	return err
}
