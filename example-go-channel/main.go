package main

import (
	"fmt"
	"log"
	"time"
)

var message = make(chan string)

func printLoop(){

	fmt.Println("start Loop")
	i:=0
	for{
		select {
		case msg := <-message:
			fmt.Println("Message received:",msg)
			return
		default:
			fmt.Println("nothing received continue")
			i+=1
			fmt.Println("the",i,"loop")
			time.Sleep(1*time.Second)
		}
	}
}

func main() {

	go printLoop()

	log.Println("wait for 3 s")
	time.Sleep(3*time.Second)

	message <- "end"

	log.Println("wait for 10 s")
	time.Sleep(10*time.Second)




}
