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
			args{strings.NewReader(`3
5
1 9 2 6 4
3
1 8 3
3 10 1
4 7 4
5
1 9 2 6 4
3
1 8 3
3 10 2
4 7 4
8
100 37 19 2 46 4 15 88
4
27 80 1
1 46 2
41 83 1
1 75 2
`)},
			`1 -1 1 2 1 
1 2 1 2 1 
-1 1 4 2 3 2 4 -1 
`,
			true,
		},
		{
			"2",
			args{strings.NewReader(`5
1
1
1
1 1000000000 1
1
1
1
2 1000000000 1
1
1000000000
1
1 999999999 1
1
500000000
1
1 1000000000 1
1
1000000000
1
1 1000000000 1
`)},
			`1 
-1 
-1 
1 
1 
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
			if gotOut := out.String(); strings.TrimSuffix(gotOut, "\n") != strings.TrimSuffix(linesTrimSpace(tt.wantOut), "\n") {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func linesTrimSpace(s string) string {
	lines := strings.Split(s, "\n")
	for i := range lines {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return strings.Join(lines, "\n")
}

func Test_run2(t *testing.T) {
	type args struct {
		in io.Reader
	}
	type test struct {
		name    string
		args    args
		wantOut string
	}

	files, err := os.ReadDir(dataDir)
	if err != nil {
		panic(err)
	}

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

		wantOut, err := os.ReadFile(filepath.Join(dataDir, fileName))
		if err != nil {
			panic(err)
		}

		in, err := os.Open(filepath.Join(dataDir, fileName[:len(fileName)-2]))
		if err != nil {
			panic(err)
		}

		tt := test{
			fileName,
			args{in},
			string(wantOut),
		}

		t.Run(tt.name, func(t *testing.T) {
			out := &bytes.Buffer{}
			run(tt.args.in, out)
			if gotOut := out.String(); strings.TrimSuffix(gotOut, "\n") != strings.TrimSuffix(linesTrimSpace(tt.wantOut), "\n") {
				t.Errorf("run() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Benchmark_run(b *testing.B) {
	buf, err := os.ReadFile(dataDir + "/18")
	if err != nil {
		panic(err)
	}
	for i := 0; i < b.N; i++ {
		run(bytes.NewReader(buf), io.Discard)
	}
}
