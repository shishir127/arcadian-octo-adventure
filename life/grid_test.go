package life

import (
	"io/ioutil"
	"testing"
  "github.com/stretchr/testify/assert"
)

func TestreadGridFromConfig(t *testing.T) {
	const configFile = "/Users/shishir/side_projects/game_of_life_golang/src/life/grid_test.txt"
	contents, err := ioutil.ReadFile(configFile)
	if err == nil {
		row1 := make([]cell, 2)
		row1 = append(row1, cell{status: true}, cell{status: false})
		row2 := make([]cell, 2)
		row2 = append(row2, cell{status: false}, cell{status: false})
		testGrid := make([][]cell, 2)
		testGrid = append(testGrid, row1, row2)
		assert.Equal(t, readGridFromConfig(configFile), testGrid)
	} else {
		t.Errorf("Some error in opening file")
	}
}

func TestGridnextGeneration(t *testing.T) {
}
