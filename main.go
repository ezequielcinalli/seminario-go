package main

import (
	"fmt"

	"tpe/parser"
)

func main() {
	estructura, err := parser.GetEstructura("TXFFABCD")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("estructura:", estructura)
		fmt.Println("Type:", estructura.Type)
		fmt.Println("Length:", estructura.Length)
		fmt.Println("Value:", estructura.Value)
	}
}
