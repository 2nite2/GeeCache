package lru

import (
	"fmt"
	"testing"
)

type MyString string

func (m MyString) Len() int {
	return len(m)
}

func TestGet(t *testing.T) {
	cache := New(1000,
		func(key string, value Value) {
			fmt.Printf("removed key : %s, value: %v", key, value)
		})
	cache.Add("monkey", MyString("1234"))
	ok, value := cache.Get("monkey")
	if ok {
		fmt.Println(":111", value)
	}
}

func TestRemoveOldest(t *testing.T) {
	v1, v2, v3 := "111", "222", "333"
	k1, k2, k3 := "v1", "v2", "v3"
	cache := New(int64(len(v1+v2+k1+k2)), nil)
	cache.Add(k1, MyString(v1))
	cache.Add(k2, MyString(v2))
	cache.Add(k3, MyString(v3))

	if ok, _ := cache.Get(v1); !ok {
		fmt.Println("remove v1 success")
	}
}

func TestOnEvicted(t *testing.T) {
	cache := New(int64(10), func(key string, value Value) {
		fmt.Printf("removed key: %s,  value: %v", key, value)
	})
	cache.Add("11", MyString("111"))
	cache.Add("22", MyString("222"))
	cache.Add("33", MyString("333"))
	cache.Add("44", MyString("444"))

}
