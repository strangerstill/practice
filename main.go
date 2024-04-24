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

func (r *IdReserver) sendNotification(userId int) {
	if userId == 0 {
		r.reserveGlobal()
	} else {
		r.reserveUser(userId)
	}
}

func (r *IdReserver) reserveUser(userId int) {
	r.localId += 1
	r.data[userId] = r.localId
}

func (r *IdReserver) reserveGlobal() {
	r.localId += 1
	r.globalId = r.localId
}

func (r *IdReserver) getNotificationId(userId int) int {
	value, ok := r.data[userId]
	if !ok {
		return r.globalId
	}
	if r.globalId > value {
		return r.globalId
	}
	return value
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
			fmt.Fprintln(out, notification.getNotificationId(userId))
		}
	}
}
