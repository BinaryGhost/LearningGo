package helpers

import (
	"fmt"
	"time"
)

func WaitingDeco() {
	time.Sleep(1 * time.Second)
	fmt.Print("..")

	time.Sleep(100 * time.Millisecond)
	fmt.Print("....")

	time.Sleep(900 * time.Millisecond)
	fmt.Print("...")

	time.Sleep(2 * time.Second)
	fmt.Print(".")
}

func Moderator() {
	fmt.Println("Welcome to this quiz game, here you will answer very simple math equations! During the quiz \nyou can type 'exit' to get out of the quiz")
	fmt.Printf("\nReady? Go!\n\n")
}

func UnderTimerSay(timer, res, w, r int) {
	fmt.Printf("Wonderfull! You finished this quiz under %d seconds\n", timer)
	fmt.Println("Here are the results: ")
	WaitingDeco()
	fmt.Printf("\nFrom %d: %d were wrong and %d right\n", res, w, r)
}

func OverTimerSay(res, w, r int) {
	fmt.Println("Time is over :(")
	fmt.Println("Here are the results: ")
	WaitingDeco()
	fmt.Printf("\nFrom %d: %d were wrong and %d right\n", res, w, r)
}
