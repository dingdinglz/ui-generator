package tool

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func ErrorLog(e error) {
	rootPath, _ := os.Getwd()
	fmt.Println(time.Now().String(), e.Error())
	os.WriteFile(filepath.Join(rootPath, "ERROR.LOG"), []byte(time.Now().String()+"\n"+e.Error()), os.ModePerm)
}
