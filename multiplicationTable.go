package main

func MultiplicationTable(width int, height int) *[][]int {
	 table := make([][]int, width)

	 for w := range table {
	 	table[w] = make([]int, height)

	 	for h := range table[w] {
	 		table[w][h] = (w + 1) * (h + 1)
		}
	 }

	 return &table
}
