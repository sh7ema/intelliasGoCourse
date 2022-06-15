package main

import "fmt"

const priseApple float32 = 5.99
const prisePear float32 = 7.00

var money float32 = 23

func main() {
	fmt.Println("How much money should be spent to buy 9 apples and 8 pears")
	priseOne := priseApple*9 + prisePear*8
	priseOneRound := fmt.Sprintf("%.2f", priseOne)
	fmt.Println(priseOneRound)

	fmt.Println("How many pears can we buy")
	countPear := int(money / prisePear)
	fmt.Println(countPear)

	fmt.Println("How many apples can we buy")
	countApple := int(money / priseApple)
	fmt.Println(countApple)

	fmt.Println("Can we buy 2 pears and 2 apples")

	sumOffruts := priseApple*2 + prisePear*2
	// fmt.Println(sumOffruts)
	if sumOffruts <= money {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}
