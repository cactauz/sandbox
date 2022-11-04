package main

import "fmt"

func main() {
	v, err := divide(10, 0)
	if err != nil {
		fmt.Println("uh oh:", err)
		return
	}
	fmt.Println("hello, %d", v)
}

func divide(n, d int) (int, error) {
	return n / d, nil
}
