package main

import (
	"fmt"
	"go-repeat-1/internal/file"
)

func main(){
	content, err := file.ReadFile("internal/testdata/inpt.txt")

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Println(content)
}