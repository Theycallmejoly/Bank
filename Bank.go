package main

import (
	"fmt"
)

func main() {

	var accountbalance = 1000.0

	for {
		fmt.Println("Welcome !")
		fmt.Println("What do you want to do? ")
		fmt.Println("1-Chech the Balance")
		fmt.Println("2-Deposit Money ")
		fmt.Println("3-withdraw money")
		fmt.Println("4-Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("your balance is: ", accountbalance)
		} else if choice == 2 {
			fmt.Print("Your Deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			if depositAmount <= 0 {
				fmt.Println("invalid amount! the amount that you enter must be greater than 0")
				continue
			}
			accountbalance += depositAmount
			fmt.Println("Balance updated to: ", accountbalance)
		} else if choice == 3 {
			fmt.Print("How much money do you want to withdraw? ")
			var withdraw float64
			fmt.Scan(&withdraw)
			if withdraw <= 0 {
				fmt.Println("invalid amount! the amount that you enter must be greater than 0")
				continue
			} else if withdraw > accountbalance {
				println("You dont have enough balance!")
				return
			}
			accountbalance -= withdraw
			fmt.Println("The updated account balance is: ", accountbalance)
		} else {
			fmt.Println("have a nice day!")
			break
		}

		fmt.Println("Thanks for choosing our bank")

	}
}

