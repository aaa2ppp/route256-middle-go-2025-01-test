package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func solve(line string) int {
	for i := 1; i < len(line); i++ {
		if line[i] > line[i-1] {
			return i - 1
		}
	}
	return len(line) - 1
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
		line, err := br.ReadString('\n')
		if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)
		if debugEnable {
			log.Printf("line:%q len:%d", line, len(line))
		}

		if len(line) < 2 {
			bw.WriteString("0\n")
			continue
		}

		idx := solve(line)
		head := line[:idx]
		tail := line[idx+1:]
		if debugEnable {
			log.Printf("idx:%d head:%q tail:%q", idx, head, tail)
		}
		bw.WriteString(head)
		bw.WriteString(tail)
		bw.WriteByte('\n')
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}
