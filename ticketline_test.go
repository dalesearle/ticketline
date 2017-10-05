package main

import "testing"

var testData = map[string]struct {
	line     []int
	position int
	expected int
}{
	"1": {[]int{1}, 0, 1},
	"2": {[]int{2}, 0, 2},
	"3": {[]int{1,2}, 1, 3},
	"4": {[]int{2,1}, 0, 3},
	"5": {[]int{1,2,1}, 2, 3},
	"6": {[]int{1,2,1}, 1, 4},
	"7": {[]int{1,2,2}, 1, 4},
	"8": {[]int{1,2,3}, 1, 4},
	"9": {[]int{1,4,5,7,1}, 2, 15},
}

func  TestEfficientSolution(t *testing.T) {
	for key, data := range testData {
		actual := getEfficientTicketCount(data.line,data.position)
		if actual != data.expected {
			t.Errorf("[%s] expected %d, got %d", key, data.expected, actual)
		}
	}
}

func  TestLiteralSolution(t *testing.T) {
	for key, data := range testData {
		actual := getLiteralTicketCount(data.line,data.position)
		if actual != data.expected {
			t.Errorf("[%s] expected %d, got %d", key, data.expected, actual)
		}
	}
}
