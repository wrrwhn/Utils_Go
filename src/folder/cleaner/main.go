package main

import (
	"flag"
	"../cleaner/helper"
	"log"
)

func main() {

	var path string
	var remove bool
	flag.StringVar(&path, "path", "", "input the path for clean repeat, please")
	flag.BoolVar(&remove, "remove", false, "is real to remove repeat file")
	flag.Parse()

	err := helper.CleanRepeater(path, remove)
	if nil != err {
		log.Printf("fail to clean folder[%v]", path)
	} else {
		log.Printf("success to clean folder[%v]", path)
	}
}
