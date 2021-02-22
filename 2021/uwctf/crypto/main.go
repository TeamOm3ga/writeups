package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type encryptingHandler struct {
	*http.ServeMux
	block cipher.Block
}

func newHandler() *encryptingHandler {
	key, _ := ioutil.ReadFile("key.txt")

	block, _ := aes.NewCipher(key)

	h := &encryptingHandler{
		ServeMux: http.NewServeMux(),
		block:    block,
	}

	h.HandleFunc("/", h.help)
	h.HandleFunc("/list", h.listFiles)
	h.HandleFunc("/get", h.getFile)

	return h
}

const helpMessage = `This API has two endpoints:
GET  /list                    - list files available for retrieval
POST /get  {filename: string} - retrieve an encrypted file
`

func (h *encryptingHandler) help(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(helpMessage))
}

type listResponse struct {
	FileNames []string `json:"filenames"`
}

func (h *encryptingHandler) listFiles(w http.ResponseWriter, r *http.Request) {
	entries, err := ioutil.ReadDir(".")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := listResponse{
		FileNames: make([]string, len(entries)),
	}

	for i, entry := range entries {
		response.FileNames[i] = entry.Name()
	}

	json.NewEncoder(w).Encode(response)
}

type fileRequest struct {
	FileName string `json:"filename"`
}

type fileResponse struct {
	Payload []byte `json:"payload"`
}

func (h *encryptingHandler) getFile(w http.ResponseWriter, r *http.Request) {
	var request fileRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bs, err := ioutil.ReadFile(request.FileName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	counter := make([]byte, h.block.BlockSize())
	response := fileResponse{
		Payload: make([]byte, len(bs)),
	}

	timestamp := time.Now().Unix()
	binary.BigEndian.PutUint64(counter, uint64(timestamp))

	stream := cipher.NewCTR(h.block, counter)
	stream.XORKeyStream(response.Payload, bs)

	json.NewEncoder(w).Encode(response)
}

func main() {
	handler := newHandler()
	server := &http.Server{
		Addr:         ":80",
		Handler:      handler,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	for {
		server.ListenAndServe()
	}
}
