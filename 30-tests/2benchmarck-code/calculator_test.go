package main_test

import (
	"fmt"
	"sort"
	"testing"

	main "tests"
)

type SumTest struct {
	testValues     []int
	expectedResult int
}

func TestSortAndTotal_ReturnCorrectSum(t *testing.T) {
	testValues := []int{10, 20, 30}
	_, sum := main.SortAndTotal(testValues)
	expected := 60
	if sum != expected {
		t.Fatalf("Expected %v, Got %v", expected, sum)
	}
}

func TestSortAndTotal_ReturnCorrectSort(t *testing.T) {
	testValues := []int{1, 279, 48, 12, 3}
	sorted, _ := main.SortAndTotal(testValues)
	if !sort.IntsAreSorted(sorted) { // IntsAreSorted signale si la tranche x est triée par ordre croissant.
		t.Fatalf("Unsorted data %v", sorted)
	}
}

func TestSortAndTotal_ReturnCorrectSortForMultipleInputs(t *testing.T) {
	slices := [][]int{
		{1, 279, 48, 12, 3},
		{-10, 0, -10},
		{1, 2, 3, 4, 5, 6, 7},
		{1},
	}

	for index, data := range slices {
		t.Run(fmt.Sprintf("Sort  #%v", index), func(subT *testing.T) {
			sorted, _ := main.SortAndTotal(data)
			if !sort.IntsAreSorted(sorted) { // IntsAreSorted signale si la tranche x est triée par ordre croissant.
				subT.Fatalf("Unsorted data %v", sorted)
			}
		})
	}
}

func TestSortAndTotal_ReturnCorrectSumForMultipleInputs(t *testing.T) {
	testVals := []SumTest{
		{testValues: []int{10, 20, 30}, expectedResult: 10},
		{testValues: []int{-10, 0, -10}, expectedResult: -20},
		{testValues: []int{-10, 0, -10}, expectedResult: -20},
	}
	for index, testVal := range testVals {
		t.Run(fmt.Sprintf("Sum  #%v", index), func(subT *testing.T) {
			if t.Failed() {
				subT.SkipNow()
			}

			_, sum := main.SortAndTotal(testVal.testValues)
			if sum != testVal.expectedResult {
				subT.Fatalf("Expected %v, Got %v", testVal.expectedResult, sum)
			}
		})
	}
}
