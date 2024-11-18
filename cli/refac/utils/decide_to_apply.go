package utils

import (
	"bufio"
	"fmt"
	"os"
)

func DecideToApply() bool {
	fmt.Println("Do you want to apply this refactored code?")
	for {
		fmt.Print("(y/n):")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		reply := scanner.Text()
		switch reply {
		case "y":
			fmt.Println("Refactored code applied!")
			return true
		case "n":
			fmt.Println("The code has been restored to its original")
			return false
		default:
			fmt.Println("Please answer in y/n.")
			continue
		}
	}
}
