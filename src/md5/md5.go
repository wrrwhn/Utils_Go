package main

import (
	"bytes"
	"crypto/md5"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
	"utils"
)

var command int
var path string
var tail string

func init() {
	flag.IntVar(&command, "c", 1, "work command, \n\t\t0: read the md5 value of file, default\n\t\t1: add temp file to change the md5 value")
	flag.StringVar(&path, "p", utils.GetCurrentPath(), "path for execute copy command")
	flag.StringVar(&tail, "t", "D:\\work\\git\\yao\\go\\Utils_Go\\template\\temp.txt", "the path of file for add at the tail")
	flag.Parse()
}

func main() {

	// switch
	switch command {
	case 0:
		log.Printf("read md5 value for file[%s]", path)
		fmt.Println(read(path))
	case 1:
		log.Printf("change md5 value for path|file[%s]", path)
		change(path)
	}
}

func change(path string) {

	// check
	if 0 == len(path) {
		log.Fatal("cannot read md5 for empty file")
	}

	// swith
	f, err := os.Stat(path)
	if err != nil {
		log.Fatalln(err)
	}

	switch m := f.Mode(); {
	case m.IsDir():
		log.Printf("change path[%s]", path)
		changeDir(path)
	case m.IsRegular():
		log.Printf("change file[%s]", path)
		changeFile(path, tail)
	}
}

func changeDir(path string) {

	// check
	if 0 == len(path) {
		log.Fatal("cannot read md5 for empty file")
	}

	// init
	list := make([]string, 0)

	// filter
	_ = filepath.Walk(path, func(p string, file os.FileInfo, err error) error {
		if file == nil {
			return err
		}
		if file.IsDir() {
			return nil
		}
		if filepath.Ext(p) == ".mp4" {
			log.Printf("add %s to list", p)
			list = append(list, p)
		}

		return nil
	})

	// foreach.change
	for _, p := range list {
		changeFile(p, tail)
	}
}

func changeFile(path, tail string) {

	// check
	if 0 == len(path) {
		log.Fatal("cannot read md5 for empty file")
	}

	// copy
	var stderr bytes.Buffer
	cmd := fmt.Sprintf("copy /Y %s /B +%s /B %s", path, tail, path)
	log.Printf("cmd.run(%v)", cmd)
	// 注：一定要将每个参数分开！！
	c := exec.Command("cmd", "/C ", "copy", "/Y", path, "/B", "+", tail, "/B", path)
	c.Stderr = &stderr
	if err := c.Run(); err != nil {
		log.Fatalf("%s: %s", err, stderr.String())
	}
	log.Printf("append [%s] to [%s]", tail, path)
}

func read(path string) string {

	// check
	if 0 == len(path) {
		log.Fatal("cannot read md5 for empty file")
	}

	// read
	log.Printf("read.start= %v", time.Now())
	f, err := os.Open(path)
	if nil != err {
		log.Fatal(err)
	}
	defer f.Close()

	// buffer
	m := md5.New()
	if _, err := io.Copy(m, f); err != nil {
		log.Fatal(err)
	}

	// sum
	val := fmt.Sprintf("%x", m.Sum(nil))

	log.Printf("read.end= %v", time.Now())
	return val
}
