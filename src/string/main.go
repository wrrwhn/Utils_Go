package main

import (
	"fmt"
	"strings"
)

func main() {

	idx()
}

func idx() {

	ps := "av,record"
	idx := strings.Index(ps, ",")
	fmt.Println(idx, ps[0:idx])
}
