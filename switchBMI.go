package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	//Insert Code Here
	var userHeight string
	fmt.Println("Please enter your height in meter. Example key in 1.75 instead 1.75 meter or 1.75m")
	_, _ = fmt.Scanln(&userHeight)
	height, err := strconv.ParseFloat(userHeight, 0)
	if err != nil {
		log.Fatal(err)
		return
	}
	var userWeight string
	fmt.Println("Please enter your weight in KG. Example key in 70 instead of 70 kilogram or 70KG")
	_, _ = fmt.Scanln(&userWeight)
	weight, err := strconv.ParseFloat(userWeight, 0)
	if err != nil {
		log.Fatal(err)
		return
	}
	sqHeight := math.Pow(height, 2)

	bmi := weight / sqHeight

	switch {
	case bmi < 18.5:
		fmt.Println(bmi, "You're Underweight")
	case bmi > 18.5 && bmi < 24.9:
		fmt.Println(bmi, "You're healthy weight")
	case bmi > 25 && bmi < 29.9:
		fmt.Println(bmi, "You're overweight")
	case bmi > 30 && bmi < 34.9:
		fmt.Println(bmi, "You're obese")
	case bmi > 35 && bmi < 39.9:
		fmt.Println(bmi, "You're severely obese")
	case bmi >= 40:
		fmt.Println(bmi, "You're morbidly obese")

	}

}
