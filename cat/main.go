package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Check if any command line arguments are provided
	if len(os.Args) > 1 {
		// Read and print the contents of each file
		for _, filename := range os.Args[1:] {
			err := printFileContents(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERORR: %s\n", err)
				os.Exit(1)
			}
		}
	} else {
		// Read from stdin and print the contents
		err := printStdinContents()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			os.Exit(1)
		}
	}
}

func printFileContents(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}

func printStdinContents() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return scanner.Err()
}
