package main

import (
	"bufio"
	"os"

	"github.com/brotherschimes/noblespaycash/input"
	"github.com/brotherschimes/noblespaycash/shop"
)

func main() {
	reader := input.Reader{Reader: bufio.NewReader(os.Stdin)}
	shop.DoSales(reader)
}
