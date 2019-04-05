package shop

import (
	"fmt"

	"github.com/BrothersChimes/NoblesPayCash/item"
	"github.com/BrothersChimes/NoblesPayCash/selection"
)

const uniqueCustName = "Ulric"

// DoSales sets up a stock and then allows the player to sell the items
func DoSales(reader selection.AnswerProvider) {
	itemTypes, storeStock := ShopSetup()
	TransactionLoop(reader, itemTypes, &storeStock)
}

// ShopSetup prepares a stock of basic items for trying out sales
func ShopSetup() (map[string]item.Type, map[string]int) {
	itemTypes := map[string]item.Type{
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

func TransactionLoop(reader selection.AnswerProvider, itemTypes map[string]item.Type, storeStock *map[string]int) {
	Says(uniqueCustName, "Hi, Bailoe!")

	customerRequests := []string{"I would like to purchase a weapon!", "I would STILL like to purchase a weapon!", "PLEASE sell me a weapon...", "A weapon, please!"}
	customerRequestIndex := 0

	isSatisfied := false

	for !isSatisfied {
		nextString := customerRequests[customerRequestIndex]
		AnnounceItemQty(*storeStock, itemTypes)
		Says(uniqueCustName, nextString)
		if customerRequestIndex < len(customerRequests)-1 {
			customerRequestIndex++
		}
		SellWeapons(reader, storeStock, itemTypes, &isSatisfied)
	}
	Says(uniqueCustName, "Bye, Mr Celhai.")
	fmt.Println(uniqueCustName + " leaves.")
}

func SellWeapons(reader selection.AnswerProvider, stock *map[string]int, types map[string]item.Type, isSatisfied *bool) {
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
		SellWeapon(reader, "sword", stock, types)
	case 2:
		SellWeapon(reader, "axe", stock, types)
	}

}

func SellWeapon(reader selection.YesNoAnswerProvider, weapon string, stock *map[string]int, types map[string]item.Type) {
	weaponName := types[weapon].Name
	weaponPluralName := types[weapon].PluralName

	if (*stock)[weapon] <= 0 {
		fmt.Println("You have no " + weaponPluralName + " left for sale.")
	}

	fmt.Println("Would you like to sell Ulric a " + weaponName + "? (y/n)")

	answer := reader.GetAnswer()

	if answer {
		Does(uniqueCustName, "happily takes the "+weaponName+".")
		(*stock)[weapon]--
	} else {
		Does(uniqueCustName, "is sad you did not sell him the "+weaponName+".")
	}
}

func AnnounceItemQty(stock map[string]int, types map[string]item.Type) {
	fmt.Printf("You have %v %s.\n", stock["sword"], types["sword"].PluralName)
	fmt.Printf("You have %v %s.\n", stock["axe"], types["axe"].PluralName)
	fmt.Printf("You have %v %s.\n", stock["trailMix"], types["trailMix"].PluralName)
}

func Does(name, action string) {
	fmt.Println(name + " " + action)
}

func Says(name, speech string) {
	fmt.Println(name + " says: " + "\"" + speech + "\"")
}
