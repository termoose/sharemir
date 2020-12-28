package main

import (
	"fmt"
	"shamir/splitter"
)

func main() {
	secret, _ := splitter.CombineFiles("file_0.part", "file_1.part", "file_2.part")
	fmt.Printf("secret: %s\n", secret)
}
