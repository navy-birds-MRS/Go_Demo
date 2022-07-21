package main

import (
	"fmt"
	"math"
)

func main() {
	var LoanAmount int
	var interest_rate float64
	var how_many_months int

	fmt.Print("Enter the Loan amount (ex: ₹1450) : ")
	fmt.Scanln(&LoanAmount)
	fmt.Print("Enter the interest rate (ex: 21.90): ")
	fmt.Scanln(&interest_rate)
	fmt.Print("Enter the Total months (ex: 6): ")
	fmt.Scanln(&how_many_months)

	T1 := (float64(LoanAmount) * float64((interest_rate/float64(12))/float64(100)) * (float64(math.Pow((1 + float64((interest_rate / float64(12) / float64(100)))), float64(how_many_months)))))

	T2 := (float64(math.Pow((1 + float64((interest_rate / float64(12) / float64(100)))), float64(how_many_months)))) - 1

	fmt.Printf("EMI = %.2f", T1/T2)
}
