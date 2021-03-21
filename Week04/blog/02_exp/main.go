package main

import (
	_ "embed"

	"github.com/Sirupsen/logrus"
)

//go:embed go.mod
var f string

func main() {
	logrus.Info(f)
}
