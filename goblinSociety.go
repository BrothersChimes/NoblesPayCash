package main

import (
	"fmt"
)

func log(output	string) {
	fmt.Println(output)
}

type GoblinSociety struct {
	totalGoblins int
	supplylist	 []string
}

func (gS *GoblinSociety) GainItem(itemName string) {
	gS.supplylist = append(gS.supplylist, itemName)
	return
}

func (gS *GoblinSociety) HasItem(itemName string) bool {
	return false
}

func main() {
	Grump := &GoblinSociety
	for i := 1;  i < 5; i++ {
		// main world loop
		log("a day has passed")
		
	}
}
