package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	// init
	c := Confige{}
	p := "conf/conf.yaml"

	// read
	r, err := os.Open(p)
	if nil != err {
		fmt.Printf("fail to read file[%s]: %s", p, err.Error())
		return
	}
	bs, err := ioutil.ReadAll(r)
	if nil != err {
		fmt.Printf("fail to read file[%s].all: %s", p, err.Error())
		return
	}

	// parse
	err = yaml.Unmarshal(bs, &c)
	if nil != err {
		fmt.Printf("fail to unmarshal data[%s] to Confige: %s", string(bs), err.Error())
		return
	}
	fmt.Println(c.Url)
	fmt.Println(len(c.Path))
	fmt.Println(c.Path[0])
}

// Confige 配置
type Confige struct {
	// Url 请求地址
	Url string
	// Path 配置路径
	Path []string
}
