package main

import (
	"bytes"
	"crypto/md5"
	"encoding/binary"
	"encoding/json"
	"hash"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var flag string

func init() {
	bs, _ := ioutil.ReadFile("flag.txt")
	flag = strings.TrimSpace(string(bs))
}

type teamHandler struct {
	*http.ServeMux
	hash   hash.Hash
	buffer []byte

	username []byte
	password []byte
	digest   []byte
}

func newTeamHandler(secret []byte) *teamHandler {
	buffer := make([]byte, 512)
	h := &teamHandler{
		ServeMux: http.NewServeMux(),
		hash:     md5.New(),

		username: buffer[0:256],
		password: buffer[256:384],
		digest:   buffer[384:512],
	}
	h.hash.Write(secret)
	h.hash.Sum(h.digest[0:0])

	h.HandleFunc("/username", h.readUsername)
	h.HandleFunc("/password", h.readPassword)
	h.HandleFunc("/flag", h.writeFlag)

	return h
}

func writeError(w http.ResponseWriter, message string, code int) {
	var payload struct {
		Error string `json:"error"`
	}
	payload.Error = message
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(&payload)
}

func writeMessage(w http.ResponseWriter, message string) {
	var payload struct {
		Message string `json:"message"`
	}
	payload.Message = message
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&payload)
}

func (s *teamHandler) readBuffer(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		writeError(w, "method must be POST", http.StatusMethodNotAllowed)
		return false
	}
	if r.Header.Get("Content-Type") != "application/octet-stream" {
		writeError(w, "content-type must be octet-stream", http.StatusUnsupportedMediaType)
		return false
	}
	r.Body.Read(s.buffer)
	return true
}

func (s *teamHandler) readUsername(w http.ResponseWriter, r *http.Request) {
	s.buffer = s.username
	if s.readBuffer(w, r) {
		writeMessage(w, "username stored")
	}
}

func (s *teamHandler) readPassword(w http.ResponseWriter, r *http.Request) {
	s.buffer = s.password
	if s.readBuffer(w, r) {
		writeMessage(w, "password stored")
	}
}

func (s *teamHandler) writeFlag(w http.ResponseWriter, r *http.Request) {
	// TODO: add username validation?
	s.hash.Reset()
	s.hash.Write(s.password)
	s.hash.Sum(s.password[0:0])
	if bytes.Equal(s.digest, s.password) {
		writeMessage(w, flag)
	} else {
		writeError(w, "password did not match secret", http.StatusForbidden)
	}
}

type toplevelHandler struct {
	teams map[string]*teamHandler
}

func (h *toplevelHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	token, _, _ := net.SplitHostPort(r.RemoteAddr)

	handler, ok := h.teams[token]
	if !ok {
		secret := make([]byte, 8)
		binary.BigEndian.PutUint64(secret, uint64(time.Now().UnixNano()))
		h.teams[token] = newTeamHandler(secret)
		handler = h.teams[token]
	}

	handler.ServeHTTP(w, r)
}

func newToplevelHandler() http.Handler {
	return &toplevelHandler{
		teams: make(map[string]*teamHandler),
	}
}

func main() {
	handler := newToplevelHandler()
	server := &http.Server{
		Addr:         ":2021",
		Handler:      handler,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	for {
		server.ListenAndServe()
	}
}
