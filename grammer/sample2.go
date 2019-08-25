package main

import "fmt"

func loop()  {
	n := 0
	for n < 10 {
			fmt.Printf("n = %d\n", n)
			n++
	}
}

func main()  {
	loop()
}


/*
無限ループは
for {
    doSomething()
}
*/