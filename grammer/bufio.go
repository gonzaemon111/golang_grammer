package main

import (
    "bufio"
    "fmt"
    "os"
)

func buffer()  {
	var sc = bufio.NewScanner(os.Stdin)
	var s, t string
	if sc.Scan() {
			s = sc.Text()
	}
	if sc.Scan() {
			t = sc.Text()
	}
	fmt.Println("----------------")
	fmt.Println(s)
	fmt.Println(t)
}



func main() {
	buffer()
}