package todolist

import (
	"fmt"
	"sync"
)

// StorageServicer Storage micro service interface
type StorageServicer interface {
	SaveTodo(content string, done bool) error
	DeleteTodo(id string) (bool, error)
	UpdateTodo(id, content string, done bool) (bool, error)
	ListTodo(list *[]*TodoItem) error
}

// MemStorageService storage service implement with mem
type MemStorageService struct {
	idc int
	db  sync.Map
}

// NewMemStorageService create an entity with default settings
func NewMemStorageService() *MemStorageService {
	return &MemStorageService{}
}

// SaveTodo save data in mem
func (mss *MemStorageService) SaveTodo(content string, done bool) error {
	mss.idc++
	id := fmt.Sprintf("%d", mss.idc)
	mss.db.Store(id, &TodoItem{
		id,
		content,
		done,
	})
	return nil
}

// DeleteTodo delete data in mem with the id
func (mss *MemStorageService) DeleteTodo(id string) (bool, error) {
	if _, ok := mss.db.Load(id); !ok {
		return false, nil
	}
	mss.db.Delete(id)
	return true, nil
}

// UpdateTodo update data in mem with the id
func (mss *MemStorageService) UpdateTodo(id, content string, done bool) (bool, error) {
	if _, ok := mss.db.Load(id); !ok {
		return false, nil
	}
	mss.db.Store(id, &TodoItem{
		id,
		content,
		done,
	})
	return true, nil
}

// ListTodo load all data in mem to list
func (mss *MemStorageService) ListTodo(list *[]*TodoItem) error {
	mss.db.Range(func(key, value interface{}) bool {
		*list = append(*list, value.(*TodoItem))
		return true
	})
	return nil
}
