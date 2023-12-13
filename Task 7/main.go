package main

import (
	"fmt"
	"sync"
)

type cMap struct {
	sync.RWMutex
	d map[int]int
}

func (m *cMap) get(ind int) int {
	m.RLock()
	defer m.RUnlock()
	return m.d[ind]
}

func (m *cMap) set(ind, v int) {
	m.Lock()
	defer m.Unlock()
	m.d[ind] = v
}

func main() {
	m := cMap{
		d: make(map[int]int),
	}

	m.set(1, 100)
	m.set(2, 200)
	m.set(5, 300)

	fmt.Println(m.d)
}