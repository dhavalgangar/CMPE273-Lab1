package main

import "fmt"
import "time"

func main() {

	fmt.Println("Hit ENTER to stop playing.")
	channel1 := make(chan string)
	channel2 := make(chan string)
	go func() {
		for {
		channel1 <- "Playing Tennis..."
		time.Sleep(time.Second * 5)
		}
	}()
	go func() {
		for {
		channel2 <- "Playing Football..."
		time.Sleep(time.Second * 8)
		}
	}()
	go func() {
		for {
			select {
			case message1 := <- channel1:
			fmt.Println(message1)
			case message2 := <- channel2:
			fmt.Println(message2)
			case <- time.After(time.Second * 5):
			fmt.Println("Taking Rest...")
			}
		}
	}()
	var input string
	fmt.Scanln(&input)
}


