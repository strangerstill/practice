package main

import (
	"bufio"
	"fmt"
	"os"
)

type IdReserver struct {
	globalId int
	localId  int
	data     map[int]int
}

func NewIdReserver() *IdReserver {
	return &IdReserver{
		data: map[int]int{},
	}
}

func (n *IdReserver) sendNotification(userId int) {
	n.localId += 1
	if userId == 0 {
		n.globalId = n.localId
	} else {
		n.data[userId] = n.localId
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var userCount, repeatCount int

	fmt.Fscan(in, &userCount, &repeatCount)

	notification := NewIdReserver()

	for i := 0; i < repeatCount; i++ {
		var (
			requestType, userId int
		)
		fmt.Fscan(in, &requestType, &userId)
		if requestType == 1 {
			notification.sendNotification(userId)
		} else if requestType == 2 {
			if value, ok := notification.data[userId]; ok {
				if notification.globalId > value {
					fmt.Fprintln(out, notification.globalId)
				} else {
					fmt.Fprintln(out, value)
				}
			} else {
				fmt.Fprintln(out, notification.globalId)
			}
		}
	}
}
