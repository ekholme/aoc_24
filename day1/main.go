package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Puzzle struct {
	FirstCol  int
	SecondCol int
}

type PuzzleCounts struct {
	Key  int
	Val  int
	Mult int
}

const inp = "day1/input.txt"

func main() {
	d, err := readData(inp)

	if err != nil {
		fmt.Println(err.Error())
	}

	pzls, err := splitRows(d)

	if err != nil {
		fmt.Println(err.Error())
	}

	c1, c2 := getCols(pzls)

	p1 := calcDist(c1, c2)

	fmt.Println(p1)

	//part 2
	pcs := countMap(c1, c2)

	p2 := sumMults(pcs)

	fmt.Println("Solution to part 2: ", p2)

}

// read the data in from a .txt file as a slice of strings
func readData(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println("Error opening file: ", err)
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file: ", err)
		return nil, err
	}

	return lines, nil
}

// split each row into ints and assign to a Puzzle struct
func splitRow(inp string) (*Puzzle, error) {
	res := strings.Fields(inp)

	v1, err := strconv.Atoi(res[0])

	if err != nil {
		return nil, err
	}

	v2, err := strconv.Atoi(res[1])

	if err != nil {
		return nil, err
	}

	out := &Puzzle{
		FirstCol:  v1,
		SecondCol: v2,
	}

	return out, nil
}

func splitRows(inp []string) ([]*Puzzle, error) {
	var puzzles []*Puzzle

	for i := 0; i < len(inp); i++ {
		p, err := splitRow(inp[i])

		if err != nil {
			return nil, err
		}

		puzzles = append(puzzles, p)
	}

	return puzzles, nil

}

func getCols(puzzles []*Puzzle) ([]int, []int) {
	var c1 []int
	var c2 []int

	for i := 0; i < len(puzzles); i++ {
		c1 = append(c1, puzzles[i].FirstCol)
		c2 = append(c2, puzzles[i].SecondCol)
	}

	sort.Ints(c1)
	sort.Ints(c2)

	return c1, c2
}

func calcDist(c1 []int, c2 []int) int {
	var s = 0

	for i := 0; i < len(c1); i++ {
		d := Abs(c1[i] - c2[i])
		s += d
	}

	return s
}

// apparently Go doesn't have an absolute value function for integers?
func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

// part 2 stuff -----------
func countMap(c1 []int, c2 []int) []*PuzzleCounts {

	counts := make(map[int]int)

	for _, v := range c2 {
		if _, ok := counts[v]; ok {
			counts[v]++
		} else {
			counts[v] = 1
		}
	}

	var ps []*PuzzleCounts

	for _, k := range c1 {
		m := k * counts[k]
		p := &PuzzleCounts{
			Key:  k,
			Val:  counts[k],
			Mult: m,
		}

		ps = append(ps, p)
	}

	return ps
}

func sumMults(pcs []*PuzzleCounts) int {
	var s = 0

	for i := 0; i < len(pcs); i++ {
		s += pcs[i].Mult
	}

	return s
}

//RESUME HERE -- SHOULD JUST NEED TO SUM THE MULTS TOMORROW
