package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Initialize() {
	rand.Seed(time.Now().Unix())
}

func SystemLog(output string) {
	fmt.Println(output)
}

func PlayerLog(output string) {
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
	squads        []*Squad
	clutch        []*child
	spawnProgress int
}

func (goblin *Goblin) Inventory() map[string]int {
	return goblin.inventory
}

func (squad *Squad) Inventory() map[string]int {
	squadInventory := make(map[string]int)
	for _, goblin := range squad.members {
		goblinInventory := goblin.Inventory()
		for itemName, itemCount := range goblinInventory {
			squadInventory[itemName] += itemCount
		}
	}
	return squadInventory
}

func (society *GoblinSociety) Inventory() map[string]int {
	societyInventory := make(map[string]int)
	for _, squad := range society.squads {
		squadInventory := squad.Inventory()
		for itemName, itemCount := range squadInventory {
			societyInventory[itemName] += itemCount
		}
	}
	return societyInventory
}

func (society *GoblinSociety) ListItems() {
	items := society.Inventory()
	for k, v := range items {
		SystemLog(k + " : " + strconv.Itoa(v))
	}
}

func (society *GoblinSociety) HasItem(itemName string) bool {
	inventory := society.Inventory()
	_, found := inventory[itemName]
	return found
}

func (society *GoblinSociety) CountItem(itemName string) int {
	inventory := society.Inventory()
	count := inventory[itemName]
	return count
}

type child struct {
	name             string
	growthPercentage int
}

func (society *GoblinSociety) spawnChild() {
	birthName := RandomGoblinName()
	wretchling := child{name: birthName, growthPercentage: 0}
	society.clutch = append(society.clutch, &wretchling)
	SystemLog("A wretchling named " + wretchling.name + " has been born.")
}

func RandomGoblinName() string {
	names := []string{"Wrterc", "Jonatan", "Rrgal"}
	return names[rand.Intn(len(names))]
}

func (squad *Squad) Size() int {
	return len(squad.members)
}

func (society *GoblinSociety) Size() int {
	totalGoblins := 0
	for _, squad := range society.squads {
		totalGoblins += squad.Size()
	}
	return totalGoblins
}

func (society *GoblinSociety) SpawnChildren() {
	spawnThreshold := 100
	goblinFertility := 5
	society.spawnProgress += goblinFertility * society.Size()
	for society.spawnProgress >= spawnThreshold {
		society.spawnChild()
		society.spawnProgress -= spawnThreshold
	}
}

func (society *GoblinSociety) PassDay() {
	society.SpawnChildren()
}

func main() {
	Initialize()
	goblinItems := map[string]int{
		"spears": 1,
		"meat":   5,
		"sticks": 20,
	}
	skinnard := Goblin{name: "skinnard", inventory: goblinItems, job: "skinner"}
	squadMembers := []*Goblin{&skinnard}
	skinsquad := Squad{members: squadMembers}
	squads := []*Squad{&skinsquad}
	skinSociety := GoblinSociety{squads: squads}
	for i := 1; i < 41; i++ {
		// main world loop
		SystemLog("a day has passed")
		skinSociety.PassDay()
	}
	skinSociety.ListItems()

}
