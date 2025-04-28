package main

import (
	"fmt"
	"github.com/djang0man/sdui/copyschemas"
)

func main() {
	copiedFiles, err := copyschemas.CopyGraphqlSchemas("./graphql")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	for _, path := range copiedFiles {
		fmt.Println(path)
	}
}
