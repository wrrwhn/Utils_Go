package main

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

func main() {
	key := uuid.New().String()
	fmt.Printf("uuid[%s]-> key[%s]", key, strings.Replace(key, "-", "", -1))
}
