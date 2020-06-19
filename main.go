package main

import "fmt"

func main() {
	width := 10
	table := MultiplicationTable(width, width)
	tableP := PrintTable(table)

	fmt.Printf("Mulitplication Table (%d):\n%s\n", width, tableP)
}
