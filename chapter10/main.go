package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ": ", i)
		amt := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * amt)
	}
}

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func printer(c <-chan string) {
	for {
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
}

func mySleep(c time.Duration) {
	<-time.After(c)
}

func channel() {
	var c chan string = make(chan string)

	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}

func main() {
	// for i := 0; i < 10; i++ {
	// 	go f(0)
	// }
	// var input string
	// fmt.Scanln(&input)

	// channel()
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			mySleep(time.Second * time.Duration(2))
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			mySleep(time.Second * time.Duration(3))
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <-c1:
				fmt.Println("Message 1", msg1)
			case msg2 := <-c2:
				fmt.Println("Message 2", msg2)
			case <-time.After(time.Second):
				fmt.Println("timeout")
				// default:
				// 	fmt.Println("nothing ready")
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}

/**
Problems:
1. You can specify the direction of a channel with the left arrow (<-). chan<- can only be send. <-chan can only receive only.
2.
3. A buffered channel an asynchronous channel that does not wait unless the channel is already full.
*/
