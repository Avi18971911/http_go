package main

import (
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
				p.processWinHandler(player, resp)
			case http.MethodGet:
				p.scoreHandler(player, resp)
			}
		}))
	r.Handle("/league", http.HandlerFunc(
		func(resp http.ResponseWriter, req *http.Request) {
			p.leagueHandler(resp)
		}))
	return r
}
