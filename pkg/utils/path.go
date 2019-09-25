package utils

import (
	"path"
	"path/filepath"
	"runtime"
)

// Get current running file path
func CurrentRunFilePath() string {
	// 0 is current file, so except
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		panic(`Can not get current file info`)
	}

	return file
}

// Get current running directory path
func CurrentRunDirPath() string {
	file := CurrentRunFilePath()
	return path.Dir(file)
}

// Get current directory relative path
func CurrentRelativePath(path string) string {
	path, err := filepath.Abs(filepath.Join(CurrentRunDirPath(), path))
	if err != nil {
		panic(err)
	}

	return path
}
