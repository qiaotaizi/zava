package main

import (
	"fmt"
	"path"
	"testing"
)

func TestPath(t *testing.T) {
	p := "/c/Go/bin/zava.exe"
	fmt.Println(path.Base(p))
	fmt.Println(path.Clean(p))
	fmt.Println(path.Dir(p))
	fmt.Println(path.Ext(p))
	dir, file := path.Split(p)
	fmt.Printf("%s, %s\n", dir, file)
}
