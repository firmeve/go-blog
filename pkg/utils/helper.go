package utils

import (
	"path"
	"runtime"
)

func CurrentFile() string {
	// 0 是当前helper
	_ , file , _ , ok  := runtime.Caller(1)
	if !ok {
		panic(`Can not get current file info`)
	}

	return file
}

func CurrentDir() string {
	file := CurrentFile()
	return path.Dir(file)
}