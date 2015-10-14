package main

import (
	"encoding/json"
	"fmt"
	"os"

	"../../ulambda"
)

func main() {
	var expr []interface{}
	err := json.Unmarshal([]byte(os.Args[1]), &expr)
	if err != nil {
		fmt.Print(err)
	}
	n := ulambda.NewNode(expr)
	fmt.Printf("%s\n", n)
}
