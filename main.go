package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func main() {
	clearTerminal()
	var minutes int
	minutesPtr := flag.Int("m", 0, "minutes to study")

	// Old flag.Usage() kept printing the whole dir where it was located lmao
	// It was pissing me off lwk
	flag.Usage = func() {
		fmt.Println("\nFocus CLI - A simple terminal study timer")
		fmt.Println("Written in Go/Golang")
		fmt.Println("Example: focus -m <minutes>")
		fmt.Println()
		fmt.Println("Flags:")
		flag.PrintDefaults()
		fmt.Println()
	}

	flag.Parse()


	if *minutesPtr <= 0 {
		flag.Usage()
		return
	}
	
	minutes = *minutesPtr
	fmt.Println("Starting study timer for", minutes, "minute(s)...")
	
	totalSeconds := minutes * 60
	seconds := totalSeconds
	for seconds > 0 {
		time.Sleep(1 * time.Second)
		seconds -= 1

		if seconds >= 60 &&  seconds % 60 == 0 {
			clearTerminal()
			fmt.Println(seconds/60, "minute(s) left!")
		} else if seconds % 15 == 0 && seconds > 10 && seconds <= 120 {
			clearTerminal()
			fmt.Println(seconds, "seconds remaining!")
		} else if seconds <= 10 && seconds > 0 {
			clearTerminal()
			fmt.Println(seconds, "remaining!")
		}

	}

	fmt.Println("Time is up!")
}

// This is to clear the terminal cuz it was lwk getting cluttered and I wasn't a big fan
func clearTerminal() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}