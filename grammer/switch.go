package main

import "fmt"

func main() {
	fmt.Println("何か数値を入力してください : ")
	var a int
	fmt.Scan(&a)
	switch a {
	case 15:
			fmt.Println("FizzBuzz")
	case 5, 10:
			fmt.Println("Buzz")
	case 3, 6, 9, 12:
			fmt.Println("Fizz")
	default:
			fmt.Println(a)
	}
}