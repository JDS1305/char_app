package services

import (
	"chat_app/internal/models"
	"chat_app/internal/repositories"
)

type FileService struct {
	FileRepo *repositories.FileRepository
}

func NewFileService(fileRepo *repositories.FileRepository) *FileService {
	return &FileService{FileRepo: fileRepo}
}

func (fs *FileService) UploadFile(file *models.File) error {
	err := fs.FileRepo.UploadFile(file)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileService) GetFileByID(fileID int) (*models.File, error) {
	file, err := fs.FileRepo.GetFileByID(fileID)
	if err != nil {
		return nil, err
	}

	return file, nil
}
