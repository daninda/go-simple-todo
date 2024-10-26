package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"simple-todo/internal/entity"
	"simple-todo/internal/usecase/interface/service"
	"strconv"

	"github.com/gorilla/mux"
)

type TodoHandler struct {
	service service.TodoService
}

func NewTodoHandler(service service.TodoService) *TodoHandler {
	return &TodoHandler{
		service: service,
	}
}

func (h *TodoHandler) FindAll(res http.ResponseWriter, req *http.Request) {
	todos, err := h.service.FindAll()
	if err != nil {
		log.Printf("Could not get todos: %v", err)
		http.Error(res, "Could not get todos", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(todos)
}

func (h *TodoHandler) FindOne(res http.ResponseWriter, req *http.Request) {
	idParam := mux.Vars(req)["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Could not get id: %v", err)
		http.Error(res, "", http.StatusBadRequest)
		return
	}

	todo, err := h.service.FindOne(id)
	if err != nil {
		log.Printf("Could not get todo: %v", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(todo)
}

func (h *TodoHandler) Create(res http.ResponseWriter, req *http.Request) {
	var todoBody entity.Todo

	err := json.NewDecoder(req.Body).Decode(&todoBody)
	if err != nil {
		log.Printf("Could not decode body: %v", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.service.Create(&todoBody)
	if err != nil {
		log.Printf("Could not create todo: %v", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	json.NewEncoder(res).Encode(todo)
}

func (h *TodoHandler) Update(res http.ResponseWriter, req *http.Request) {
	idParam := mux.Vars(req)["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Could not get id: %v", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	var todoBody entity.Todo

	err = json.NewDecoder(req.Body).Decode(&todoBody)
	if err != nil {
		log.Printf("Could not decode body: %v", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := h.service.Update(id, &todoBody)
	if err != nil {
		log.Printf("Could not update todo: %v", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(todo)
}

func (h *TodoHandler) Delete(res http.ResponseWriter, req *http.Request) {
	idParam := mux.Vars(req)["id"]

	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Printf("Could not get id: %v", err)
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.Delete(id)
	if err != nil {
		log.Printf("Could not delete todo: %v", err)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusNoContent)
}
