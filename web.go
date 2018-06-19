//go:generate browserify web/index.js web/js/ws.js -o web/bundle.js
package main

import (
	"log"
	"net/http"
)

func startWeb(ws *WSServer, bind string) {
	fs := http.FileServer(http.Dir("web"))
	http.Handle("/", noCacheMiddleware(fs))
	http.HandleFunc("/ws", ws.Handle)
	log.Fatal(http.ListenAndServe(bind, nil))
}

func noCacheMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "max-age=0, no-cache")
		h.ServeHTTP(w, r)
	})
}
