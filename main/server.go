package main

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string) int
}

type PlayerServer struct {
	store PlayerStore
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{store}
}

func (p *PlayerServer) ProcessWin(player string) int {
	return p.store.RecordWin(player)
}

func (p *PlayerServer) GetScore(player string) int {
	return p.store.GetPlayerScore(player)
}
