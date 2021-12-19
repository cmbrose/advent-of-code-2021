package main

import (
	"container/heap"
	"fmt"
	"math"
	"sort"

	"../util"
)

type Step struct {
	X    int
	Y    int
	Cost int
}

type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}

type StepPriorityQueue []Step

func (h StepPriorityQueue) Len() int           { return len(h) }
func (h StepPriorityQueue) Less(i, j int) bool { return h[i].Cost < h[j].Cost }
func (h StepPriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *StepPriorityQueue) Push(x interface{}) {
	*h = append(*h, x.(Step))
}

func (h *StepPriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func main() {
	grid := util.ParseIntGrid()

	costs := makeCostGrid(len(grid[0]), len(grid))

	util.PrintIntGrid(grid)

	targetY := len(grid) - 1
	targetX := len(grid[targetY]) - 1

	queue := &StepPriorityQueue{}
	queue.Push(Step{X: 0, Y: 0, Cost: 0})

	heap.Init(queue)

	for queue.Len() > 0 {
		curr := heap.Pop(queue).(Step)

		// Already have a cheaper path
		if costs[curr.Y][curr.X] <= curr.Cost {
			continue
		}

		costs[curr.Y][curr.X] = curr.Cost

		fmt.Printf("At (%d,%d) with cost %d\n", curr.X, curr.Y, curr.Cost)

		if curr.X == targetX && curr.Y == targetY {
			fmt.Println(curr.Cost)
			return
		}

		for i := -1; i <= 1; i++ {
			if curr.X+i >= 0 && curr.X+i < len(grid[curr.Y]) && i != 0 {
				step := Step{curr.X + i, curr.Y, curr.Cost + grid[curr.Y][curr.X+i]}
				heap.Push(queue, step)
			}

			if curr.Y+i >= 0 && curr.Y+i < len(grid) && i != 0 {
				step := Step{curr.X, curr.Y + i, curr.Cost + grid[curr.Y+i][curr.X]}
				heap.Push(queue, step)
			}
		}
	}
}

func makeCostGrid(x, y int) [][]int {
	grid := make([][]int, y)

	for i := 0; i < y; i++ {
		grid[i] = make([]int, x)

		for j := 0; j < x; j++ {
			grid[i][j] = math.MaxInt
		}
	}

	return grid
}
