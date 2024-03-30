package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const jsonContentType = "application/json"

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string) int
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
}

type Player struct {
	Name string
	Score int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	return &PlayerServer{store}
}

func (p *PlayerServer) processWinHandler(player string, resp http.ResponseWriter) {
	numWins := p.store.RecordWin(player)
	resp.WriteHeader(http.StatusAccepted)
	fmt.Fprint(resp, numWins)
}

func (p *PlayerServer) scoreHandler(player string, resp http.ResponseWriter) {
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		resp.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(resp, score)
}

func (p *PlayerServer) leagueHandler(resp http.ResponseWriter) {
	resp.Header().Set("content-type", jsonContentType)
	json.NewEncoder(resp).Encode(p.store.GetLeague())
}
