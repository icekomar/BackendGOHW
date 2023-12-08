package storage

import (
	"dz/pkg/models"
	"fmt"
	"strconv"
	"sync"
)

type LocalStorage struct {
	stor []models.Record
	mu   *sync.Mutex
}

func New() *LocalStorage {
	return &LocalStorage{stor: make([]models.Record, 0), mu: new(sync.Mutex)}
}

func (localstorage *LocalStorage) AddFromJson(m []models.Record) {
	localstorage.mu.Lock()
	for _, v := range m {
		localstorage.stor = append(localstorage.stor, v)
	}
	localstorage.mu.Unlock()
}

func (localstorage *LocalStorage) ToString(t bool) string {

	s := "List:\n"
	if len(localstorage.stor) == 0 {
		return s
	}
	for _, v := range localstorage.stor {
		if v.IsQuestion == t {
			s = fmt.Sprintf("%s %s : %s \n", s, strconv.Itoa(v.Num), v.Body)
		}
	}
	return s
}
