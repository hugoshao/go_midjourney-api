package Util

import "fmt"

func SendLog(string2 string) {
	// 用于调试
	fmt.Printf("\n=============================\n")
	fmt.Printf("\033[1;31;40m%s\033[0m\n", string2)
	fmt.Printf("=============================\n")
}
