package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pm5/byohc-workshop/ulambda"
)

func main() {
	ast, err := ulambda.ParseExpr(string(os.Args[1]))
	if err != nil {
		log.Fatal(err)
	}
	node := ulambda.NewNode(ast)
	fmt.Printf("%s\n", ulambda.NormalForm(node))
}
