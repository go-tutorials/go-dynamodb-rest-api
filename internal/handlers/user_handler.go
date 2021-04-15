package handlers

import (
	"encoding/json"
	server "github.com/common-go/http"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"

	. "go-service/internal/models"
	. "go-service/internal/services"
)

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	result, err := h.service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func (h *UserHandler) Load(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	result, err := h.service.Load(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func (h *UserHandler) Insert(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}

	result, er2 := h.service.Insert(&user)
	if er2 != nil {
		http.Error(w, er1.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func (h *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var user User
	er1 := json.NewDecoder(r.Body).Decode(&user)
	defer r.Body.Close()
	if er1 != nil {
		http.Error(w, er1.Error(), http.StatusBadRequest)
		return
	}
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	if len(user.Id) == 0 {
		user.Id = id
	} else if id != user.Id {
		http.Error(w, "Id not match", http.StatusBadRequest)
		return
	}

	result, er2 := h.service.Update(&user)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func (h *UserHandler) Patch(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}

	ids := []string{"id"}

	userType := reflect.TypeOf(User{})
	_, jsonMap := server.BuildMapField(userType)
	body, _, er1 := server.BodyToJson(r, userType, ids, jsonMap, nil)
	if er1 != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	result, er2 := h.service.Patch(r.Context(),body)
	if er2 != nil {
		http.Error(w, er2.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}
func (h *UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if len(id) == 0 {
		http.Error(w, "Id cannot be empty", http.StatusBadRequest)
		return
	}
	result, err := h.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	respond(w, result)
}

func respond(w http.ResponseWriter, result interface{}) {
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
