package lru

import "container/list"

type Cache struct {
	maxBytes int64
	nBytes   int64
	ll       *list.List
	cache    map[string]*list.Element
	Evicted  func(key string, value Value)
}

type entry struct {
	key   string
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, evicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes: maxBytes,
		nBytes:   0,
		ll:       list.New(),
		cache:    make(map[string]*list.Element),
		Evicted:  evicted,
	}
}
func (c *Cache) Get(key string) (ok bool, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.PushBack(ele)
		kv := ele.Value.(*entry)
		return true, kv.value
	}
	return
}

func (c *Cache) RemoveOldest() {
	ele := c.ll.Front()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nBytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.Evicted != nil {
			c.Evicted(kv.key, kv.value)
		}
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToBack(ele)
		kv := ele.Value.(*entry)
		c.nBytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushBack(&entry{key, value})
		c.cache[key] = ele
		kv := ele.Value.(*entry)
		c.nBytes += int64(len(key)) + int64(kv.value.Len())
	}

	for c.maxBytes != 0 && c.nBytes > c.maxBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}
