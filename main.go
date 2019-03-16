package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

const uniqueCustName = "Ulric"

func main() {
	reader := bufio.NewReader(os.Stdin)

	var numSwords = 1
	var numTrailMix = 1
	announceItemQty("swords", numSwords)
	announceItemQty("bags of trail mix", numTrailMix)

	says(uniqueCustName, "Greetings!")
	says(uniqueCustName, "I would like to purchase a sword!")

	sellSwords(reader, &numSwords)

	says(uniqueCustName, "I would still like to purchase a sword!")
	announceItemQty("swords", numSwords)
	announceItemQty("bag of trail mix", numSwords)

	sellSwords(reader, &numSwords)
	announceItemQty("swords", numSwords)
	announceItemQty("bag of trail mix", numSwords)

}

func sellSwords(reader *bufio.Reader, numSwords *int) {
	if *numSwords <= 0 {
		fmt.Println("You inform Ulric that you have no swords left for sale.")
		return
	}

	fmt.Println("Would you like to sell Ulric a sword? (y/n)")

	answer, _ := reader.ReadString('\n')
	var answerChar = unicode.ToUpper(rune(answer[0]))

	switch answerChar {
	case 'Y':
		does(uniqueCustName, "happily takes the sword!")
		*numSwords--
	case 'N':
		does(uniqueCustName, "is sad you did not sell him the sword.")
	default:
		does(uniqueCustName, "did not understand what you said.")
	}
}

func announceItemQty(itemNamePlural string, numSwords int) {
	fmt.Printf("You have %v %s.\n", numSwords, itemNamePlural)
}

func does(name, action string) {
	fmt.Println(name + " " + action)
}

func says(name, speech string) {
	fmt.Println(name + " says: " + "\"" + speech + ".\"")
}
