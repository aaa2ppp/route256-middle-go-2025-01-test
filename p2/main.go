package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"unsafe"
)

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
		if _, err := fmt.Fscanln(br, &n); err != nil {
			panic(err)
		}

		a := make([]int, 0, n)

		s, err := br.ReadString('\n')
		if err != nil {
			panic(err)
		}

		sc := bufio.NewScanner(strings.NewReader(s))
		sc.Split(bufio.ScanWords)
		for sc.Scan() {
			v, err := strconv.Atoi(unsafeString(sc.Bytes()))
			if err != nil {
				panic(err)
			}
			a = append(a, v)
		}
		if err := sc.Err(); err != nil {
			panic(err)
		}

		sort.Ints(a)

		var sb strings.Builder
		for _, v := range a {
			sb.WriteString(strconv.Itoa(v))
			sb.WriteByte(' ')
		}

		want := sb.String()
		if len(want) > 0 {
			want = want[:len(want)-1] // trim last space
		}

		got, err := br.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if len(got) > 0 {
			got = got[:len(got)-1] // trim last \n
		}

		if got == want {
			bw.WriteString("yes\n")
		} else {
			bw.WriteString("no\n")
		}
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}

func unsafeString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
