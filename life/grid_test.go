package life

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestReadGridFromConfig(t *testing.T) {
	const configFile = "/Users/shishir/side_projects/game_of_life_golang/src/life/grid_test.txt"
	_, err := ioutil.ReadFile(configFile)
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

func TestParseGridRow(t *testing.T) {
	testString := "alive, dead, alive, alive"
	testRow := make([]cell, 0)
	testRow = append(testRow, cell{status: true}, cell{status: false}, cell{status: true}, cell{status: true})
	gridRow := parseGridRow(testString)
	if !reflect.DeepEqual(testRow, gridRow) {
		t.Errorf("Failed TestparseGridRow")
	}
}

func TestAlertNeighboursAboutLiveCell(t *testing.T) {
	const configFile = "/Users/shishir/side_projects/game_of_life_golang/src/life/grid_test.txt"
	grid, gridError := ReadGridFromConfig(configFile)
	if gridError == nil {
		alertNeighboursAboutLiveCell(grid, 0, 0, 2, 2)
		assert.Equal(t, 1, grid[0][1].liveNeighbours)
		assert.Equal(t, 1, grid[1][0].liveNeighbours)
		assert.Equal(t, 1, grid[1][1].liveNeighbours)
	} else {
		t.Errorf("Error in opening file")
	}
}

func TestGridnextGeneration(t *testing.T) {
	const configFile = "/Users/shishir/side_projects/game_of_life_golang/src/life/grid_test.txt"
	grid, gridError := ReadGridFromConfig(configFile)
	if gridError == nil {
		row1 := make([]cell, 0)
		row1 = append(row1, cell{status: false}, cell{status: false})
		row2 := make([]cell, 0)
		row2 = append(row2, cell{status: false}, cell{status: false})
		nextGrid := make([][]cell, 0)
		nextGrid = append(nextGrid, row1, row2)
		if !reflect.DeepEqual(nextGrid, nextGeneration(grid)) {
			t.Errorf("Error in TestGridnextGeneration")
		}
	} else {
		t.Errorf("Error in opening file")
	}
}
