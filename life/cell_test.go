package life

import "testing"
import "math/rand"

func TestIsAlive(t *testing.T) {
	cell1 := cell{status: true}
	status1 := cell1.IsAlive()
	if status1 != true {
		t.Errorf("Failed TestingIsAlive for a live cell - %b", status1)
	}
	cell2 := cell{status: false}
	status2 := cell2.IsAlive()
	if status2 != false {
		t.Errorf("Failed TestingIsAlive for a dead cell - %b", status2)
	}
}

func TestRegisterLiveNeighbour(t *testing.T) {
	cell1 := cell{status: true}
	previousLiveNeighbours := cell1.liveNeighbours
	cell1.RegisterLiveNeighbour()
	currentLiveNeighbours := cell1.liveNeighbours
	if previousLiveNeighbours+1 != currentLiveNeighbours {
		t.Errorf("Failed TestingRegisterLiveNeighbour")
	}
}

func TestNextGeneration(t *testing.T) {
	deadCell := cell{status: false, liveNeighbours: 3}
	nextGeneration1 := deadCell.NextGeneration()
	if !nextGeneration1.IsAlive() {
		t.Errorf("Failed TestingNextGeneration for the dead->alive transition")
	}
	liveCell1 := cell{status: true, liveNeighbours: (rand.Intn(2))}
	nextGeneration2 := liveCell1.NextGeneration()
	if nextGeneration2.IsAlive() {
		t.Errorf("Failed TestingNextGeneration for the alive->dead by under-population transition")
	}
	liveCell2 := cell{status: true, liveNeighbours: (rand.Intn(2) + 2)}
	nextGeneration3 := liveCell2.NextGeneration()
	if !nextGeneration3.IsAlive() {
		t.Errorf("Failed TestingNextGeneration for the alive->alive transition")
	}
	liveCell3 := cell{status: true, liveNeighbours: (rand.Intn(5) + 4)}
	nextGeneration4 := liveCell3.NextGeneration()
	if nextGeneration4.IsAlive() {
		t.Errorf("Failed TestingNextGeneration for the alive->dead by overcrowding transition")
	}
}
