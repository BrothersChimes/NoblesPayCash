package shop

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/brotherschimes/noblespaycash/input"
	"github.com/brotherschimes/noblespaycash/item"
	"github.com/brotherschimes/noblespaycash/selection"
)

const uniqueCustName = "Ulric"

// DoSales sets up a stock and then allows the player to sell the items
func DoSales(provider selection.AnswerProvider) {
	stock := Setup()
	TransactionLoop(provider, stock)
}

// StockItem represents an item combined with its stock-relevant information such as quantity and price
type StockItem struct {
	*item.Type
	qty int
}

// Setup prepares a stock of basic items for trying out sales
func Setup() []StockItem {
	stock := make([]StockItem, 3)

	stock[0] = StockItem{&item.Type{Name: "sword", PluralName: "swords", UseAn: false, IsWeapon: true}, 1}
	stock[1] = StockItem{&item.Type{Name: "axe", PluralName: "axes", UseAn: true, IsWeapon: true}, 1}
	stock[2] = StockItem{&item.Type{Name: "bag of trail mix", UseAn: false, PluralName: "bags of trail mix", IsWeapon: false}, 1}

	return stock
}

// TransactionLoop loops across the customer until they are satisfied
func TransactionLoop(provider selection.AnswerProvider, stock []StockItem) {
	says(uniqueCustName, "Hi, Bailoe!")

	customerRequests := []string{"I would like to purchase a weapon!", "I would STILL like to purchase a weapon!", "PLEASE sell me a weapon...", "A weapon, please!"}
	customerRequestIndex := 0

	isSoldOut := false

	for !isSoldOut {
		nextString := customerRequests[customerRequestIndex]
		announceItemQty(stock)
		says(uniqueCustName, nextString)
		if customerRequestIndex < len(customerRequests)-1 {
			customerRequestIndex++
		}
		isSoldOut = SellWeapons(provider, stock)
	}
	says(uniqueCustName, "Bye, Celhai.")
	fmt.Println(uniqueCustName + " leaves.")
}

// SellWeapons provides a list of available weapons and asks which one should be sold
func SellWeapons(provider selection.AnswerProvider, stock []StockItem) (isSoldOut bool) {

	isSoldOut = true
	for _, item := range stock {
		if item.IsWeapon && item.qty > 0 {
			isSoldOut = false
		}
	}

	if isSoldOut {
		fmt.Printf("You inform %s that you have no weapons left for sale.\n", uniqueCustName)
		return isSoldOut
	}

	SellWeaponsFoo(provider, stock)
	return isSoldOut
}

func sortedKeys(stock *map[string]int) []string {
	keys := make([]string, len(*stock))

	i := 0
	for k := range *stock {
		keys[i] = k
		i++
	}

	sort.Strings(keys)
	return keys
}

// SellWeaponsFoo ...
func SellWeaponsFoo(provider selection.AnswerProvider, stock []StockItem) {
	fmt.Printf("What would you like to sell %s?\n", uniqueCustName)

	stock2 := make([]*StockItem, len(stock)+1)

	i := 1
	for j, item := range stock {
		if item.qty > 0 {
			stock2[i] = &stock[j]
			fmt.Printf("[%v] %s\n", i, strings.Title(item.Name))
			i++
		}
	}
	fmt.Println("[0] Nothing")

	selection := provider.GetSelection(len(stock2))

	if selection == 0 {
		fmt.Println("You decide not to sell anything at this point.")
		return
	}

	SellWeapon(provider, stock2[selection])
}

// SellWeapon tries to sell a single weapon
func SellWeapon(provider selection.YesNoAnswerProvider, item *StockItem) {
	weaponName := item.Name
	weaponPluralName := item.PluralName
	aOrAn := "a"
	if item.UseAn {
		aOrAn = "an"
	}

	if item.qty <= 0 {
		fmt.Println("You have no " + weaponPluralName + " left for sale.")
		return
	}

	if !item.IsWeapon {
		fmt.Printf("%s asks how they are meant to kill goblins with %s %s ?\n", uniqueCustName, aOrAn, weaponName)
		return
	}

	fmt.Printf("Would you like to sell %s %s %s ? (y/n)\n", uniqueCustName, aOrAn, weaponName)

	answer := provider.GetAnswer()

	if answer {
		does(uniqueCustName, "happily takes the "+weaponName+".")
		item.qty--
	} else {
		does(uniqueCustName, "is sad you did not sell them the "+weaponName+".")
	}
}

func announceItemQty(stock []StockItem) {
	for _, item := range stock {
		fmt.Printf("You have %v %s.\n", item.qty, item.PluralName)
	}
}

func does(name, action string) {
	fmt.Println(name + " " + action)
}

func says(name, speech string) {
	fmt.Println(name + " says: " + "\"" + speech + "\"")
}

func main() {
	reader := input.Reader{Reader: bufio.NewReader(os.Stdin)}
	DoSales(reader)
}
