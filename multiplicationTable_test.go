package main

import "testing"

func TestMultiplicationTable(t *testing.T) {
	table1x1 := MultiplicationTable(1, 1)
	value := (*table1x1)[0][0]
	if value != 1 {
		t.Errorf("Expected (*table1x1)[0][0] == 1 but was %d", value)
	}

	table10x10 := MultiplicationTable(10, 10)
	value = (*table10x10)[9][9]
	if value != 100 {
		t.Errorf("Expected (*table10x10)[9][9] == 100 but was %d", value)
	}
}
