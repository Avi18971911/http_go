package main

import (
	"fmt"
	"net/http"
	"strings"
)

func createRouter(p *PlayerServer) http.Handler {
	r := http.NewServeMux()
	r.Handle("/players/", http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			player := strings.TrimPrefix(req.URL.Path, "/players/")
			switch req.Method {
			case http.MethodPost:
				score := p.ProcessWin(player)
				fmt.Fprint(resp, score)
			case http.MethodGet:
				score := p.GetScore(player)
				fmt.Fprint(resp, score)
			}
		}))
	return r
}
