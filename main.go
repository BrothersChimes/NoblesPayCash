package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type itemType struct {
	name       string
	pluralName string
	isWeapon   bool
}

const uniqueCustName = "Ulric"

func main() {
	reader := bufio.NewReader(os.Stdin)

	itemTypes := map[string]itemType{
		"sword":    {"sword", "swords", true},
		"axe":      {"axe", "axes", true},
		"trailMix": {"bag of trail mix", "bags of trail mix", false},
	}

	storeStock := map[string]int{
		"sword":    1,
		"axe":      1,
		"trailMix": 1,
	}

	announceItemQty(storeStock, itemTypes)

	says(uniqueCustName, "Greetings!")
	says(uniqueCustName, "I would like to purchase a sword!")

	sellSwords(reader, &storeStock, itemTypes)

	says(uniqueCustName, "I would still like to purchase a sword!")
	announceItemQty(storeStock, itemTypes)

	sellSwords(reader, &storeStock, itemTypes)
	announceItemQty(storeStock, itemTypes)

}

func sellSwords(reader *bufio.Reader, stock *map[string]int, types map[string]itemType) {
	if (*stock)["sword"] <= 0 {
		fmt.Println("You inform " + uniqueCustName + " that you have no " + types["sword"].pluralName + " left for sale.")
		fmt.Println(uniqueCustName + " leaves.")

		return
	}

	fmt.Println("Would you like to sell Ulric a sword? (y/n)")

	answer, _ := reader.ReadString('\n')
	answerChar := unicode.ToUpper(rune(answer[0]))

	switch answerChar {
	case 'Y':
		does(uniqueCustName, "happily takes the sword!")
		(*stock)["sword"]--
	case 'N':
		does(uniqueCustName, "is sad you did not sell him the sword.")
	default:
		does(uniqueCustName, "did not understand what you said.")
	}
}

func announceItemQty(stock map[string]int, types map[string]itemType) {
	fmt.Printf("You have %v %s.\n", stock["sword"], types["sword"].pluralName)
	fmt.Printf("You have %v %s.\n", stock["axe"], types["axe"].pluralName)
	fmt.Printf("You have %v %s.\n", stock["trailMix"], types["trailMix"].pluralName)
}

func does(name, action string) {
	fmt.Println(name + " " + action)
}

func says(name, speech string) {
	fmt.Println(name + " says: " + "\"" + speech + ".\"")
}
