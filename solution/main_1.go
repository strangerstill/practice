package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var repeatCount int

	fmt.Fscan(in, &repeatCount)
	for i := 0; i < repeatCount; i++ {
		var (
			a, b int
		)
		fmt.Fscan(in, &a, &b)
		fmt.Fprintln(out, a+b)
	}
}
