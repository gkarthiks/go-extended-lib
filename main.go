package main

import (
	"strings"
	"fmt"
)

func main() {

	data := []string{"a", "b", "c"}
	strData := strings.Join(data, "a")

	fmt.Println(strData)

	Join(data, "a", 1)

}

func Join(data []string, sep string, incrementer interface{}) {
	fmt.Println(fmt.Sprintf("%T", incrementer))

	for x:=0 ;; x++ {
		if x < 26 {
			fmt.Println(string(int('A')+x))
		} else {
			fmt.Println(string(int('A')+(x % 26)))
		}
	}

}