package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string{

	var splitText []string = strings.Split(text, " ")

	return splitText
}

func main(){
	fmt.Println("Hello, World!")
}
