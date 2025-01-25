package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const dataDir = "../data"

// func Test_run(t *testing.T) {
// 	type args struct {
// 		in io.Reader
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		wantOut string
// 		debug   bool
// 	}{
// 		{
// 			"1",
// 			args{strings.NewReader(`1
// 3 3
// B..
// .#.
// ..A
// `)},
// 			`B..
// .#.
// ..A
// `,
// 			true,
// 		},
// 		{
// 			"2",
// 			args{strings.NewReader(`2
// 5 5
// .....
// .#A#.
// ...B.
// .#.#.
// .....
// 7 9
// .........
// .#.#.#.#.
// ..AB.....
// .#.#.#.#.
// .........
// .#.#.#.#.
// .........
// `)},
// 			`aaa..
// .#A#.
// ...Bb
// .#.#b
// ....b
// aaa......
// .#a#.#.#.
// ..ABb....
// .#.#b#.#.
// ....b....
// .#.#b#.#.
// ....bbbbb
// `,
// 			true,
// 		},
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			defer func(v bool) { debugEnable = v }(debugEnable)
// 			debugEnable = tt.debug

// 			out := &bytes.Buffer{}
// 			run(tt.args.in, out)
// 			if gotOut := out.String(); strings.TrimSuffix(gotOut, "\n") != strings.TrimSuffix(tt.wantOut, "\n") {
// 				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
// 			}
// 		})
// 	}
// }

func Test_run(t *testing.T) {
	type args struct {
		in io.Reader
	}
	tests := []struct {
		name    string
		args    args
		wantOut string
		debug   bool
	}{
		{
			"1",
			args{strings.NewReader(`1
3 3
B..
.#.
..A
`)},
			`B..
.#.
..A
`,
			true,
		},
		{
			"2",
			args{strings.NewReader(`2
5 5
.....
.#A#.
...B.
.#.#.
.....
7 9
.........
.#.#.#.#.
..AB.....
.#.#.#.#.
.........
.#.#.#.#.
.........
`)},
			`aaa..
.#A#.
...Bb
.#.#b
....b
a........
a#.#.#.#.
aaABbbbbb
.#.#.#.#b
........b
.#.#.#.#b
........b
`,
			true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(v bool) { debugEnable = v }(debugEnable)
			debugEnable = tt.debug

			out := &bytes.Buffer{}
			run(tt.args.in, out)
			if gotOut := out.String(); strings.TrimSuffix(gotOut, "\n") != strings.TrimSuffix(tt.wantOut, "\n") {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func getTestNames() []string {
	files, err := os.ReadDir(dataDir)
	if err != nil {
		panic(err)
	}

	names := make([]string, 0, len(files))
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := file.Name()
		if ok, err := filepath.Match("*.a", fileName); err != nil {
			panic(err)
		} else if !ok {
			continue
		}

		fileName = fileName[:len(fileName)-2]
		names = append(names, fileName)
	}

	return names
}

// func Test_run2(t *testing.T) {
// 	tests := getTestNames()

// 	for _, test := range tests {
// 		input, err := os.ReadFile(filepath.Join(dataDir, test))
// 		if err != nil {
// 			panic(err)
// 		}

// 		t.Run(test, func(t *testing.T) {
// 			wantOut := &strings.Builder{}
// 			run(bytes.NewReader(input), wantOut)
// 			want := wantOut.String()

// 			gotOut := &strings.Builder{}
// 			run2(bytes.NewReader(input), gotOut)
// 			got := gotOut.String()

// 			if got != want {
// 				t.Errorf("run() = %v, \nwant %v", got, want)
// 			}
// 		})
// 	}
// }

func Benchmark_run(b *testing.B) {
	benchs := []struct {
		name string
		run  func(io.Reader, io.Writer)
	}{
		{"copy", func(in io.Reader, out io.Writer) { io.Copy(out, in) }}, // погреваем кеш
		{"copy", func(in io.Reader, out io.Writer) { io.Copy(out, in) }}, // сравниваем
		{"run", run},
		// {"run2", run2},
	}

	tests := getTestNames()

	for _, bb := range benchs {
		b.Run(bb.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				for _, test := range tests {
					func() {
						in, err := os.Open(filepath.Join(dataDir, test))
						if err != nil {
							panic(err)
						}
						defer in.Close()
						bb.run(in, io.Discard)
					}()
				}
			}
		})
	}
}
