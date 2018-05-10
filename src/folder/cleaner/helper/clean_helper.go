package helper

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var keys []string = make([]string, 0)

func init() {
	keys = append(keys, "（")
	keys = append(keys, "(")
}

// CleanRepeater  清除文件夹内的重复文件
func CleanRepeater(path string, remove bool) error {

	// check
	if 0 == len(path) {
		return errors.New(fmt.Sprintf("invalid folder path[%v]", path))
	}

	// init
	rps := make([]string, 0)
	pm := make(map[string]string)

	// filter
	_ = filepath.Walk(path, func(p string, file os.FileInfo, err error) error {
		if nil == file {
			return err
		}
		bn, err := getBasename(p)
		if nil != err {
			return err
		}
		pm[bn] = p
		if checkKeyword(bn) {
			rps = append(rps, bn)
		}
		return nil
	})

	// check & delete
	for _, bn := range rps {
		obn, err := removeAfterKeyword(bn)
		if nil != err {
			log.Printf("fail to get origin basename for path[%v]", bn)
			return nil
		}
		bnp, bnOk := pm[bn]
		obnp, obnOk := pm[obn]
		if bnOk && obnOk {
			log.Printf("[%v{%v}] copy from [%v{%v}]", bn, bnp, obn, obnp)
			if remove {
				os.Remove(bnp)
			}
		} else {
			log.Printf("[%v{%v}] is only one[%v{%v}]", bn, bnp, obn, obnp)
		}
	}
	return nil
}

// checkKeyword 校验是否有关键字
func checkKeyword(s string) bool {

	for _, k := range keys {
		if strings.Contains(s, k) {
			return true
		}
	}
	return false
}

// getBasename 获取基础文件名，如 /a/b.txt -> b
func getBasename(s string) (string, error) {

	if 0 == len(s) {
		return "", errors.New("filename is empty")
	}

	fn := filepath.Base(s)
	e := filepath.Ext(s)
	bn := fn[0 : len(fn)-len(e)]
	// log.Printf("\n%s\n%s\n%s", fn, e, bn)
	// log.Printf("getBasename(%s)= %s", s, bn)

	return bn, nil
}

// removeAfterKeyword 移除关键符号后元素内容
func removeAfterKeyword(s string) (string, error) {

	if 0 == len(s) {
		return "", errors.New("basename is empty")
	}

	for _, k := range keys {
		if strings.Contains(s, k) {
			return strings.TrimRight(s[0:strings.LastIndex(s, k)], " "), nil
		}
	}
	return s, nil
}
