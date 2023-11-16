package Structs

import (
	"strconv"
	"sync"
)

type LocalStorage struct {
	storage []Record
	mu      *sync.Mutex
}

func New() *LocalStorage {
	return &LocalStorage{storage: make([]Record, 0), mu: new(sync.Mutex)}
}

func (localstorage *LocalStorage) AddFromJson(m []Record) {
	for _, v := range m {
		localstorage.storage = append(localstorage.storage, v)
	}
}

func (localstorage *LocalStorage) Lock() {
	localstorage.mu.Lock()
}

func (localstorage *LocalStorage) Unlock() {
	localstorage.mu.Unlock()
}
func (localstorage *LocalStorage) ToString(t bool) string {
	s := "List:\n"
	if len(localstorage.storage) == 0 {
		return s
	}
	for _, v := range localstorage.storage {
		if v.T == t {
			s += strconv.Itoa(v.Num) + ":" + v.Body + "\n"
		}
	}
	return s
}
