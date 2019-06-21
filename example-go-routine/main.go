package main

import (
	"fmt"
	"log"
)

func f(from string){
	for i :=0; i<30; i++ {
		log.Println(from,":",i)
	}
}

func main() {

	f("direct")

	go f("routine")

	go func(msg string) {
		fmt.Println("annonymized:",msg)
	}("dddddddd")

	var input string
	_,err :=fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("done")


}
