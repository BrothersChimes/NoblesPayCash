package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type itemType struct {
	name       string
	pluralName string
	isWeapon   bool
}

func (i itemType) getAnswer() bool {
	return i.isWeapon
}

type reader struct {
	*bufio.Reader
}

func (r reader) getAnswer() bool {
	for true {
		fmt.Println("Please enter (y/n)")
		answer, _ := r.ReadString('\n')
		answerChar := unicode.ToUpper(rune(answer[0]))
		switch answerChar {
		case 'Y':
			return true
		case 'N':
			return false
		default:
			fmt.Println("That character was not understood.")
		}
	}
	return false
}

func (r reader) getSelection(maxSelect int) int {
	for true {
		fmt.Println("Please enter a number.")
		answer, _ := r.ReadString('\n')
		answerNum, err := strconv.Atoi(strings.TrimSpace(answer))
		if err != nil {
			log.Fatal(err)
			fmt.Println("That entry was not understood.")
			continue
		}

		if answerNum >= maxSelect {
			fmt.Println("That number is too high.")
			continue
		}

		if answerNum < 0 {
			fmt.Println("Please enter a non-negative number.")
			continue
		}

		return answerNum
	}
	return 0
}

type yesNoAnswerProvider interface {
	getAnswer() bool
}

type numericalSelectionProvider interface {
	getSelection(maxSelect int) int
}

type answerProvider interface {
	yesNoAnswerProvider
	numericalSelectionProvider
}

const uniqueCustName = "Ulric"

func main() {
	reader := reader{bufio.NewReader(os.Stdin)}
	doSales(reader)
}

func doSales(reader reader) {
	itemTypes, storeStock := shopSetup()
	transactionLoop(reader, itemTypes, &storeStock)
}

func shopSetup() (map[string]itemType, map[string]int) {
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

	return itemTypes, storeStock
}

func transactionLoop(reader reader, itemTypes map[string]itemType, storeStock *map[string]int) {
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

func sellWeapons(reader answerProvider, stock *map[string]int, types map[string]itemType, isSatisfied *bool) {
	if (*stock)["sword"] <= 0 && (*stock)["axe"] <= 0 {
		fmt.Println("You inform " + uniqueCustName + " that you have no weapon left for sale.")
		*isSatisfied = true
		return
	}

	fmt.Println("What would you like to sell Ulric?")
	fmt.Println("[1] Sword")
	fmt.Println("[2] Axe")
	fmt.Println("[0] Nothing")

	selection := reader.getSelection(3)
	switch selection {
	case 0:
		fmt.Println("You decide not to sell anything at this point.")
	case 1:
		sellWeapon(reader, "sword", stock, types)
	case 2:
		sellWeapon(reader, "axe", stock, types)
	}

}

func sellWeapon(reader yesNoAnswerProvider, weapon string, stock *map[string]int, types map[string]itemType) {
	weaponName := types[weapon].name
	weaponPluralName := types[weapon].pluralName

	if (*stock)[weapon] <= 0 {
		fmt.Println("You have no " + weaponPluralName + " left for sale.")
		return
	}

	fmt.Println("Would you like to sell Ulric a " + weaponName + "? (y/n)")

	answer := reader.getAnswer()

	if answer {
		does(uniqueCustName, "happily takes the "+weaponName+".")
		(*stock)[weapon]--
	} else {
		does(uniqueCustName, "is sad you did not sell him the "+weaponName+".")
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
	fmt.Println(name + " says: " + "\"" + speech + "\"")
}
