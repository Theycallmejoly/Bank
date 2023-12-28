package main

import (
	"fmt"
)

func main() {

	var accountbalance = 1000.0

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
		accountbalance += depositAmount
		fmt.Println("Balance updated to: ", accountbalance)
	}

	fmt.Println("your choice: ", choice)

}
