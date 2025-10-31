package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Items available in the market and their stock
var items = map[string]uint{
	"Apple": 10,
	"Bread": 8,
	"Milk":  6,
	"Egg":   12,
}

// Price multipliers for items (cost per unit)
var prices = map[string]uint{
	"Apple": 5,
	"Bread": 7,
	"Milk":  9,
	"Egg":   3,
}

// Generates a slice of available items according to stock
func generateItemArray(items map[string]uint) []string {
	itemArr := []string{}
	for item, count := range items {
		for i := uint(0); i < count; i++ {
			itemArr = append(itemArr, item)
		}
	}
	return itemArr
}

// Gets the user's name
func getName() string {
	fmt.Println("Welcome to Mostafa's Market!")
	fmt.Printf("Enter your name: ")
	var name string
	_, err := fmt.Scanln(&name)
	if err != nil {
		return ""
	}
	fmt.Printf("Hello %s, ready to shop?\n", name)
	return name
}

// Gets how much the user wants to spend
func getBudget(balance uint) uint {
	var spend uint
	for {
		fmt.Printf("Enter amount to spend (balance = $%d) or 0 to quit: ", balance)
		_, err := fmt.Scan(&spend)
		if err != nil {
			continue
		}
		if spend > balance {
			fmt.Println("You can't spend more than your balance!")
		} else {
			break
		}
	}
	return spend
}

// "Buys" random items for the given amount, returns what you got
func buyItems(budget uint, itemArr []string, prices map[string]uint) []string {
	rand.Seed(time.Now().UnixNano())
	bought := []string{}
	remaining := budget
	attempts := 0
	for remaining > 0 && attempts < 100 {
		idx := rand.Intn(len(itemArr))
		item := itemArr[idx]
		price := prices[item]
		if price <= remaining {
			bought = append(bought, item)
			remaining -= price
		}
		attempts++
	}
	return bought
}

// Prints your basket
func printBasket(basket []string) {
	fmt.Print("You purchased: ")
	if len(basket) == 0 {
		fmt.Println("Nothing. Try again next time!")
	} else {
		for _, item := range basket {
			fmt.Print(item, " ")
		}
		fmt.Println()
	}
}

func main() {
	name := getName()
	balance := uint(100)

	itemArr := generateItemArray(items)

	for balance > 0 {
		spend := getBudget(balance)
		if spend == 0 {
			break
		}
		balance -= spend

		basket := buyItems(spend, itemArr, prices)
		printBasket(basket)
		fmt.Printf("Current balance: $%d\n\n", balance)
	}

	fmt.Printf("%s, you left the market with $%d. Thanks for shopping!\n", name, balance)
}
