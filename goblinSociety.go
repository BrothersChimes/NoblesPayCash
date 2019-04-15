package main

import (
	"fmt"
	"strconv"
)

func log(output	string) {
	fmt.Println(output)
}

type Goblin struct {
	name      string
	inventory map[string]int
	job       string
}

type Squad struct {
	members []*Goblin
}

type GoblinSociety struct {
	squads []*Squad
}

func (goblin *Goblin) GetInventory() map[string]int {
	return goblin.inventory
}

func (squad *Squad) GetInventory() map[string]int {
	squadInventory := make(map[string]int)
	for _, goblin := range squad.members {
		goblinInventory := goblin.GetInventory()
		for itemName, itemCount := range goblinInventory {
			squadInventory[itemName] += itemCount
		}
	}
	return squadInventory
}

func (society *GoblinSociety) GetInventory() map[string]int {
	societyInventory := make(map[string]int)
	for _, squad := range society.squads {
		squadInventory := squad.GetInventory()
		for itemName, itemCount := range squadInventory {
			societyInventory[itemName] += itemCount
		}
	}
	return societyInventory
}

func (society *GoblinSociety) ListItems() {
	items := society.GetInventory()
	for k, v := range items {
		log(k + " : " + strconv.Itoa(v));
	}
}

func (society *GoblinSociety) HasItem(itemName string) bool {
	inventory := society.GetInventory()
	_, found := inventory[itemName]
	return found
}

func (society *GoblinSociety) CountItem(itemName string) int {
	inventory := society.GetInventory()
	count := inventory[itemName]
	return count
}

func main() {
	goblinItems := map[string]int{
		"spears": 1,
		"meat": 5,
		"sticks": 20,
	}
	skinnard := Goblin{name: "skinnard", inventory: goblinItems, job: "skinner"}
	squadMembers := []*Goblin{&skinnard}
	skinsquad := Squad{members: squadMembers}
	squads := []*Squad{&skinsquad}
	skinSociety := GoblinSociety{squads: squads}
	for i := 1;  i < 5; i++ {
		// main world loop
		log("a day has passed")
		
	}
	skinSociety.ListItems()

}

