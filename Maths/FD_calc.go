package main

import "fmt"

func f32(val int) float32 {
	return float32(val)
}

func main() {
	var Investment_amount float32
	var percentage, SCDS_sum, TDS_sum float32
	var days, years int

	fmt.Print("Enter the investment amount = ₹")
	fmt.Scanln(&Investment_amount)
	Investment_amount = float32(int(Investment_amount))

	fmt.Print("Enter the percentage given by bank (eg. 5.30) = ")
	fmt.Scanln(&percentage)

	fmt.Print("Enter the days you want to calculate = ")
	fmt.Scanln(&days)

	years, days = (days / 365), (days % 365)

	//fmt.Println(years, days)

	for i := 0; i < years; i++ {
		TDS_sum = TDS_sum + (percentage * Investment_amount * ((f32(1) / f32(365)) * f32(365)) / 100)
	}

	for i := 0; i < years; i++ {
		SCDS_sum = SCDS_sum + (percentage * Investment_amount * ((f32(1) / f32(365)) * f32(365)) / 100)
		Investment_amount = Investment_amount + (percentage * Investment_amount * ((f32(1) / f32(365)) * f32(365)) / 100)
		//fmt.Print("SCDS_sum", SCDS_sum)
	}

	sum2 := (percentage * Investment_amount * ((f32(1) / f32(365)) * f32(days)) / 100)

	//fmt.Print("sum2", sum2)

	fmt.Print("Total amount after ", years*365+days, " days SCDS= ₹", int(SCDS_sum+sum2), " ,TDS= ₹", int(TDS_sum+sum2))

	fmt.Print("per month you will earn in SCDS method = ₹", int(SCDS_sum+sum2)/(12*years), " ,TDS method = ₹", int(TDS_sum+sum2)/(12*years), "\n")
}
