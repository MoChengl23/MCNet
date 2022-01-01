package main

import (
	"container/list"
	"fmt"
)

type test struct {
	q list.List
}

func (test *test) AddQueue() {
	test.q.PushBack("1iljhgbuil2")
	for i := test.q.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}

func (test *test) GetLen() int {

	return test.q.Len()
}
