package main

// compound interest calculator
/*
P = Principal (initial amount)
r = Annual interest rate (as a decimal, so 5% becomes 0.05)
n = Number of times interest is compounded per year (e.g., 12 for monthly, 4 for quarterly)
t = Time in years
A = Final amount

A = principle * (1+ rate/100*n)^n*time
*/

import (
	"fmt"

	"math"
)

func main() {
	var principal float64
	var rate float64
	var n int
	var time float64
	
	fmt.Print("enter the principal amount : ")
	fmt.Scan(&principal)
	fmt.Print("enter the rate of interest : ")
	fmt.Scan(&rate)
	fmt.Print("enter the number of times the interest is compounded : ")
	fmt.Scan(&n)
	fmt.Print("enter the years : ")
	fmt.Scan(&time)

	totalAmount := principal * math.Pow(1+rate/float64(n),float64(n)*time)
	fmt.Printf("the amount paid after %v is %.2f\n",time,totalAmount)
	interest := totalAmount - principal
	fmt.Printf("the interest is %.2f",interest)

}
