package life

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIsAlive(t *testing.T) {
	cell1 := cell{status: true}
	status1 := cell1.isAlive()
	assert.Equal(t, true, status1)
	cell2 := cell{status: false}
	status2 := cell2.isAlive()
	assert.Equal(t, false, status2)
}

func TestRegisterLiveNeighbour(t *testing.T) {
	cell1 := cell{status: true}
	previousLiveNeighbours := cell1.liveNeighbours
	cell1.registerLiveNeighbour()
	currentLiveNeighbours := cell1.liveNeighbours
	assert.Equal(t, previousLiveNeighbours+1, currentLiveNeighbours)
}

func TestCellNextGeneration(t *testing.T) {
	deadCell := cell{status: false, liveNeighbours: 3}
	nextGeneration1 := deadCell.nextGeneration()
	assert.Equal(t, true, nextGeneration1.isAlive())
	liveCell1 := cell{status: true, liveNeighbours: (rand.Intn(2))}
	nextGeneration2 := liveCell1.nextGeneration()
	assert.Equal(t, false, nextGeneration2.isAlive())
	liveCell2 := cell{status: true, liveNeighbours: (rand.Intn(2) + 2)}
	nextGeneration3 := liveCell2.nextGeneration()
	assert.Equal(t, true, nextGeneration3.isAlive())
	liveCell3 := cell{status: true, liveNeighbours: (rand.Intn(5) + 4)}
	nextGeneration4 := liveCell3.nextGeneration()
	assert.Equal(t, false, nextGeneration4.isAlive())
}
