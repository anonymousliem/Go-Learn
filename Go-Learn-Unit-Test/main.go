package main

import (
	"Go-learn-Unit-Test/helper"
	"fmt"
)

func main()  {
	name := helper.HelloWorld("A")
	fmt.Println(name)

	c := add(1,2);
	fmt.Println(c)
}

func add(a, b int)  int{
	return a + b
}
