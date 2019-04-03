package customer

import (
	"bufio"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type Reader struct {
	*bufio.Reader
}

func (r Reader) GetAnswer() bool {
	for true {
		fmt.Println("Please enter (y/n)")
		answer, _ := r.ReadString('\n')
		answerChar := unicode.ToUpper(rune(answer[0]))
		switch answerChar {
		case 'Y':
			return true
		case 'N':
			return false
		default:
			fmt.Println("That character was not understood.")
		}
	}
	return false
}

func (r Reader) GetSelection(maxSelect int) int {
	for true {
		fmt.Println("Please enter a number.")
		answer, _ := r.ReadString('\n')
		answerNum, err := strconv.Atoi(strings.TrimSpace(answer))
		if err != nil {
			log.Fatal(err)
			fmt.Println("That entry was not understood.")
			continue
		}

		if answerNum >= maxSelect {
			fmt.Println("That number is too high.")
			continue
		}

		if answerNum < 0 {
			fmt.Println("Please enter a non-negative number.")
			continue
		}

		return answerNum
	}
	return 0
}
