package main

import (
	"MathSkills/utils"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println("Error. Incorrect number of arguments passed.\n\nUsage: go run . data.txt")
		os.Exit(1)
	}
	// if filepath.Ext(args[0]) != ".txt" {
	// 	fmt.Println("Error. Incorrect file name !")
	// 	os.Exit(1)
	// }
	if args[0] != "data.txt" {
		fmt.Println("Error. Incorrect file name !")
		os.Exit(1)
	}
	Processed := utils.ReadFile(args[0])
	
}
