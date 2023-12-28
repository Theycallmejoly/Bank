package main

import (
	"fmt"
)

func main() {

	var accountbalance = 1000.0

	for {
		fmt.Println("Welcome!")
		fmt.Println("What do you want to do?")
		fmt.Println("1-Check the Balance")
		fmt.Println("2-Deposit Money")
		fmt.Println("3-Withdraw Money")
		fmt.Println("4-Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountbalance)
		case 2:
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("Invalid amount! The amount you enter must be greater than 0")
				continue
			}
			accountbalance += depositAmount
			fmt.Println("Balance updated to: ", accountbalance)
		case 3:
			fmt.Print("How much money do you want to withdraw? ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("Invalid amount! The amount you enter must be greater than 0")
				continue
			} else if withdraw > accountbalance {
				fmt.Println("You don't have enough balance!")
				return
			}
			accountbalance -= withdraw
			fmt.Println("The updated account balance is: ", accountbalance)
		case 4:
			fmt.Println("Have a nice day!")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
		fmt.Println("Thanks for choosing our bank")
	}
}
