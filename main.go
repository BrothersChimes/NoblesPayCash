package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

type itemStock struct {
	qty        int
	pluralName string
}

const uniqueCustName = "Ulric"

func main() {
	reader := bufio.NewReader(os.Stdin)

	storeStock := map[string]itemStock{
		"sword":    itemStock{1, "swords"},
		"trailMix": itemStock{1, "bags of trail mix"},
	}

	announceItemQty(storeStock)

	says(uniqueCustName, "Greetings!")
	says(uniqueCustName, "I would like to purchase a sword!")

	sellSwords(reader, &storeStock)

	says(uniqueCustName, "I would still like to purchase a sword!")
	announceItemQty(storeStock)

	sellSwords(reader, &storeStock)
	announceItemQty(storeStock)

}

func sellSwords(reader *bufio.Reader, stock *map[string]itemStock) {
	if (*stock)["sword"].qty <= 0 {
		fmt.Println("You inform " + uniqueCustName + " that you have no swords left for sale.")
		fmt.Println(uniqueCustName + " leaves.")

		return
	}

	fmt.Println("Would you like to sell Ulric a sword? (y/n)")

	answer, _ := reader.ReadString('\n')
	answerChar := unicode.ToUpper(rune(answer[0]))

	switch answerChar {
	case 'Y':
		does(uniqueCustName, "happily takes the sword!")
		newSwordStock := (*stock)["sword"]
		newSwordStock.qty--
		(*stock)["sword"] = newSwordStock
	case 'N':
		does(uniqueCustName, "is sad you did not sell him the sword.")
	default:
		does(uniqueCustName, "did not understand what you said.")
	}
}

func announceItemQty(stock map[string]itemStock) {
	fmt.Printf("You have %v %s.\n", stock["sword"].qty, stock["sword"].pluralName)
	fmt.Printf("You have %v %s.\n", stock["trailMix"].qty, stock["trailMix"].pluralName)
}

func does(name, action string) {
	fmt.Println(name + " " + action)
}

func says(name, speech string) {
	fmt.Println(name + " says: " + "\"" + speech + ".\"")
}
