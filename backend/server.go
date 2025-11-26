package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct {
	store   Storage
	currKey int
}

func SetupServer(store Storage) *Server {
	return &Server{store: store, currKey: 0}
}

func (server *Server) CreateToDo(w http.ResponseWriter, r *http.Request) {
	var reqData TodoItem
	err := json.NewDecoder(r.Body).Decode(&reqData)
	if err != nil {
		http.Error(w, "Error while adding new todo item", http.StatusBadRequest)
		return
	}

	server.store.Set(reqData)
	json.NewEncoder(w).Encode(reqData)
	fmt.Printf("Succesfully added new todo item: %v, %v \n", reqData.Title, reqData.Description)
}

func (server *Server) GetToDos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(server.store.items)
}

type DeleteKey struct {
	Key int `json:"key"`
}

func (server *Server) DeleteTodoItem(w http.ResponseWriter, r *http.Request) {
	var key DeleteKey
	err := json.NewDecoder(r.Body).Decode(&key)
	if err != nil {
		http.Error(w, "Error while deleting todo item", http.StatusBadRequest)
		return
	}
	//fmt.Println(server.store.items[key.Key])
	server.store.Delete(key.Key)
	json.NewEncoder(w).Encode(server.store.items)
}

func (server *Server) HandleTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

	if r.Method == http.MethodOptions {
		return
	}

	switch r.Method {
	case http.MethodPost:
		server.CreateToDo(w, r)
	case http.MethodGet:
		server.GetToDos(w, r)
	case http.MethodDelete:
		server.DeleteTodoItem(w, r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
