package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

func run(in io.Reader, out io.Writer) {
	br := bufio.NewReader(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()
	var t int

	if _, err := fmt.Fscanln(br, &t); err != nil {
		panic(err)
	}

	var n, m int
	var matrix [][]byte

	for i := 1; i <= t; i++ {
		if _, err := fmt.Fscanln(br, &n, &m); err != nil {
			panic(err)
		}
		matrix = matrix[:0]
		for i := 0; i < n; i++ {
			line, err := br.ReadBytes('\n')
			if err != nil {
				panic(err)
			}
			matrix = append(matrix, line[:m])
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
