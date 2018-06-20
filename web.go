//go:generate browserify web/index.js web/js/ws.js -o static/bundle.js
//go:generate cp web/node_modules/three/build/three.min.js static/three.min.js
//go:generate cp web/js/controls/TrackballControls.js static/
//go:generate cp web/js/controls/FlyControls.js static/
//go:generate go-bindata-assetfs static/...
package main

import (
	"log"
	"net/http"
)

func startWeb(ws *WSServer, bind string) {
	fs := http.FileServer(assetFS())
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
