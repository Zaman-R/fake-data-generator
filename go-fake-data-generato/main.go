package main

import (
	"fakerer/genpackage"
	"fmt"
)

func main() {
	var n int
	fmt.Println("Enter the number of Info structs to generate:")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	infos := genpackage.GenerateInfo(n)
	for _, info := range infos {
		fmt.Printf("%+v\n", info)
	}
}
