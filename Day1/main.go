package main

import (
	"container/list"
	"fmt"
)

func main() {
	myList := list.New()
	myList.PushFront(1)
	myList.PushFront(2)
	myList.PushBack(1)
	fmt.Println("loha")
}
