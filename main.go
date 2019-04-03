package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/BrothersChimes/NoblesPayCash/customer"
	"github.com/BrothersChimes/NoblesPayCash/item"
)

type yesNoAnswerProvider interface {
	GetAnswer() bool
}

type numericalSelectionProvider interface {
	GetSelection(maxSelect int) int
}

type answerProvider interface {
	yesNoAnswerProvider
	numericalSelectionProvider
}

const uniqueCustName = "Ulric"

func main() {
	reader := customer.Reader{Reader: bufio.NewReader(os.Stdin)}
	doSales(reader)
}

func doSales(reader answerProvider) {
	itemTypes, storeStock := shopSetup()
	transactionLoop(reader, itemTypes, &storeStock)
}

func shopSetup() (map[string]item.ItemType, map[string]int) {
	itemTypes := map[string]item.ItemType{
		"sword":    {Name: "sword", PluralName: "swords", IsWeapon: true},
		"axe":      {Name: "axe", PluralName: "axes", IsWeapon: true},
		"trailMix": {Name: "bag of trail mix", PluralName: "bags of trail mix", IsWeapon: false},
	}

	storeStock := map[string]int{
		"sword":    1,
		"axe":      1,
		"trailMix": 1,
	}

	return itemTypes, storeStock
}

func transactionLoop(reader answerProvider, itemTypes map[string]item.ItemType, storeStock *map[string]int) {
	says(uniqueCustName, "Hi, Bailoe!")

	customerRequests := []string{"I would like to purchase a weapon!", "I would STILL like to purchase a weapon!", "PLEASE sell me a weapon...", "A weapon, please!"}
	customerRequestIndex := 0

	isSatisfied := false

	for !isSatisfied {
		nextString := customerRequests[customerRequestIndex]
		announceItemQty(*storeStock, itemTypes)
		says(uniqueCustName, nextString)
		if customerRequestIndex < len(customerRequests)-1 {
			customerRequestIndex++
		}
		sellWeapons(reader, storeStock, itemTypes, &isSatisfied)
	}
	says(uniqueCustName, "Bye, Mr Celhai.")
	fmt.Println(uniqueCustName + " leaves.")
}

func sellWeapons(reader answerProvider, stock *map[string]int, types map[string]item.ItemType, isSatisfied *bool) {
	if (*stock)["sword"] <= 0 && (*stock)["axe"] <= 0 {
		fmt.Println("You inform " + uniqueCustName + " that you have no weapon left for sale.")
		*isSatisfied = true
		return
	}

	fmt.Println("What would you like to sell Ulric?")
	fmt.Println("[1] Sword")
	fmt.Println("[2] Axe")
	fmt.Println("[0] Nothing")

	selection := reader.GetSelection(3)
	switch selection {
	case 0:
		fmt.Println("You decide not to sell anything at this point.")
	case 1:
		sellWeapon(reader, "sword", stock, types)
	case 2:
		sellWeapon(reader, "axe", stock, types)
	}

}

func sellWeapon(reader yesNoAnswerProvider, weapon string, stock *map[string]int, types map[string]item.ItemType) {
	weaponName := types[weapon].Name
	weaponPluralName := types[weapon].PluralName

	if (*stock)[weapon] <= 0 {
		fmt.Println("You have no " + weaponPluralName + " left for sale.")
	}

	fmt.Println("Would you like to sell Ulric a " + weaponName + "? (y/n)")

	answer := reader.GetAnswer()

	if answer {
		does(uniqueCustName, "happily takes the "+weaponName+".")
		(*stock)[weapon]--
	} else {
		does(uniqueCustName, "is sad you did not sell him the "+weaponName+".")
	}
}

func announceItemQty(stock map[string]int, types map[string]item.ItemType) {
	fmt.Printf("You have %v %s.\n", stock["sword"], types["sword"].PluralName)
	fmt.Printf("You have %v %s.\n", stock["axe"], types["axe"].PluralName)
	fmt.Printf("You have %v %s.\n", stock["trailMix"], types["trailMix"].PluralName)
}

func does(name, action string) {
	fmt.Println(name + " " + action)
}

func says(name, speech string) {
	fmt.Println(name + " says: " + "\"" + speech + "\"")
}
