package goworktest

import "fmt"

func PrintStr(str string, num int) string {
	return fmt.Sprintf("project C %s_%d", str, num)
}

func ForGwsprjTest() string {
	return "hi go workspace test prj1"
}
