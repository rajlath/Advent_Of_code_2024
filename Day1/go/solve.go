package main

import (
	"fmt"

	"io/ioutil"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func getData() ([]int, []int) {
	var listA []int
	var listB []int
	fileContent, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(fileContent), "\n")
	re := regexp.MustCompile("[0-9]+")
	for _, line := range lines {
		abc := re.FindAllString(line, -1)

		A, err := strconv.Atoi(abc[0])
		if err != nil {
			panic(err)
		}
		listA = append(listA, A)

		B, err := strconv.Atoi(abc[1])
		if err != nil {
			panic(err)
		}
		listB = append(listB, B)

	}
	return listA, listB

}

func stringToInt(arr []string) []int {
	var ret = make([]int, len(arr))
	for idx, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		ret[idx] = j
	}
	return ret
}

func part1(list1 []int, list2 []int) int {
	distance := 0
	for idx, i := range list1 {
		distance += int(math.Abs(float64(i) - float64(list2[idx])))
	}
	return distance

}
func countValue(arr []int) map[int]int {
	//Create a   dictionary of values for each element
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	return dict
}

func part2(A []int, B []int) int {
	counts := countValue(B)

	simillarityScore := 0
	for _, vals := range A {
		if val, ok := counts[vals]; ok {

			simillarityScore += (vals * val)
		}
	}
	return simillarityScore
}

func main() {
	var listA, listB = getData()
	sort.Slice(listA, func(i, j int) bool {
		return listA[j] > listA[i]
	})
	sort.Slice(listB, func(i, j int) bool {
		return listB[j] > listB[i]
	})

	fmt.Println(part1(listA, listB)) // 1590491
	fmt.Println(part2(listA, listB)) // 22588371

}
