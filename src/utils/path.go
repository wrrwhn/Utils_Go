package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

// GetCurrentPath 获取当前目录
func GetCurrentPath() string {

	for _, p := range os.Args {
		log.Println(p)
	}

	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if nil != err {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
