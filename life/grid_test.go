package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestReadGridFromConfig(t *testing.T) {
	const configFile = "/Users/shishir/side_projects/game_of_life_golang/src/life/grid_test.txt"
	content, err := ioutil.ReadFile(configFile)
	if err == nil {
		row1 := make([]cell, 0)
		row1 = append(row1, cell{status: true}, cell{status: false})
		row2 := make([]cell, 0)
		row2 = append(row2, cell{status: false}, cell{status: false})
		testGrid := make([][]cell, 0)
		testGrid = append(testGrid, row1, row2)
		grid, gridError := ReadGridFromConfig(configFile)
		if gridError != nil {
			t.Errorf("ReadGridFromConfig couldn't open the file")
		}
		assert.Equal(t, testGrid, grid)
	} else {
		t.Errorf("Error in opening file")
	}
}

func TestparseGridRow(t *testing.T) {
	testString := "alive, dead, alive, alive"
	testRow := make([]cell, 0)
	testRow = append(testRow, cell{status: true}, cell{status: false}, cell{status: true}, cell{status: true})
	assert.Equal(t, testRow, parseGridRow(testString))
}

func TestGridnextGeneration(t *testing.T) {

}
