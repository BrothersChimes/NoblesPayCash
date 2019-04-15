package main

import (
	"fmt"
	"strconv"
	"math/rand"
	"time"
)

func Initialize() {
	rand.Seed(time.Now().Unix())
}

func SystemLog(output	string) {
	fmt.Println(output)
}

func PlayerLog(output	string) {
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
		SystemLog(k + " : " + strconv.Itoa(v));
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

type child struct {
	name             string
	growthPercentage int
}

func (society *GoblinSociety) spawnChild() {
	birthName := GetRandomGoblinName()
	wretchling := child{name: birthName, growthPercentage: 0}
	society.clutch = append(society.clutch, &wretchling)
	SystemLog("A wretchling named " + wretchling.name + " has been born.")
}

func GetRandomGoblinName() string {
	names := []string{"Wrterc", "Jonatan", "Rrgal",}
	return names[rand.Intn(len(names))]
}

func (squad *Squad) getSize() int {
	return len(squad.members)
}

func (society *GoblinSociety) getSize() int {
	totalGoblins := 0
	for _, squad := range society.squads {
		totalGoblins += squad.getSize()
	}
	return totalGoblins
}

func (society *GoblinSociety) spawnChildren() {
	spawnThreshold := 100
	goblinFertility := 5
	society.spawnProgress += goblinFertility * society.getSize()
	for society.spawnProgress >= spawnThreshold {
		society.spawnChild()
		society.spawnProgress -= spawnThreshold
	}
}

func (society *GoblinSociety) PassDay() {
	society.spawnChildren()
}

func main() {
	Initialize()
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
	for i := 1;  i < 40; i++ {
		// main world loop
		SystemLog("a day has passed")
		skinSociety.PassDay()
	}
	skinSociety.ListItems()

}

