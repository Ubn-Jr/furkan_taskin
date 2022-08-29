package main

import (
	"fmt"
)

//Everyone needs a nice crosshair color in fps games and a
//crosshair size that will hit the
//enemy right in the head and turn him off from the game.
//My favorite color is white(255,255,255).
//Let's enter your favorite crosshair color and size,
//then let's go to the game.
//LET'S GOOO

//?RGB Color Code Values

//RED :(255,0,0)
//BLUE :(0,0,255)
//YELLOW :(255,255,0)
//PURPLE :(255,250,0)
//ORANGE :(255,255,250)
//GREEN :(0,255,0)
//WHİTE :(255,255,255)
//BLACK :(0,0,0)
//GREY :(100,100,100)
func main() {
	//!____________________________________________________________________
	var colorCode []int = []int{}
	var primaryColors []string = []string{"RED", "BLUE", "YELLOW"}
	var secondaryColors []string = []string{"PURPLE", "ORANGE", "GREEN"}
	var neutralColors []string = []string{"WHİTE", "BLACK", "GREY"}
	//!____________________________________________________________________
	for i := 0; i < 3; i++ {
		var addColorCode int
		fmt.Printf("Enter color code -->  :")
		fmt.Scan(&addColorCode)
		colorCode = append(colorCode, addColorCode)

	}
	fmt.Println("_____________________________________________________")
	switch {
	case colorCode[0] == 255 && colorCode[1] == 0 && colorCode[2] == 0:
		fmt.Printf("your crosshair color %s has been selected \n", primaryColors[0])

	case colorCode[0] == 0 && colorCode[1] == 0 && colorCode[2] == 255:
		fmt.Printf("your crosshair color %s has been selected \n", primaryColors[1])

	case colorCode[0] == 255 && colorCode[1] == 255 && colorCode[2] == 0:
		fmt.Printf("your crosshair color %s has been selected \n", primaryColors[2])

	case colorCode[0] == 255 && colorCode[1] == 250 && colorCode[2] == 0:
		fmt.Printf("your crosshair color %s has been selected \n", secondaryColors[0])

	case colorCode[0] == 255 && colorCode[1] == 255 && colorCode[2] == 250:
		fmt.Printf("your crosshair color %s has been selected \n", secondaryColors[1])

	case colorCode[0] == 0 && colorCode[1] == 255 && colorCode[2] == 0:
		fmt.Printf("your crosshair color %s has been selected \n ", secondaryColors[2])

	case colorCode[0] == 255 && colorCode[1] == 255 && colorCode[2] == 255:
		fmt.Printf("your crosshair color %s has been selected \n ", neutralColors[0])

	case colorCode[0] == 0 && colorCode[1] == 0 && colorCode[2] == 0:
		fmt.Printf("your crosshair color %s has been selected \n ", neutralColors[1])

	case colorCode[0] == 100 && colorCode[1] == 100 && colorCode[2] == 100:
		fmt.Printf("your crosshair color %s has been selected \n ", neutralColors[2])

	case colorCode[0] != 0 && colorCode[1] != 0 && colorCode[2] != 0:
		fallthrough

	default:
		fmt.Println("you entered wrong code")
		fmt.Println("your crosshair color is the default value, it is set to green")

	}
	fmt.Println("_____________________________________________________")
	//!______________________________________
	var crossHairSize []string = []string{}

	crosshairSizeSelect(crossHairSize)
	//!______________________________________
	fmt.Println("LET'S GO...")
}

//?Crosshair Size Values:
//small
//medium
//large
func crosshairSizeSelect(crossHairSize []string) {

	for i := 0; i < 1; i++ {
		var setCross string
		fmt.Printf("Enter crosshair size --> : ")
		fmt.Scanln(&setCross)
		crossHairSize = append(crossHairSize, setCross)

	}
	fmt.Println("_____________________________________________________")
	switch {

	case crossHairSize[0] == "small":
		fmt.Printf("your crosshair size %s has been selected...\n", crossHairSize[0])
	case crossHairSize[0] == "medium":
		fmt.Printf("your crosshair size %s has been selected...\n", crossHairSize[0])
	case crossHairSize[0] == "large":
		fmt.Printf("your crosshair size %s has been selected...\n", crossHairSize[0])
	default:
		fmt.Println("you entered wrong size")
		fmt.Println("your crosshair is assigned to the standard setting : medium")

	}
	fmt.Println("_____________________________________________________")

}
