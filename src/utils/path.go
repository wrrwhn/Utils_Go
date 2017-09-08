package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var imageTypeLst []string
var docTypeLst []string

// TypeFormat 文件类型
type TypeFormat string

var (
	// DefaultType 其它类型
	DefaultType = TypeFormat("%v")
	// ImageType 图片类型
	ImageType = TypeFormat("![%v](%v)")
	// DocType 文档类型
	DocType = TypeFormat("[%v](%v)")
)

// ModeFormat 文件类型
type ModeFormat int

var (
	// ErrorModeFormat 其它类型
	ErrorModeFormat = ModeFormat(0)
	// DicModeFormat 文件夹类型
	DicModeFormat = ModeFormat(1)
	// FileModeFormat 文件类型
	FileModeFormat = ModeFormat(2)
)

func init() {
	imageTypeLst = []string{"ico", "jpeg", "jpg", "png", "svg", "tif", "tiff", "webp", "gif", "bmp"}
	docTypeLst = []string{"doc", "docx", "xls", "xlsx", "ppt", "pptx", "pdf", "txt", "md", "sql"}
}

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

// PathExist 判断文件路径是否存在
func PathExist(path string) (bExist bool, err error) {

	if len(path) == 0 {
		return false, errors.New("args[path] is blank")
	}

	_, err = os.Stat(path)
	if nil != err {
		return false, err
	}
	return true, nil
}

// CheckType 校验文件类型
func CheckType(path string) (format TypeFormat, err error) {

	if len(path) == 0 {
		err = errors.New("args[path] is blank")
		return
	}

	ext, err := GetExtension(path)
	if nil != err {
		return
	}

	if contain(imageTypeLst, ext) {
		format = ImageType
	} else if contain(docTypeLst, ext) {
		format = DocType
	} else {
		format = DefaultType
	}

	return
}

func contain(lst []string, key string) bool {

	for _, p := range lst {
		if p == key {
			return true
		}
	}
	return false
}

// GetExtension 获取文件后缀(xxx)
func GetExtension(path string) (ext string, err error) {

	if len(path) == 0 {
		err = errors.New("args[path] is blank")
		return
	}

	idx := strings.LastIndex(path, ".")
	if idx < 0 {
		err = fmt.Errorf("path[%v].extension is empty", path)
		return
	}

	ext = path[idx+1:]
	return
}

// GetBasename 获取文件名
func GetBasename(path string) (basename string, err error) {

	if len(path) == 0 {
		err = errors.New("args[path] is blank")
		return
	}

	basename = filepath.Base(path)
	return
}

// CheckMode 判断文件类型
func CheckMode(path string) (format ModeFormat, err error) {

	if len(path) == 0 {
		err = errors.New("args[path] is blank")
		return
	}

	format = ErrorModeFormat
	file, err := os.Stat(path)
	if err != nil {
		err = fmt.Errorf("path[%v] is not exists", path)
		return
	}

	switch mode := file.Mode(); {
	case mode.IsDir():
		format = DicModeFormat
	case mode.IsRegular():
		format = FileModeFormat
	}

	return
}
