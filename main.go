package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, informe o CEP")
		os.Exit(1)
	}
	fmt.Println("CEP:", os.Args[1])
}
