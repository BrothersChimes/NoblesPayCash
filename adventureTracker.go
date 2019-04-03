package main

import (
	"fmt"
)

func log(output	string) {
	fmt.Println(output)
}

type adventurer struct {
	home		string
	location	string
	gold		int
	items		[]string
}

func (a *adventurer) go_home() string {
	//adventurer goes to his home
	a.location = a.home
	return a.location
}

func (a *adventurer) remove_gold(amt int) int {
	//adventurer loses amt gold		returns remainder
	a.gold -= amt
	return a.gold
}

func (a *adventurer) add_item(itemName string) string {
	a.items = append(a.items, itemName)
	return itemName
}

func (a *adventurer) buy_item(itemName string) string {
	if (itemName == "sword" && a.gold >= 8) {
			a.remove_gold(8)
			return a.add_item("sword")
	} else {
		return "big fat sack of nothing"
	}
}

func main() {
	dirk := adventurer{home: "town", location: "goblin caves", gold: 10}
	log("Dirk the Adventurer has arrived in " + dirk.go_home());
	log("Dirk the Adventurer has bought a " + dirk.buy_item("sword"));
	log("Dirk the Adventurer has left town");
	log("Dirk the Adventurer has arrived at the mouth of the caves");
	log("Dirk the Adventurer has encounted 2 goblins");
	log("Dirk the Adventurer has killed a goblin");
	log("Dirk the Adventurer has taken 1 damage");
	log("Dirk the Adventurer has killed a goblin");
	log("Dirk the Adventurer has entered the cave");
	log("Dirk the Adventurer has picked up a gemstone");
	log("Dirk the Adventurer has completed objective 'retrive gemstone'");
	log("Dirk the Adventurer has left the cave");
	log("Dirk the Adventurer has left the cave mouth");
	log("Dirk the Adventurer has arrived in town");
	log("Dirk the Adventurer has sold a sword");
	log("Dirk the Adventurer has handed in quest 'deal with the goblin raiders'");
	log("Dirk the Adventurer has received a scroll of healing");
}

