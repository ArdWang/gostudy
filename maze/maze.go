package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][] int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int

	fmt.Fscanf(file,"%d %d", &row, &col)

	maze := make([][] int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", maze[i][j])
		}
	}

	return maze

}

type ponit struct {
	i, j int
}

var dirs = [4]ponit {
	{-1,0},{0,-1},{1,0},{0,1},
}

func (p ponit) add(r ponit) ponit{
	return ponit{p.i+r.i,p.j+r.j}
}

func (p ponit) at(grid [][] int) (int,bool) {
	return grid[p.i][p.j]
}

func walk(maze [][] int, start, end ponit) {
	steps := make([][] int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}
	Q := [] ponit{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0
			// and next != start
		}
	}
}


func main() {
	 maze := readMaze("maze/maze.in")
	 for _, row := range maze{
	 	for _, val := range row {
	 		fmt.Printf("%d", val)
		}
		fmt.Println()
	 }

	 walk(maze, ponit{0,0}, ponit{len(maze)-1, len(maze[0])-1})
}
