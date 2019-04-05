package main

import (
	"bufio"
	"os"

	"github.com/BrothersChimes/NoblesPayCash/input"
	"github.com/BrothersChimes/NoblesPayCash/shop"
)

func main() {
	reader := input.Reader{Reader: bufio.NewReader(os.Stdin)}
	shop.DoSales(reader)
}
