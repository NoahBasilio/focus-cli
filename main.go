package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func main() {
	var minutes int
	var seconds int
	

	clearTerminal()
	minutesPtr := flag.Int("m", 0, "minutes to study")
	secondsPtr := flag.Int("s", 0, "seconds to study")

	// Old flag.Usage() kept printing the whole dir where it was located lmao
	// It was pissing me off lwk
	flag.Usage = func() {
		fmt.Println("\nFocus CLI - A simple terminal study timer")
		fmt.Println("Written in Go/Golang")
		fmt.Println("Example: focus -m <minutes> -s <seconds>")
		fmt.Println()
		fmt.Println("Flags:")
		flag.PrintDefaults()
		fmt.Println()
	}

	flag.Parse()


	if *minutesPtr <= 0 && *secondsPtr <= 0 {
		flag.Usage()
		return
	} 
	
	minutes = *minutesPtr
	seconds = *secondsPtr
	if *secondsPtr <= 0 {
		fmt.Println("Starting study timer for", minutes, "minute(s)...")
	} else if *minutesPtr <= 0 {
		fmt.Println("Starting study timer for", seconds, "second(s)...")
	} else {
		fmt.Println("Starting a study timer for", minutes, "minute(s) and", seconds, "second(s)...")
	}
	
	
	totalSeconds := (minutes * 60) + seconds
	seconds = totalSeconds
	for seconds > 0 {
		time.Sleep(1 * time.Second)
		seconds -= 1
		
		// Finally got this ASCII bar to work, it wasn't even that hard idk why I struggled so much
		green := "\033[32m"
		reset := "\033[0m"

		elapsed := totalSeconds - seconds
		filled := (elapsed * 25) / totalSeconds
		empty := 25 - filled
		bar := green + strings.Repeat("#", filled) + reset + strings.Repeat("-", empty)

		minsLeft := seconds / 60
		secsLeft := seconds % 60
		fmt.Printf("\r[%s] %02d:%02d remaining", bar, minsLeft, secsLeft)
	}
	fmt.Println()
	terminalBeep()
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

func terminalBeep() {
	var beep *exec.Cmd
	switch runtime.GOOS {
		case "windows":
			beep = exec.Command("powershell", "-c", `(New-Object Media.SoundPlayer "C:\Windows\Media\Alarm01.wav").PlaySync()`)
		case "darwin":
			beep = exec.Command("afplay", "/System/Library/Sounds/Glass.aiff")
		default:
			beep = exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
	}

	beep.Stdout = os.Stdout
	beep.Run()
}