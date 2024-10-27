package server

/*
	JSON/HTTP Web Server - Producer/Consumer Model
*/

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Produce - Writing to Log
// Consume - Reading from Log

// Steps for our handler
// Umarshal req JSON to struct
// Run endpoint logic and get result
// Marshal result to JSON response

type httpServer struct {
	Log *Log
}

func newHTTPServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

// GET Request
type ProduceRequest struct {
	Record Record `json: "record"`
}

// GET Response
type ProduceResponse struct {
	Offset uint64 `json: "offset"`
}

// POST Request
type ConsumeRequest struct {
	Offset uint64 `json: "offset"`
}

// POST Response - newly created record
type ConsumeResponse struct {
	Record Record `json: "record"`
}

func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	var req ProduceRequest
	// Unmarshal req to struct
	err := json.NewDecoder(r.Body).Decode(&req)
	// If we have error w/ req
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	off, err := s.Log.Append(req.Record)
	// Check for append error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Marshal struct to json
	res := ProduceResponse{Offset: off}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Consume Handler
func (s *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
	var req ConsumeRequest
	// Unmarshal into struct
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Read record
	rec, err := s.Log.Read(req.Offset)
	if err == ErrOffsetNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	// Marshal struct to JSON
	res := ConsumeResponse{Record: rec}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

/*
Takes in address for server to run on, returns address to it
Gorilla/mux for RESTful routes to match requests to handlers
net/http.Server for easy listening and handling of requests
*/
func NewHttpServer(addr string) *http.Server {
	httpserver := newHTTPServer()
	r := mux.NewRouter()
	r.HandleFunc("/", httpserver.handleProduce).Methods("POST")
	r.HandleFunc("/", httpserver.handleConsume).Methods("GET")
	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}
