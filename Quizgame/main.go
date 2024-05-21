package main

import (
	"flag"
	"fmt"
	"os"
	"time"
	"utils/helpers"
)

func main() {
	timer := flag.NewFlagSet("timer", flag.ExitOnError)
	enableTimer := timer.Bool("enableTimer", false, "Enables a timer for this quiz")
	timerSeconds := timer.Int("timerSeconds", 30, "Sets the seconds for the test")

	timer.Parse(os.Args[2:]) //enables enableTimer and timerSeconds, but not main.go -h
	// timer -h works though

	helpers.Moderator()

	if *enableTimer {
		sig := make(chan int, 1)
		ch := make(chan bool)

		// Checks if at least timer was set in the command-line
		if len(os.Args) < 2 {
			fmt.Println("Expected the 'timer' subcommand: file.go timer -options...")
			return
		}

		go func() {
			pre := helpers.Quizgame()
			if pre == nil {
				sig <- 0
				close(sig)

				ch <- true
			}
		}()

		go func() {
			if exitSleep := <-ch; exitSleep {
				return
			} else {
				time.Sleep(time.Duration(*timerSeconds) * time.Second)
				sig <- 1
			}
		}()

		//Exiting the programm
		if exitOverTimer := <-sig; exitOverTimer == 1 {
			helpers.OverTimerSay(helpers.Total, helpers.Wrong, helpers.Right)
			return
		}

		// User under timer
		if exitUnderTimer := <-sig; exitUnderTimer == 0 {
			helpers.UnderTimerSay(*timerSeconds, helpers.Total, helpers.Wrong, helpers.Right)
			return
		}

	} else {
		// fmt.Println("Else-Block") // - for debugging purposes
		helpers.Quizgame()

		fmt.Println("Here are the results: ")
		helpers.WaitingDeco()
		fmt.Printf("\nFrom %d: %d were wrong and %d right\n", helpers.Total, helpers.Wrong, helpers.Right)
	}
}
