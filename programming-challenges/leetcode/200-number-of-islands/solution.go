package main

import (
	"container/list"
	"fmt"
)

const (
	land  byte = '1'
	water byte = '0'
)

type Coordenate struct {
	X int
	Y int
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}

	result := numIslands(grid)
	fmt.Println(result)
}

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	// visited matrix
	visited := make([][]bool, rows)
	for i := range visited {
		visited[i] = make([]bool, cols)
	}

	islandCount := 0

	for y := range rows {
		for x := range cols {
			topography := grid[y][x]
			if topography == land && !visited[y][x] {
				startCoordenate := Coordenate{
					X: x,
					Y: y,
				}
				bfs(startCoordenate, grid, visited)
				islandCount++
			}
		}
	}

	return islandCount
}

func bfs(start Coordenate, grid [][]byte, visited [][]bool) {
	rows := len(grid)
	cols := len(grid[0])

	dx := [4]int{-1, 0, 1, 0}
	dy := [4]int{0, 1, 0, -1}

	queue := list.New()
	queue.PushBack(start)

	for queue.Len() > 0 {
		visitElement := queue.Front()
		coordenate := visitElement.Value.(Coordenate)
		topography := grid[coordenate.Y][coordenate.X]
		queue.Remove(visitElement)

		if topography == water || visited[coordenate.Y][coordenate.X] {
			continue
		}

		visited[coordenate.Y][coordenate.X] = true
		for i := range dx {
			nextX := coordenate.X + dx[i]
			nextY := coordenate.Y + dy[i]
			if nextX >= 0 && nextX < cols && nextY >= 0 && nextY < rows {
				nextCoordenate := Coordenate{
					X: nextX,
					Y: nextY,
				}
				queue.PushBack(nextCoordenate)
			}
		}
	}
}
