package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	a, b := 10, 5
	fmt.Println(operate(a, b, "￥"))
	fmt.Println(swap(1, 20))
	fmt.Println(swap2(30, 50))

	file, error := os.Open("1.txt")
	if error != nil {
		fmt.Println("打开文件有误")
	} else {
		fmt.Println(file)
	}

	fmt.Println(compute(op, 1, 2))
	fmt.Println(sum(1, 2, 3))
}

func operate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New(fmt.Sprintf("not support operate:%s", op))
	}
}

func swap(a, b int) (int, int) {
	return b, a
}

func swap2(a, b int) (x, y int) {
	x, y = b, a
	return
}

func compute(op func(int, int) int, a, b int) int {

	return op(a, b)

}

func op(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func sum(args ...int) int {
	s := 0
	for i := 0; i < len(args); i++ {
		s += args[i]
	}
	return s
}
