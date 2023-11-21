package BMI

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
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

	bmi := getBmi(newPerson)
	fmt.Println(newPerson.name, "has a BMI of", bmi)
}

func getInput() Person {
	//Get input from the user
	color.Set(color.FgGreen)
	fmt.Println("\n--------------------------\n----- BMI CALCULATOR -----\n--------------------------")
	color.Unset()

	reader := bufio.NewReader(os.Stdin)
	newUser := Person{}

	color.Set(color.FgYellow)
	fmt.Println("\nPlease enter your name:")
	name, _ := reader.ReadString('\n')
	newUser.name = strings.TrimSpace(name)

	fmt.Println("\nPlease enter your weight(kg):")
	weight, _ := reader.ReadString('\n')
	newUser.weight, _ = strconv.ParseFloat(strings.TrimSpace(weight), 64)

	fmt.Println("\nPlease enter your height(m):")
	height, _ := reader.ReadString('\n')
	newUser.height, _ = strconv.ParseFloat(strings.TrimSpace(height), 64)
	color.Unset()

	return newUser
}

func getBmi(input Person) float64 {
	//Calculates BMI from input user data
	return input.weight / (input.height * input.height)

}
