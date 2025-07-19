package main

import "container/list"

func main() {
	myList := list.New()
	myList.PushFront(1)
	myList.PushFront(2)
	myList.PushBack(1)
}
