package main

import (
	"fmt"
	"regexp"
	"path/filepath"
)

var eightDotThreeRe = regexp.MustCompile(`^\w{1,8}\.\w{1,3}$`)

func check8dot3(name string) bool {
	return eightDotThreeRe.MatchString(name)
}

func main() {
	files, err := filepath.Glob("./*")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if !check8dot3(f) {
			fmt.Println("NG:", f)
		}
	}
}
