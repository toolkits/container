package maps

import (
	"sync"
)

type StringIntMap struct {
	sync.RWMutex
	M map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{M: make(map[string]int)}
}

func (this *StringIntMap) Put(key string, val int) {
	this.Lock()
	defer this.Unlock()
	this.M[key] = val
}

func (this *StringIntMap) Get(key string) (int, bool) {
	this.RLock()
	defer this.RUnlock()
	val, exists := this.M[key]
	return val, exists
}

func (this *StringIntMap) Exists(key string) bool {
	_, exists := this.Get(key)
	return exists
}

func (this *StringIntMap) Remove(key string) {
	this.Lock()
	defer this.Unlock()
	delete(this.M, key)
}
