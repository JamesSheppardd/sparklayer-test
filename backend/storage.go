package main

type TodoItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type StorageInterface interface {
	Get(key int) TodoItem
	Set(item TodoItem)
	Delete(key int)
}

type Storage struct {
	items []TodoItem
}

func SetupStorage() *Storage {
	return &Storage{items: make([]TodoItem, 0)}
}

func (s *Storage) Get(key int) TodoItem {
	item := s.items[key]
	return item
}

func (s *Storage) Set(item TodoItem) {
	s.items = append(s.items, item)
}

func (s *Storage) Delete(key int) {
	s.items = append(s.items[:key], s.items[key+1:]...)
}
