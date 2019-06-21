package main

import (
	"fmt"
	"github.com/aliciahan/leet-code/questions"
)

func main() {
	string_example := "reverseme"
	result := questions.ReverseStringRecursive(string_example)
	fmt.Println("The result is: ", result)
}
