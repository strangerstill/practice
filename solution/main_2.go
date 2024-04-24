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

	var initStr string
	var repeatCount int

	fmt.Fscan(in, &initStr, &repeatCount)
	data := []byte(initStr)
	for i := 0; i < repeatCount; i++ {
		var (
			start, end int
			newStr     string
		)
		fmt.Fscan(in, &start, &end, &newStr)
		strIndex := start - 1
		if strIndex < end {
			copy(data[strIndex:end], []byte(newStr))
		}
	}
	fmt.Fprint(out, string(data))
}
