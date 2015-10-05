package main

import (
	"encoding/json"
	"fmt"
	"os"

	"./lambda"
)

func main() {
	var expr []interface{}
	err := json.Unmarshal([]byte(os.Args[1]), &expr)
	if err != nil {
		fmt.Print(err)
	}
	n := lambda.NewNode(expr)
	fmt.Printf("%s\n", n)
}
