package controllers

import (
	"chat_app/internal/models"
	"chat_app/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type GroupController struct {
	GroupService *services.GroupService
}

func NewGroupController(groupService *services.GroupService) *GroupController {
	return &GroupController{GroupService: groupService}
}

func (gc *GroupController) CreateGroup(w http.ResponseWriter, r *http.Request) {
	var newGroup models.Group
	err := json.NewDecoder(r.Body).Decode(&newGroup)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = gc.GroupService.CreateGroup(&newGroup)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (gc *GroupController) GetGroupByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	groupID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}

	// Mendapatkan grup berdasarkan ID
	group, err := gc.GroupService.GetGroupByID(groupID)
	if err != nil {
		http.Error(w, "Group not found", http.StatusNotFound)
		return
	}

	// Mengembalikan data grup dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(group)
}
