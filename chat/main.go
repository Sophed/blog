package main

import (
	"fmt"
	"net/http"

	"github.com/olahol/melody"
	"github.com/sophed/lg"
)

const PORT = "1337"

var m *melody.Melody

func main() {
	m = melody.New()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.HandleRequest(w, r)
	})
	m.HandleConnect(connect)
	m.HandleDisconnect(disconnect)
	m.HandleMessage(message)
	lg.Dbug("listening at ws://127.0.0.1:" + PORT)
	http.ListenAndServe(":"+PORT, nil)
}

func connect(s *melody.Session) {
	username := randomAdjective() + "-" + randomAnimal()
	s.Set("username", username)
	log := fmt.Sprintf(
		"[+] %s connected", username,
	)
	m.Broadcast([]byte(log))
	lg.Dbug(log)
}

func disconnect(s *melody.Session) {
	username, ok := s.Get("username")
	if !ok {
		return
	}
	log := fmt.Sprintf(
		"[-] %s disconnected", username,
	)
	m.Broadcast([]byte(log))
	lg.Dbug(log)
}

func message(s *melody.Session, msg []byte) {
	username, ok := s.Get("username")
	if !ok {
		return
	}
	log := fmt.Sprintf(
		"<%s> %s", username, msg,
	)
	m.Broadcast([]byte(log))
	lg.Info(log)
}
