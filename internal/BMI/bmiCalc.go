package BMI

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

	MenuLogo()
	newPerson := getInput()

	bmi := getBmi(newPerson)
	fmt.Printf("\n%v has a BMI of %.2f \n", newPerson.name, bmi)
	Menu()

}

func getInput() Person {

	reader := bufio.NewReader(os.Stdin)
	newUser := Person{}
	var nameVer, weightVer, heightVer bool = false, false, false

	color.Set(color.FgYellow)

	for !nameVer {
		fmt.Println("\nPlease enter your name:")
		name, _ := reader.ReadString('\n')
		newUser.name = strings.TrimSpace(name)
		if len(newUser.name) < 1 || len(newUser.name) > 15 {
			fmt.Println("ERROR: Enter a name with a length of 1-15 char")
		} else {
			nameVer = true
		}
	}

	for !weightVer {
		fmt.Println("\nPlease enter your weight(kg):")
		weight, _ := reader.ReadString('\n')
		newUser.weight, _ = strconv.ParseFloat(strings.TrimSpace(weight), 64)
		if newUser.weight < 0 || newUser.weight > 635 {
			fmt.Println("ERROR: Enter a valid weight between 0-635kg")
		} else {
			weightVer = true
		}
	}

	for !heightVer {
		fmt.Println("\nPlease enter your height(m):")
		height, _ := reader.ReadString('\n')
		newUser.height, _ = strconv.ParseFloat(strings.TrimSpace(height), 64)
		if newUser.height < 0 || newUser.height > 3 {
			fmt.Println("ERROR: Enter a valid height beween 0-3m")
		} else {
			heightVer = true
		}
	}
	color.Unset()
	return newUser
}

func getBmi(input Person) float64 {
	//Calculates BMI from input user data
	return input.weight / (input.height * input.height)

}

func Clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
