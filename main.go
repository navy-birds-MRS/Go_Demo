package main

import (
	"fmt"
	"strconv"
	"github.com/navy-birds-MRS/Go_Demo/card_check"
)

func main() {
	card := ""
	fmt.Print("hello user plz,Enter the card number :")
	fmt.Scan(&card)
	temp, _ := strconv.Atoi(string(card[len(card)-1]))
	if card_check.SixTeen_dig(card) == int8(temp) {
		fmt.Println("Card status :Card Number valid")
	}else{
		fmt.Println("Card status :Invalid")
	}
	fmt.Println("Card Name:",card_check.Category(string(card)))
}
