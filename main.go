package main

import (
	"fmt"
	"time"
)

func main() {
	var minutes int
	fmt.Print("How many minutes are you studying for? ")
	fmt.Scanln(&minutes)

	seconds := minutes * 60

	for seconds > 0 {
		time.Sleep(1 * time.Second)
		seconds -= 1

		if seconds >= 60 &&  seconds % 60 == 0 {
			fmt.Println(seconds/60, "minute(s) left!")
		} else if seconds % 10 == 0 && seconds > 10 {
			fmt.Println(seconds, "seconds remaining!")
		} else if seconds <= 10 && seconds > 0 {
			fmt.Println(seconds, "remaining!")
		}

	}

	fmt.Println("Time is up!")
}