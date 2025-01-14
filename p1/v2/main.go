package main

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
	"unsafe"
)

func run(in io.Reader, out io.Writer) {
	br := bufio.NewReader(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()
	var t int

	line, isPrefix, err := br.ReadLine()
	if err != nil {
		panic(err)
	}
	if isPrefix {
		panic("line too long")
	}

	t, err = strconv.Atoi(strings.TrimSpace(unsafeString(line)))
	if err != nil {
		panic(err)
	}

taskLoop:
	for i := 1; i <= t; i++ {
		count := 0
		prev, err := br.ReadByte()
		if err != nil {
			panic(err)
		}

		for {
			cur, err := br.ReadByte()
			if err != nil {
				panic(err)
			}

			if cur == '\n' {
				if count == 0 {
					bw.WriteString("0\n")
				} else {
					bw.WriteByte('\n')
				}
				continue taskLoop
			}

			if cur > prev {
				bw.WriteByte(cur)
				break
			}

			bw.WriteByte(prev)
			prev = cur
			count++
		}

		for {
			cur, err := br.ReadByte()
			if err != nil {
				panic(err)
			}
			bw.WriteByte(cur)
			if cur == '\n' {
				break
			}
		}
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}

// ---------------------------------

func unsafeString(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}
