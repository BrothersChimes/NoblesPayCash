package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {
	says("Ulric", "Greetings!")
	says("Ulric", "I would like to purchase a sword!")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Would you like to sell Ulric a sword? (y/n)")
	text, _ := reader.ReadString('\n')
	var answerChar = unicode.ToUpper(rune(text[0]))
	switch answerChar {
	case 'Y':
		does("Ulric", "happily takes the sword!")
	case 'N':
		does("Ulric", "is sad you did not sell him the sword.")
	default:
		does("Ulric", "did not understand what you said.")
	}
}

func does(name, action string) {
	fmt.Println(name + " " + action)
}

func says(name, speech string) {
	fmt.Println(name + " says: " + "\"" + speech + ".\"")
}
