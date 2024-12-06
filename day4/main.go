package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

const inp = "day4/input.txt"

type Puzzle struct {
	Lines   []string
	Letters [][]string
	Rows    int
	Cols    int
}

func main() {

	data, err := readData(inp)

	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	m1 := data.detectHorizontal("forward")
	m2 := data.detectHorizontal("backward")

	r := m1 + m2

	fmt.Println("Horizontal matches: ", r)

	m3 := data.detectVertical("forward")
	m4 := data.detectVertical("backward")

	v := m3 + m4

	fmt.Println("Vertical matches: ", v)
}

func readData(path string) (*Puzzle, error) {
	file, err := os.Open(path)

	if err != nil {
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
		return nil, err
	}

	//getting letters
	rows := len(lines)
	cols := utf8.RuneCountInString(lines[0])

	var letters [][]string

	for i := 0; i < rows; i++ {
		l := strings.Split(lines[i], "")
		letters = append(letters, l)
	}

	ret := &Puzzle{
		Lines:   lines,
		Letters: letters,
		Rows:    rows,
		Cols:    cols,
	}

	return ret, nil

}

// methods to detect
func (p *Puzzle) detectHorizontal(direction string) int {
	s := 0

	var r string

	if direction == "forward" {
		r = `XMAS`
	} else {
		r = `SAMX`
	}

	for i := 0; i < len(p.Lines); i++ {
		re := regexp.MustCompile(r)
		b := []byte(p.Lines[i])
		m := re.FindAll(b, -1)
		num_match := len(m)
		s += num_match
	}

	return s
}

func (p *Puzzle) detectVertical(direction string) int {
	tl := p.transposeLetters()

	var lines []string

	for _, l := range tl {
		line := strings.Join(l, "")
		lines = append(lines, line)
	}

	//i should just re-implement my detect function so i don't have to copy it here, but w/e
	s := 0

	var r string

	if direction == "forward" {
		r = `XMAS`
	} else {
		r = `SAMX`
	}

	for i := 0; i < len(p.Lines); i++ {
		re := regexp.MustCompile(r)
		b := []byte(lines[i])
		m := re.FindAll(b, -1)
		num_match := len(m)
		s += num_match
	}

	return s

}

func (p *Puzzle) transposeLetters() [][]string {
	result := make([][]string, p.Cols)

	for i := range result {
		result[i] = make([]string, p.Rows)
	}

	for i := 0; i < p.Rows; i++ {
		for j := 0; j < p.Cols; j++ {
			result[j][i] = p.Letters[i][j]
		}
	}

	return result
}
