package main

import "fmt"

// Declare constant
const Title = "Personal Details"

//Declare package variables
var Country = "USA"

func main() {

	fname, lname := "fname", "lname"

	age := 35

	fmt.Println("fname: ", fname)
	fmt.Println("lname: ", lname)
	fmt.Println("age: ", age)
	fmt.Println("country: ", Country)
	fmt.Println("Title: ", Title)

}
