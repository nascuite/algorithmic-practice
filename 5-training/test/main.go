package main

import (
	"os"
)

func main() {
	err := foo()
	if err == nil {

	} // false
}

func foo() error {
	var err *os.PathError // nil
	return err
}
