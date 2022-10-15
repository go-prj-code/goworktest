package main

import (
	"fmt"

	"github.com/go-prj-code/goworktest"
)

func main() {
	fmt.Println("test go workspace")

	v := goworktest.PrintStr("Hello", 100)
	fmt.Println(v)

	s := goworktest.ForGwsprjTest()
	fmt.Println(s)

}
