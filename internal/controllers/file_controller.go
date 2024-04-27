package controllers

import (
	"chat_app/internal/models"
	"chat_app/internal/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type FileController struct {
	FileService *services.FileService
}

func NewFileController(fileService *services.FileService) *FileController {
	return &FileController{FileService: fileService}
}

func (fc *FileController) UploadFile(w http.ResponseWriter, r *http.Request) {
	var newFile models.File
	err := json.NewDecoder(r.Body).Decode(&newFile)
	if err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err = fc.FileService.UploadFile(&newFile)
	if err != nil {
		http.Error(w, "Failed to upload file", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (fc *FileController) GetFileByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fileID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid file ID", http.StatusBadRequest)
		return
	}

	file, err := fc.FileService.GetFileByID(fileID)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	// Mengembalikan data berkas dalam format JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(file)
}
