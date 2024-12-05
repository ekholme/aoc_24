package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const inp = "day3/input.txt"

// i should probably clean this up to reduce the duplication, but w/e
func main() {
	bs, err := os.ReadFile(inp)

	if err != nil {
		log.Fatalf("Error reading file: %s", err)
	}

	matches := detect(bs)

	res, err := multiplyAllNums(matches)

	if err != nil {
		log.Fatalf("Error extracting and multiplying numbers: %s", err)
	}

	p1 := 0

	for _, num := range res {
		p1 += num
	}

	fmt.Println("solution to part1: ", p1)

	//p2 ----------
	m2 := detect2(bs)

	res2, err := multiplyAllNums(m2)

	if err != nil {
		log.Fatalf("Error extracting and multiplying numbers: %s", err)
	}

	p2 := 0

	for _, num := range res2 {
		p2 += num
	}

	fmt.Println("solution to part2: ", p2)

}

func detect(inp []byte) [][]byte {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	matches := re.FindAll(inp, -1)

	return matches
}

func multiplyNums(b []byte) (int, error) {
	re := regexp.MustCompile(`\d{1,3},\d{1,3}`)

	match := re.Find(b)

	s := string(match)

	nums := strings.Split(s, ",")

	n1, err := strconv.Atoi(nums[0])

	if err != nil {
		return 0, err
	}

	n2, err := strconv.Atoi(nums[1])

	if err != nil {
		return 0, err
	}

	res := n1 * n2

	return res, nil
}

func multiplyAllNums(bs [][]byte) ([]int, error) {
	var res []int
	for i := 0; i < len(bs); i++ {
		v, err := multiplyNums(bs[i])
		if err != nil {
			return nil, err
		}
		res = append(res, v)
	}

	return res, nil

}

// part2 ---------------

func detectFirst(b []byte) []byte {
	re := regexp.MustCompile(`(?s:.*?)don\'t\(\)`)

	m := re.Find(b)

	return m
}

func detect2(b []byte) [][]byte {
	pref := "do()"

	pb := []byte(pref)

	b = append(pb, b...)

	re := regexp.MustCompile(`do\(\)(?s:.*?)don\'t\(\)`)

	matches := re.FindAll(b, -1)

	//flatten slice of slice back into a slice
	var rs []byte
	for _, m := range matches {
		rs = append(rs, m...)
	}

	//then run the regular detect
	res := detect(rs)

	return res
}
