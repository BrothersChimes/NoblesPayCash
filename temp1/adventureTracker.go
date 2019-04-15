package main

import (
	"fmt"
	"strconv"
)

func log(output string) {
	fmt.Println(output)
}

type Human interface {
	Gold() int
	GoHome() string
}

type GoldHaver struct {
	gold int
}

type Glutton struct {
	GoldHaver
}

type shopkeeper struct {
	GoldHaver
	Glutton
	items map[string]bool
	shop  string
}

type adventurer struct {
	GoldHaver
	home     string
	items    []string
	location string
	weapon   string
}

func (g *GoldHaver) Gold() int {
	return g.gold
}

func (s *shopkeeper) GoHome() string {
	return s.shop
}

func (s *shopkeeper) AddGold(amt int) int {
	s.gold += amt
	return s.gold
}

func (s *shopkeeper) RemoveItem(itemName string) string {
	delete(s.items, itemName)
	return itemName
}

func (s *shopkeeper) SellItem(itemName string) string {
	if itemName == "Blackstaff" {
		s.AddGold(20)
	} else if itemName == "Longstaff" {
		s.AddGold(10)
	} else {
		return "big fat sack of nothing"
	}
	return s.RemoveItem(itemName)
}

func (a *adventurer) SwingWeapon() string {
	return "hit with his " + a.weapon + " for 3 damage"
}

func (a *adventurer) GoHome() string {
	//adventurer goes to his home
	a.location = a.home
	return a.location
}

func (a *adventurer) RemoveGold(amt int) int {
	//adventurer loses amt gold		returns remainder
	a.gold -= amt
	return a.gold
}

func (a *adventurer) AddItem(itemName string) string {
	a.items = append(a.items, itemName)
	return itemName
}

func (a *adventurer) BuyItem(itemName string) string {
	if itemName == "sword" && a.gold >= 8 {
		a.RemoveGold(8)
		return a.AddItem("sword")
	}
	return "big fat sack of nothing"
}

func GoHomeCountMoney(h Human) {
	log("The Human has arrived in " + h.GoHome())
	log("The Human has " + strconv.Itoa(h.Gold()) + "GP")
}

func main() {
	dirk := &adventurer{GoldHaver: GoldHaver{gold: 10}, home: "town", location: "goblin caves"}
	shillelagh := &shopkeeper{GoldHaver: GoldHaver{gold: 1000}, shop: "Shilling Staves", items: map[string]bool{"Longstaff": true, "Blackstaff": true}}
	log("Dirk the Adventurer has arrived in " + dirk.GoHome())
	log("Dirk the Adventurer has " + strconv.Itoa(dirk.GoldHaver.Gold()) + "GP")
	log("Dirk the Adventurer has bought a " + dirk.BuyItem("sword"))
	log("Dirk the Adventurer has " + strconv.Itoa(dirk.Gold()) + "GP")
	log("Dirk the Adventurer has left town")
	log("Dirk the Adventurer has arrived at the mouth of the caves")
	log("Dirk the Adventurer has encounted 2 goblins")
	log("Dirk the Adventurer has " + dirk.SwingWeapon())
	log("Dirk the Adventurer has killed a goblin")
	log("Dirk the Adventurer has taken 1 damage")
	log("Dirk the Adventurer has killed a goblin")
	log("Dirk the Adventurer has entered the cave")
	log("Dirk the Adventurer has picked up a gemstone")
	log("Dirk the Adventurer has completed objective 'retrive gemstone'")
	log("Dirk the Adventurer has left the cave")
	log("Dirk the Adventurer has left the cave mouth")
	log("Dirk the Adventurer has arrived in town")
	log("Dirk the Adventurer has sold a sword")
	log("Dirk the Adventurer has handed in quest 'deal with the goblin raiders'")
	log("Dirk the Adventurer has received a scroll of healing")
	log("Shillelagh the Shoppekeep has " + strconv.Itoa(shillelagh.Gold()) + "GP")
	log("Shillelagh the Shoppekeep has sold " + shillelagh.SellItem("Blackstaff"))
	log("Shillelagh the Shoppekeep has " + strconv.Itoa(shillelagh.Gold()) + "GP")
	GoHomeCountMoney(shillelagh)
	GoHomeCountMoney(dirk)

}
