package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

type Truck struct {
	ID        int
	Start     int
	End       int
	Cap       int
	starIndex int
	endIndex  int
}

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
		return trucks[i].Start < trucks[j].Start || (trucks[i].Start == trucks[j].Start && trucks[i].ID < trucks[j].ID)
	})

	if debugEnable {
		log.Println("order:", arrivalIndex)
		log.Println("trucks:", trucks)
	}

	var (
		startQ StartQueue
		endQ   EndQueue
	)

	for _, ai := range arrivalIndex {
		t := arrival[ai]
		if debugEnable {
			log.Println("ii:", ai, "t:", t)
		}

		for len(trucks) > 0 && trucks[0].Start <= t {
			if debugEnable {
				log.Println("+", trucks[0].ID)
			}
			if trucks[0].End >= t {
				heap.Push(&startQ, &trucks[0])
				heap.Push(&endQ, &trucks[0])
			}
			trucks = trucks[1:]
		}

		for endQ.Len() > 0 && endQ[0].End < t {
			if debugEnable {
				log.Println("-", endQ[0].ID)
			}
			heap.Remove(&startQ, endQ[0].starIndex)
			heap.Pop(&endQ)
		}

		for startQ.Len() > 0 && startQ[0].Cap == 0 {
			if debugEnable {
				log.Println("*", startQ[0].ID)
			}
			heap.Remove(&endQ, startQ[0].endIndex)
			heap.Pop(&startQ)
		}

		// if startQ.Len() == 0 && len(trucks) > 0 {
		// 	if debugEnable {
		// 		log.Println("+", trucks[0].ID)
		// 	}
		// 	heap.Push(&startQ, &trucks[0])
		// 	heap.Push(&endQ, &trucks[0])
		// 	trucks = trucks[1:]
		// }

		if debugEnable {
			log.Println("startQ:", startQ)
			log.Println("endQ:", endQ)
		}

		if startQ.Len() == 0 {
			if debugEnable {
				log.Println(ai, "<-", -1)
			}
			ans[ai] = -1
			continue
		}

		if debugEnable {
			log.Println(ai, "<-", startQ[0].ID)
		}
		ans[ai] = startQ[0].ID
		startQ[0].Cap--
	}

	return ans
}

// A StartQueue implements heap.Interface and holds Items.
type StartQueue []*Truck

func (pq StartQueue) Len() int { return len(pq) }

func (pq StartQueue) Less(i, j int) bool {
	return pq[i].Start < pq[j].Start || (pq[i].Start == pq[j].Start && pq[i].ID < pq[j].ID)
}

func (pq StartQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].starIndex = i
	pq[j].starIndex = j
}

func (pq *StartQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Truck)
	item.starIndex = n
	*pq = append(*pq, item)
}

func (pq *StartQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil      // don't stop the GC from reclaiming the item eventually
	item.starIndex = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

type EndQueue []*Truck

func (pq EndQueue) Len() int { return len(pq) }

func (pq EndQueue) Less(i, j int) bool {
	return pq[i].End < pq[j].End
}

func (pq EndQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].endIndex = i
	pq[j].endIndex = j
}

func (pq *EndQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Truck)
	item.endIndex = n
	*pq = append(*pq, item)
}

func (pq *EndQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil     // don't stop the GC from reclaiming the item eventually
	item.endIndex = -1 // for safety
	*pq = old[0 : n-1]
	return item
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
			var v int
			if _, err := fmt.Fscan(br, &v); err != nil {
				panic(err)
			}
			arrival = append(arrival, v)
		}

		var m int
		if _, err := fmt.Fscan(br, &m); err != nil {
			panic(err)
		}

		if debugEnable {
			log.Println("m:", m)
		}

		trucks := make([]Truck, 0, m)
		for i := 0; i < m; i++ {
			var start, end, capacity int
			if _, err := fmt.Fscan(br, &start, &end, &capacity); err != nil {
				panic(err)
			}
			trucks = append(trucks, Truck{ID: i + 1, Start: start, End: end, Cap: capacity})
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
