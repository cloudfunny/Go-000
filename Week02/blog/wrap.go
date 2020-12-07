package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	fmt.Printf("err: %+v", c())
}

func a() error {
	return errors.Wrap(fmt.Errorf("xxx"), "a")
}

func b() error {
	return errors.Wrap(a(), "b")
}

func c() error {
	return errors.Wrap(b(), "c")
}
