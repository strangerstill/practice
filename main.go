package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Data struct {
	userTime   int
	userNumber int
}

type ResultData struct {
	win        int
	userNumber int
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var repeatCount int

	fmt.Fscan(in, &repeatCount)
	for i := 0; i < repeatCount; i++ {
		var usersCount int
		fmt.Fscan(in, &usersCount)
		var data []Data
		for j := 0; j < usersCount; j++ {
			var time int
			fmt.Fscan(in, &time)
			data = append(data, Data{userTime: time, userNumber: j})
		}
		sort.Slice(data, func(i, j int) bool {
			return data[i].userTime < data[j].userTime
		})
		sameWinCount := 0
		win := 1
		var result []ResultData
		for index := range data {
			if index == 0 {
				result = append(result, ResultData{
					win:        1,
					userNumber: data[index].userNumber,
				})
			} else if diff(data[index-1].userTime, data[index].userTime) <= 1 {
				if sameWinCount == 0 {
					sameWinCount = 1
				}
				sameWinCount += 1
				result = append(result, ResultData{
					win:        win,
					userNumber: data[index].userNumber,
				})
			} else {
				if sameWinCount > 0 {
					win += sameWinCount
					sameWinCount = 0
				} else {
					win += 1
				}
				result = append(result, ResultData{
					win:        win,
					userNumber: data[index].userNumber,
				})
			}
		}
		if len(result) == usersCount {
			var winResult []string
			sort.Slice(result, func(i, j int) bool {
				return result[i].userNumber < result[j].userNumber
			})
			for item := range result {
				winResult = append(winResult, strconv.Itoa(result[item].win))
			}
			fmt.Fprintln(out, strings.Join(winResult, " "))
		}
	}
}
