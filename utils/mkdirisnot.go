package utils

import "os"

func MkDirIfDirNotExist(name string) {
	os.Stat(name)
}
