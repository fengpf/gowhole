package main

import (
	"bufio"
	"bytes"
	"os"
	"path"
	"runtime"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	println(dir)
	println(getPkgName() + "/" + getFileName())
}

func getPkgName() string {
	_, filenamePath, _, _ := runtime.Caller(0)
	file, _ := os.Open(filenamePath)
	r := bufio.NewReader(file)
	line, isPrefix, err := r.ReadLine()
	println(string(line), isPrefix, err)
	pkgName := bytes.TrimPrefix(line, []byte("package "))
	return string(pkgName)
}

func getFileName() string {
	pc, file, line, ok := runtime.Caller(0)
	println(pc, file, line, ok)
	filenamePath := path.Base(file)
	suffix := path.Ext(filenamePath)
	filename := strings.TrimSuffix(filenamePath, suffix)
	println(filenamePath, suffix, filename)
	return filename
}
