package main

import (
	_ "embed"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type monkey struct {
	starting  []uint64
	operation func(item uint64) uint64
	test      func(item uint64) bool
	testint   uint64
	t         int
	f         int
	activity  int
}

func main() {
	input := strings.ReplaceAll(input, "\n\n", "\n")
	inputsp := strings.Split(input, "\n")
	monkey_count := len(inputsp) / 6
	var monkies = make([]monkey, monkey_count)
	var monkiesone = make([]monkey, monkey_count)
	for i := 0; i < monkey_count; i++ {
		s := i * 6
		// Start with s+5, as this proves s+(below 5) exists
		f := strings.Split(inputsp[s+5], "to monkey ")[1]
		fnumber, err := strconv.Atoi(f)
		if err != nil {
			panic("Couldn't find out what happens if false")
		}
		itemss := strings.Split(inputsp[s+1], ": ")[1]
		items := strings.Split(itemss, ", ")
		items_count := len(items)
		var starting = make([]uint64, items_count)
		for j, c := range strings.Split(itemss, ", ") {
			number, err := strconv.ParseUint(c, 10, 64)
			if err != nil {
				panic("bad")
			}
			starting[j] = number
		}
		ostring := strings.Split(inputsp[s+2], "= ")[1]
		osp := strings.Split(ostring, " ")
		onum := osp[2]
		var operation func(item uint64) uint64
		if onum == "old" {
			operation = func(item uint64) uint64 {
				return item * item
			}
		} else {
			onumber, err := strconv.ParseUint(onum, 10, 64)
			if err != nil {
				panic("Couldn't find operation")
			}
			if osp[1] == "*" {
				operation = func(item uint64) uint64 {
					return item * onumber
				}
			} else {
				operation = func(item uint64) uint64 {
					return item + onumber
				}
			}
		}
		divisible := strings.Split(inputsp[s+3], "by ")[1]
		number, err := strconv.ParseUint(divisible, 10, 64)
		//fmt.Printf("Monkey %d tests if worry level is divisible by %d\n", i, number)
		if err != nil {
			panic("Couldn't find test")
		}
		test := func(item uint64) bool {
			//fmt.Printf("Is %d mod %d zero? %t", item, number, (item%number) == 0)
			return (item % number) == 0
		}
		t := strings.Split(inputsp[s+4], "to monkey ")[1]
		tnumber, err := strconv.Atoi(t)
		if err != nil {
			panic("Couldn't find out what happens if true")
		}
		//fmt.Printf("If true, monkey %d throws to %d\n", i, tnumber)
		ape := monkey{starting: starting, operation: operation, test: test, testint: number, t: tnumber, f: fnumber}
		monkies[i] = ape
	}
	copy(monkiesone, monkies)
	res := solveOne(monkiesone, 20)
	resTwo := solve(monkies, 10000)
	activities := make([]int, monkey_count)
	for i, ape := range res {
		activities[i] = ape.activity
	}
	activitiestwo := make([]int, monkey_count)
	for i, ape := range resTwo {
		activitiestwo[i] = ape.activity
	}
	sort.Ints(activities)
	sort.Ints(activitiestwo)
	fmt.Printf("Part 1: %d\n", activities[monkey_count-1]*activities[monkey_count-2])
	fmt.Printf("Part 2: %d\n", activitiestwo[monkey_count-1]*activitiestwo[monkey_count-2])
}

func solve(monkies []monkey, rounds int) []monkey {
	for i := 0; i < rounds; i++ {
		for j, ape := range monkies {
			for _, item := range ape.starting {
				var worry uint64
				monkies[j].activity++
				var product uint64 = 1
				for _, monkey := range monkies {
					product *= monkey.testint
				}
				worry = ape.operation(item) % product
				if ape.test(worry) {
					//fmt.Printf("Monkey %d threw item with level %d to monkey %d as test was true\n", j, worry, ape.t)
					monkies[ape.t].starting = append(monkies[ape.t].starting, worry)
				} else {
					//fmt.Printf("Monkey %d threw item with level %d to monkey %d as test was false\n", j, worry, ape.f)
					monkies[ape.f].starting = append(monkies[ape.f].starting, worry)
				}
			}
			monkies[j].starting = make([]uint64, 0)
		}
	}
	return monkies
}

func solveOne(monkies []monkey, rounds int) []monkey {
	for i := 0; i < rounds; i++ {
		for j, ape := range monkies {
			for _, item := range ape.starting {
				var worry uint64
				monkies[j].activity++
				worryfloat := math.Floor(float64(ape.operation(item)) / 3.0)
				worry = uint64(worryfloat)
				if ape.test(worry) {
					//fmt.Printf("Monkey %d threw item with level %d to monkey %d as test was true\n", j, worry, ape.t)
					monkies[ape.t].starting = append(monkies[ape.t].starting, worry)
				} else {
					//fmt.Printf("Monkey %d threw item with level %d to monkey %d as test was false\n", j, worry, ape.f)
					monkies[ape.f].starting = append(monkies[ape.f].starting, worry)
				}
			}
			monkies[j].starting = make([]uint64, 0)
		}
	}
	return monkies
}
