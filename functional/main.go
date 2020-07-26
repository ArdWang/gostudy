package main

import (
	"bufio"
	"fmt"
	"github.com/gostudy/functional/fib"
	"io"
	"strings"
)

// 1, 1, 2, 3, 5, 8, 13, ...
// a, b

type intGen func() int

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000{
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader)  {
	scanncer := bufio.NewScanner(reader)
	for scanncer.Scan() {
		fmt.Println(scanncer.Text())
	}
}

func main() {
	var f intGen = fib.Fibonacci()

	printFileContents(f)
}
