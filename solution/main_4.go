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
		var usersCount int
		fmt.Fscan(in, &usersCount)
		for j := 0; j < usersCount; j++ {
			var time int
			fmt.Fscan(in, &time)
			fmt.Println(time, "TIME")
		}
	}
}
