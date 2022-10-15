package goworktest

import "fmt"

func PrintStr(str string, num int) string {
	return fmt.Sprintf("project C %s_%d", str, num)
}
