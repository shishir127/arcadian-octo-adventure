package main

import (
	"bufio"
	"os"
	"strings"
)

func createGrid(rows, columns int) {

}

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

func nextGeneration() {

}
