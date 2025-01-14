package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Truck struct {
	ID    int
	Start int
	End   int
	Cap   int
}

// на всякого мудреца довольно простоты :) достаточно двух сортировок
func solve(arrival []int, trucks []Truck) []int {
	ans := make([]int, len(arrival))

	arrivalIndex := make([]int, len(arrival))
	for i := range arrivalIndex {
		arrivalIndex[i] = i
	}

	sort.Slice(arrivalIndex, func(i, j int) bool {
		ai, aj := arrivalIndex[i], arrivalIndex[j]
		return arrival[ai] < arrival[aj]
	})

	sort.Slice(trucks, func(i, j int) bool {
		return trucks[i].Start < trucks[j].Start ||
			(trucks[i].Start == trucks[j].Start && trucks[i].ID < trucks[j].ID)
	})

	if debugEnable {
		log.Println("arrival     :", arrival)
		log.Println("arrivalIndex:", arrivalIndex)
		log.Println("trucks:", trucks)
	}

	for _, ai := range arrivalIndex {
		t := arrival[ai]
		if debugEnable {
			log.Println("ai:", ai, "t:", t)
		}

		for len(trucks) > 0 && (trucks[0].End < t || trucks[0].Cap == 0) {
			if debugEnable {
				log.Println("-", trucks[0].ID)
			}
			trucks = trucks[1:]
		}

		if len(trucks) == 0 || !(trucks[0].Start <= t && t <= trucks[0].End) {
			if debugEnable {
				log.Println(ai, "<-", -1)
			}
			ans[ai] = -1
			continue
		}

		if debugEnable {
			log.Println(ai, "<-", trucks[0].ID)
		}
		ans[ai] = trucks[0].ID
		trucks[0].Cap--
	}

	return ans
}

func run(in io.Reader, out io.Writer) {
	br := bufio.NewReader(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()
	var t int

	if _, err := fmt.Fscanln(br, &t); err != nil {
		panic(err)
	}

	for i := 1; i <= t; i++ {
		var n int
		if _, err := fmt.Fscan(br, &n); err != nil {
			panic(err)
		}

		arrival := make([]int, 0, n)
		for i := 0; i < n; i++ {
			var av int
			if _, err := fmt.Fscan(br, &av); err != nil {
				panic(err)
			}
			arrival = append(arrival, av)
		}

		var m int
		if _, err := fmt.Fscan(br, &m); err != nil {
			panic(err)
		}

		if debugEnable {
			log.Println("m:", m)
		}

		trucks := make([]Truck, 0, m)
		for id := 1; id <= m; id++ {
			var start, end, capacity int
			if _, err := fmt.Fscan(br, &start, &end, &capacity); err != nil {
				panic(err)
			}
			trucks = append(trucks, Truck{ID: id, Start: start, End: end, Cap: capacity})
		}

		ans := solve(arrival, trucks)

		bw.WriteString(strconv.Itoa(ans[0]))
		for _, v := range ans[1:] {
			bw.WriteByte(' ')
			bw.WriteString(strconv.Itoa(v))
		}
		bw.WriteByte('\n')
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}
