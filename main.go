package main

import (
	"fmt"
	"shamir/splitter"
)

func main() {
	parts, _ := splitter.Split("main.go", 4, 3)

	for key, part := range parts {
		fmt.Printf("part %d length: %d\n", key, len(part))
	}

	err := splitter.PartsToFiles(parts, "file_%d.part")

	if err != nil {
		panic(err)
	}

	//secret, _ := splitter.Combine(parts[1:])
	secret, err := splitter.CombineFiles(
		"file_0.part", "file_1.part", "file_3.part")
	if err != nil {
		panic(err)
	}
	fmt.Printf("secret:\n%v\n", string(secret))
}
