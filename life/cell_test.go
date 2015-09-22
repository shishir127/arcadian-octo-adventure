package life

import (
	"math/rand"
	"testing"
)

func TestIsAlive(t *testing.T) {
	cell1 := cell{status: true}
	status1 := cell1.isAlive()
	if status1 != true {
		t.Errorf("Failed TestingisAlive for a live cell - %b", status1)
	}
	cell2 := cell{status: false}
	status2 := cell2.isAlive()
	if status2 != false {
		t.Errorf("Failed TestingisAlive for a dead cell - %b", status2)
	}
}

func TestRegisterLiveNeighbour(t *testing.T) {
	cell1 := cell{status: true}
	previousLiveNeighbours := cell1.liveNeighbours
	cell1.registerLiveNeighbour()
	currentLiveNeighbours := cell1.liveNeighbours
	if previousLiveNeighbours+1 != currentLiveNeighbours {
		t.Errorf("Failed TestingregisterLiveNeighbour")
	}
}

func TestCellNextGeneration(t *testing.T) {
	deadCell := cell{status: false, liveNeighbours: 3}
	nextGeneration1 := deadCell.nextGeneration()
	if !nextGeneration1.isAlive() {
		t.Errorf("Failed TestingnextGeneration for the dead->alive transition")
	}
	liveCell1 := cell{status: true, liveNeighbours: (rand.Intn(2))}
	nextGeneration2 := liveCell1.nextGeneration()
	if nextGeneration2.isAlive() {
		t.Errorf("Failed TestingnextGeneration for the alive->dead by under-population transition")
	}
	liveCell2 := cell{status: true, liveNeighbours: (rand.Intn(2) + 2)}
	nextGeneration3 := liveCell2.nextGeneration()
	if !nextGeneration3.isAlive() {
		t.Errorf("Failed TestingnextGeneration for the alive->alive transition")
	}
	liveCell3 := cell{status: true, liveNeighbours: (rand.Intn(5) + 4)}
	nextGeneration4 := liveCell3.nextGeneration()
	if nextGeneration4.isAlive() {
		t.Errorf("Failed TestingnextGeneration for the alive->dead by overcrowding transition")
	}
}
