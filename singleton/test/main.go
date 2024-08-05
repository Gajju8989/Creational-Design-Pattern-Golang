package main

import (
	"design_patterns/singleton"
	"fmt"
)

func main() {

	for i := 0; i < 100; i++ {
		fmt.Println(singleton.GetInstance())
	}

}
