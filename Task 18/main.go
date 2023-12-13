package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type MutexCounter struct {
	num int
	m sync.RWMutex
}

func (mc *MutexCounter) Add() {
	mc.m.Lock()
	mc.num++
	mc.m.Unlock()
}

func (mc *MutexCounter) Get() int {
	mc.m.RLock()
	defer mc.m.RUnlock()
	return mc.num
}

type AtomicCounter struct {
	num int64
}

func (ac *AtomicCounter) Add() {
	atomic.AddInt64(&ac.num, 1)
}

func (ac *AtomicCounter) Get() int {
	return int(atomic.LoadInt64(&ac.num))
}

func main() {
	mc := MutexCounter{}
	ac := AtomicCounter{}

	wg := sync.WaitGroup{}
	wg.Add(500)

	mutexStartTime := time.Now()
	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			mc.Add()
		}()
	}
	wg.Wait()
	mutexEndTime := time.Since(mutexStartTime)

	wg.Add(500)
	atomicStartTime := time.Now()
	for i := 0; i < 500; i++ {
		go func() {
			defer wg.Done()
			ac.Add()
		}()
	}
	wg.Wait()
	atomicEndTime := time.Since(atomicStartTime)

	fmt.Printf("AtomicCounter time: %v, counter: %v\n", atomicEndTime, ac.Get())
	fmt.Printf("MutexCounter time: %v, counter: %v", mutexEndTime, mc.Get())
}