package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("File name missing")
		return
	} else if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	file, _ := os.Open(os.Args[1])

	arr := make([]byte, 14)

	file.Read(arr)

	fmt.Println(string(arr))

	file.Close()
}
