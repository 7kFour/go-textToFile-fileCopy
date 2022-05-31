package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	// casting returned slice of bytes to string for readability
	fmt.Println(string(readFile("test.py")))

	fmt.Printf("\n====below is copy====\n\n")

	copyFile("test.py")
}

// readFile reads from a file and returns a slice of bytes
func readFile(s string) []byte {

	// ioutil.ReadFile() is deprecated
	a, err := os.ReadFile(s)
	errCheck(err)

	return a
}

func copyFile(s string) {

	src := s
	dst := "theCopy.py"

	bytesRead, err := os.ReadFile(src)
	errCheck(err)

	err = ioutil.WriteFile(dst, bytesRead, 0644)
	errCheck(err)
}

func errCheck(err error) {
	if err != nil {
		log.Fatalf("error = %s", err)
	}
}
