package main

import (
	"fmt"
	"math"
)

/*
 输入任意类型，转换成int返回
 */
func typeSwitch(x interface{}) int {
	switch t := x.(type) {
	case int:
		return t
	case float64:
		return int(math.Ceil(t))
	}
	return 0
}

func main() {
	num := 56.7890
	num2 := typeSwitch(num)
	fmt.Println(num2)
}