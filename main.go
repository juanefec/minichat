package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path == "/" {
		http.ServeFile(w, r, "home.html")
	}
	if r.URL.Path == "/sketch/" {
		http.ServeFile(w, r, "sketch/index.html")
	}
	if r.URL.Path == "/sketch/sketch.js" {
		http.ServeFile(w, r, "sketch/sketch.js")
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.Error(w, "Not found", http.StatusNotFound)
	return

}

func main() {
	flag.Parse()
	hub := newHub()
	go hub.run()
	http.HandleFunc("/", serveHome)
	http.Handle("/sketch", http.FileServer(http.Dir("./sketch")))
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
