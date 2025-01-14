package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Folder struct {
	Dir     string   `json:"dir"`
	Files   []string `json:"files"`
	Folders []Folder `json:"folders"`
}

func (f Folder) IsHacked() bool {
	for _, file := range f.Files {
		if strings.HasSuffix(file, ".hack") {
			return true
		}
	}
	return false
}

func (f Folder) CountFiles() int {
	count := len(f.Files)
	for _, subFolder := range f.Folders {
		count += subFolder.CountFiles()
	}
	return count
}

func (f Folder) CountHackedFiles() int {
	if f.IsHacked() {
		return f.CountFiles()
	}

	count := 0
	for _, subFolder := range f.Folders {
		count += subFolder.CountHackedFiles()
	}
	return count
}

func solve(buf []byte) int {
	var tree Folder

	if err := json.Unmarshal(buf, &tree); err != nil {
		panic(err)
	}

	return tree.CountHackedFiles()
}

func run(in io.Reader, out io.Writer) {
	br := bufio.NewReader(in)
	bw := bufio.NewWriter(out)
	defer bw.Flush()
	var t int

	if _, err := fmt.Fscanln(br, &t); err != nil {
		panic(err)
	}

	var buf []byte

	for i := 1; i <= t; i++ {
		var n int
		if _, err := fmt.Fscanln(br, &n); err != nil {
			panic(err)
		}

		buf = buf[:0]
		for i := 0; i < n; {
			line, isPreffix, err := br.ReadLine()
			if err != nil && err != io.EOF {
				panic(err)
			}
			
			buf = append(buf, line...)
			if !isPreffix {
				buf = append(buf, '\n')
				i++
			}
		}

		ans := solve(buf)
		fmt.Fprintln(bw, ans)
	}
}

var debugEnable bool

func main() {
	_, debugEnable = os.LookupEnv("DEBUG")
	run(os.Stdin, os.Stdout)
}
