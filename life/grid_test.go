package life

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestReadGridFromConfig(t *testing.T) {
	const configFile = "/Users/shishir/side_projects/game_of_life_golang/src/life/grid_test.txt"
	_, err := ioutil.ReadFile(configFile)
	assert.Nil(t, err)
	row1 := make([]cell, 0)
	row1 = append(row1, cell{status: true}, cell{status: false})
	row2 := make([]cell, 0)
	row2 = append(row2, cell{status: false}, cell{status: false})
	testGrid := make([][]cell, 0)
	testGrid = append(testGrid, row1, row2)
	grid, gridError := ReadGridFromConfig(configFile)
	assert.Nil(t, gridError)
	assert.Equal(t, testGrid, grid)
}

func TestParseGridRow(t *testing.T) {
	testString := "alive, dead, alive, alive"
	testRow := make([]cell, 0)
	testRow = append(testRow, cell{status: true}, cell{status: false}, cell{status: true}, cell{status: true})
	gridRow := parseGridRow(testString)
	assert.Equal(t, testRow, gridRow)
}

func TestAlertNeighboursAboutLiveCell(t *testing.T) {
	const configFile = "/Users/shishir/side_projects/game_of_life_golang/src/life/grid_test.txt"
	grid, gridError := ReadGridFromConfig(configFile)
	assert.Nil(t, gridError)
	alertNeighboursAboutLiveCell(grid, 0, 0, 2, 2)
	assert.Equal(t, 1, grid[0][1].liveNeighbours)
	assert.Equal(t, 1, grid[1][0].liveNeighbours)
	assert.Equal(t, 1, grid[1][1].liveNeighbours)
}

func TestGridnextGeneration(t *testing.T) {
	const configFile = "/Users/shishir/side_projects/game_of_life_golang/src/life/grid_test.txt"
	grid, gridError := ReadGridFromConfig(configFile)
	assert.Nil(t, gridError)
	row1 := make([]cell, 0)
	row1 = append(row1, cell{status: false}, cell{status: false})
	row2 := make([]cell, 0)
	row2 = append(row2, cell{status: false}, cell{status: false})
	nextGrid := make([][]cell, 0)
	nextGrid = append(nextGrid, row1, row2)
	assert.Equal(t, nextGrid, nextGeneration(grid))
}
