package main

import "fmt"

func PrintTable(table *[][]int) string {
	str := ""

	for _, row := range *table {
		for _, value := range row {
			str = fmt.Sprintf("%s%4d ",str,  value)
		}
		str = fmt.Sprint(str, "\n")
	}

	return str
}
