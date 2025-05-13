package tool

import "os"

func FileExist(_path string) bool {
	_, e := os.Stat(_path)
	return e == nil
}

func DirCreate(_path string) {
	if !FileExist(_path) {
		os.Mkdir(_path, os.ModePerm)
	}
}
