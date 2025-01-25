package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
	"unsafe"
)

type point struct {
	i, j int
}

type robot struct {
	c byte
	point
}

func solve(matrix [][]byte) {
	n, m := len(matrix), len(matrix[0])

	var a, b robot
searchRobotLoop:
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch matrix[i][j] {
			case 'A':
				a = robot{'a', point{i, j}}
				if b.c != 0 {
					break searchRobotLoop
				}
			case 'B':
				b = robot{'b', point{i, j}}
				if a.c != 0 {
					break searchRobotLoop
				}
			}
			if i%2 == 1 {
				j++
			}
		}
	}

	if a.j == b.j {
		if a.i > b.i {
			a, b = b, a
		}
	} else {
		if a.j > b.j {
			a, b = b, a
		}
	}

	if a.i%2 == 1 {
		a.i--
		matrix[a.i][a.j] = a.c
	}
	for a.j > 0 {
		a.j--
		matrix[a.i][a.j] = a.c
	}
	for a.i > 0 {
		a.i--
		matrix[a.i][a.j] = a.c
	}

	if b.i%2 == 1 {
		b.i++
		matrix[b.i][b.j] = b.c
	}
	for b.j < m-1 {
		b.j++
		matrix[b.i][b.j] = b.c
	}
	for b.i < n-1 {
		b.i++
		matrix[b.i][b.j] = b.c
	}
}

func unsafeString(b []byte) string { return *(*string)(unsafe.Pointer(&b)) }

func run(in io.Reader, out io.Writer) {
	br := bufio.NewReader(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()

	buf := make([]byte, 100*100)
	matrix := make([][]byte, 100)

	line, _, err := br.ReadLine()
	if err != nil {
		panic(err)
	}

	t, err := strconv.Atoi(unsafeString(line))
	if err != nil {
		panic(err)
	}

	for i := 1; i <= t; i++ {
		line, _, err := br.ReadLine()
		if err != nil {
			panic(err)
		}

		w := bytes.Split(line, []byte(" "))
		n, err := strconv.Atoi(unsafeString(w[0]))
		if err != nil {
			panic(err)
		}

		m, err := strconv.Atoi(unsafeString(w[1]))
		if err != nil {
			panic(err)
		}

		matrix = matrix[:n]
		for i, j := 0, 0; i < n; i, j = i+1, j+m {
			line, _, err := br.ReadLine()
			if err != nil {
				panic(err)
			}

			matrix[i] = buf[j : j+m]
			copy(matrix[i], line)
		}

		solve(matrix)
		for _, row := range matrix {
			bw.Write(row)
			bw.WriteByte('\n')
		}
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}
