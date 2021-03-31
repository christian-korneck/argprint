package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("-----------------")

	fmt.Printf("args: %d\n", len(os.Args))
	for _, arg := range os.Args {
		fmt.Println(string(arg))
		_ = arg
	}

	fmt.Println("-----------------")

}

