package controllers

import (
	"chat_app/internal/models"
	"chat_app/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type MessageController struct {
	MessageService *services.MessageService
}

func NewMessageController(messageService *services.MessageService) *MessageController {
	return &MessageController{MessageService: messageService}
}

func (mc *MessageController) CreateMessage(w http.ResponseWriter, r *http.Request) {
	var newMessage models.Message
	err := json.NewDecoder(r.Body).Decode(&newMessage)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = mc.MessageService.CreateMessage(&newMessage)
	if err != nil {
		http.Error(w, "Failed to create message", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (mc *MessageController) GetMessageByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	messageID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid message ID", http.StatusBadRequest)
		return
	}

	message, err := mc.MessageService.GetMessageByID(messageID)
	if err != nil {
		http.Error(w, "Message not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
