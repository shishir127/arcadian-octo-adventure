package life

import (
	"bufio"
	"os"
	"strings"
)

func ReadGridFromConfig(configFile string) ([][]cell, error) {
	file, err := os.Open(configFile)
	if err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		grid := make([][]cell, 0)
		for scanner.Scan() {
			grid = append(grid, parseGridRow(scanner.Text()))
		}
		if err := scanner.Err(); err != nil {
			return make([][]cell, 0), err
		}
		return grid, nil
	}
	return make([][]cell, 0), err
}

func parseGridRow(row string) []cell {
	cellRow := make([]cell, 0)
	for _, s := range strings.Split(row, ", ") {
		switch strings.Trim(s, " ") {
		case "alive":
			cellRow = append(cellRow, cell{status: true})
		case "dead":
			cellRow = append(cellRow, cell{status: false})
		}
	}
	return cellRow
}

func nextGeneration(grid [][]cell) [][]cell {
	for row, rows := range grid {
		for column, cell := range rows {
			if cell.isAlive() {
				alertNeighboursAboutLiveCell(grid, row, column, len(grid), len(rows))
			}
		}
	}
	nextGenerationGrid := make([][]cell, 0)
	for _, rows := range grid {
		nextGenerationRow := make([]cell, 0)
		for _, cell := range rows {
			nextGenerationRow = append(nextGenerationRow, cell.nextGeneration())
		}
		nextGenerationGrid = append(nextGenerationGrid, nextGenerationRow)
	}
	return nextGenerationGrid
}

func alertNeighboursAboutLiveCell(grid [][]cell, row, column, numRows, numColumns int) {
	for i := row - 1; i <= row+1; i++ {
		if i < 0 || i >= numRows {
			continue
		}
		for j := column - 1; j < numColumns; j++ {
			if j < 0 || j >= numColumns {
				continue
			}
			if i == row && j == column {
				continue
			}
			grid[i][j].registerLiveNeighbour()
		}
	}
}
