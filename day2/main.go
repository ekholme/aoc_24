package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const inp = "day2/input.txt"

type Report struct {
	Levels []int
}

func main() {
	reports, err := readData(inp)

	if err != nil {
		fmt.Println("Error reading data: ", err)
	}

	p1 := checkAllSafety(reports)

	fmt.Println("answer to p1: ", p1)

	//p2
	p2 := checkAllSafety2(reports)
	fmt.Println("answer to p2: ", p2)
}

func readData(path string) ([]*Report, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var reports []*Report

	for scanner.Scan() {
		line := scanner.Text()
		res := strings.Fields(line)

		var s []int
		for i := 0; i < len(res); i++ {
			v, err := strconv.Atoi(res[i])
			if err != nil {
				return nil, err
			}
			s = append(s, v)
		}

		r := &Report{
			Levels: s,
		}

		reports = append(reports, r)
	}

	return reports, nil
}

func checkAscending(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] > s[i] {
			return false
		}
	}
	return true
}

func checkDescending(s []int) bool {
	for i := 1; i < len(s); i++ {
		if s[i-1] < s[i] {
			return false
		}
	}
	return true

}

func checkSafety(report *Report) bool {
	l := report.Levels

	if !(checkAscending(l) || checkDescending(l)) {
		return false
	}

	for i := 1; i < len(l); i++ {
		d := Abs(l[i] - l[i-1])
		if d < 1 || d > 3 {
			return false
		}
	}

	return true
}

// reimplemeting absolute value for ints
func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func checkAllSafety(reports []*Report) int {
	var s = 0

	for i := 0; i < len(reports); i++ {
		if checkSafety(reports[i]) {
			s += 1
		}
	}

	return s
}

// part 2 ---------------
func checkSafety2(report *Report) bool {
	l := report.Levels

	for i := 0; i < len(l); i++ {
		l2 := append([]int{}, l[:i]...)
		l2 = append(l2, l[i+1:]...)
		r2 := &Report{
			Levels: l2,
		}
		if checkSafety(r2) {
			return true
		}
	}

	return false
}

func checkAllSafety2(reports []*Report) int {
	var s = 0

	for i := 0; i < len(reports); i++ {
		if checkSafety2(reports[i]) || checkSafety(reports[i]) {
			s += 1
		}
	}

	return s
}
