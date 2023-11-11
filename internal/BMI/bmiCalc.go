package BMI

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	name   string
	weight float64
	height float64
}

func MainProcess() {
	//controls the operation of getting the input from the users, calculating the BMI then printing it
	fmt.Println("BMI Calc started!")
	newPerson := getInput()
	fmt.Println(newPerson)

	bmi := getBmi(newPerson)
	fmt.Println(newPerson.name, "has a BMI of", bmi)
}

func getInput() Person {
	//Get input from the user
	fmt.Println("\n--------------------------\n----- BMI CALCULATOR -----\n--------------------------")

	reader := bufio.NewReader(os.Stdin)
	newUser := Person{}

	fmt.Println("\nPlease enter your name:")
	newUser.name, _ = reader.ReadString('\n')

	fmt.Println("\nPlease enter your weight:")
	weight, _ := reader.ReadString('\n')
	newUser.weight, _ = strconv.ParseFloat(weight, 64)

	fmt.Println("\nPlease enter your height:")
	height, _ := reader.ReadString('\n')
	newUser.height, _ = strconv.ParseFloat(height, 64)

	return newUser
}

func getBmi(input Person) float64 {
	//Calculates BMI from input user data
	return input.weight / (input.height * input.height)

}
