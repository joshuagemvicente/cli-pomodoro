package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Pomodoro Timer (made with <3 by gem) ===")
	fmt.Println("1. Get Started!")
	fmt.Println("2. Exit :(")

	input_reader := bufio.NewReader(os.Stdin) // usign standar input
	choice, _ := input_reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "1" {
		selectedMinutes := selectedWorkMinutes()
		fmt.Printf("You selected a %d-minute Pomodoro session.\n", selectedMinutes)

		breakMinute := selectedBreakMinutes()
		fmt.Printf("You selected a %d-minute break.\n", breakMinute)

		if confirmStart() {
			fmt.Println("Starting your pomodoro timer...")
			startTimer(selectedMinutes, "Work")
		}

		if breakMinute > 0 {
			fmt.Println("Break time starts now...")
			startTimer(breakMinute, "Break")
		}

	} else if choice == "2" {
		fmt.Println("Thank you for using the Pomodoro timer.")
		os.Exit(0)
	}
	fmt.Println("Session complete! Great job!")
}

func selectedWorkMinutes() int {
	fmt.Println("=== Select Pomodoro Duration ===")
	fmt.Println("1. 10 minutes.")
	fmt.Println("2. 15 minutes.")
	fmt.Println("3. 25 minutes.")
	fmt.Print("Enter your choice (1-3):")

	input_reader := bufio.NewReader(os.Stdin)
	choice, _ := input_reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		return 10
	case "2":
		return 15
	case "3":
		return 25
	default:
		fmt.Println("Invalid input. defaulting to 10 minutes.")
		return 10
	}
}

func selectedBreakMinutes() int {
	fmt.Println("=== Select Pomodoro Break Duration ===")
	fmt.Println("1. 5 minutes.")
	fmt.Println("2. 10 minutes.")
	fmt.Println("3. 15 minutes.")
	fmt.Print("Enter your choice (1-3):")

	input_reader := bufio.NewReader(os.Stdin)
	choice, _ := input_reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		return 5
	case "2":
		return 10
	case "3":
		return 15
	default:
		fmt.Println("Invalid input. defaulting to no break.")
		return 0
	}
}

func confirmStart() bool {
	fmt.Println("\n=== Ready to Start? ===")
	fmt.Println("1. Yes, start the session")
	fmt.Println("2. Cancel and exit")
	fmt.Print("Enter choice (1 or 2): ")

	reader := bufio.NewReader(os.Stdin)
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		return true
	case "2":
		fmt.Println("Session cancelled. Goodbye!")
		return false
	default:
		fmt.Println("Invalid choice. Cancelling by default.")
		return false
	}
}

func startTimer(minutes int, label string) {
	duration := minutes * 60

	for i := duration; i > 0; i-- {
		fmt.Printf("\r%s Time Remaining: %02d:%02d", label, i/60, i%60)
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("\r%s complete!                            \n", label)
}
