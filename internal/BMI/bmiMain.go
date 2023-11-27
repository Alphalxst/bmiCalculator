package BMI

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func Menu() {

	MenuLogo()
	fmt.Println("Select Option:\n(1)Calculate BMI\n(2)View previous entries\n(3)Exit")

	reader := bufio.NewReader(os.Stdin)
	selection, _ := reader.ReadString('\n')
	s := strings.TrimSpace(selection)
	sInt, _ := strconv.ParseInt(s, 10, 4)

	switch sInt {
	case 1:
		Clear()
		MainProcess()
	case 2:
		Entries()
	default:
		fmt.Println("Select an option from the above list")
	}

}

func MenuLogo() {
	Clear()
	color.Set(color.FgGreen)
	fmt.Println("\n--------------------------\n----- BMI CALCULATOR -----\n--------------------------")
	color.Unset()
}
