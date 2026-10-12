package store

import "sync"

type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}

// Todo is one item on the list

//Store will keep todos in memory.
type Store struct{
	mu sync.Mutex
	todos map[int]Todo
	nextID int
}

func New() *Store{
	return &Store{todos: make(map[int]Todo), nextID:1}
}

