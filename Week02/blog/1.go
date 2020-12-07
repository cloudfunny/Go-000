package main

import (
	"errors"
	"fmt"
	"os"
)

type errorString struct {
	text string
}

func (e errorString) Error() string {
	return e.text
}

// New 创建一个自定义错误
func New(s string) error {
	return errorString{text: s}
}

var errorString1 = New("test a")
var err1 = errors.New("test b")

func main() {
	if errorString1 == New("test a") {
		fmt.Println("err string a")
	}
	os.IsExist()
	if err1 == errors.New("test b") {
		fmt.Println("err b")
	}
}
