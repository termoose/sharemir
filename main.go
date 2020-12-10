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

	_ = splitter.PartsToFiles(parts, "file_%d.part")

	secret, _ := splitter.Combine(parts[2:])
	fmt.Printf("secret: %v\n", string(secret))
}
