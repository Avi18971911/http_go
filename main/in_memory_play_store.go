package main

type InMemoryPlayerStore struct {
	store map[string]int
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

func (i *InMemoryPlayerStore) RecordWin(name string) int {
	i.store[name]++
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}