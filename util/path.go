package util

import (
	"path"
	"runtime"
)

// __FILE__
func GetCurrentFile() string {
	_, filename, _, _ := runtime.Caller(1)
	return filename
}

// __DIR__
func GetCurrentDir() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Dir(filename)
}
