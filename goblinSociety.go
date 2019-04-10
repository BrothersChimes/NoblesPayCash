package main

import (
	"fmt"
	"strconv"
)

func log(output	string) {
	fmt.Println(output)
}

type GoblinSociety struct {
	totalGoblins int
	inventory	 map[string]int
}

func (gS *GoblinSociety) GainItem(itemName string, number int) {
	gS.Inventory[itemName] += number
	return gS.Inventory[itemName]

}

func (gS *GoblinSociety) HasItem(itemName string) bool {
	_, 
	return 0
}

func ListItems
	for k, v := range gS.inventory {
		log(k + " : " + strconv.Itoa(v));
	}
}

func main() {
	goblinItems := map[string]int{
		"axe": 1,
		"meat": 20,
		"sticks": 20,
	}
	grump := &GoblinSociety{totalGoblins:9, inventory:GoblinItems}
	for i := 1;  i < 5; i++ {
		// main world loop
		log("a day has passed")
		
	}
}

