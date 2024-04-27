package repositories

import (
	"chat_app/internal/models"
	"database/sql"
)

type FileRepository struct {
	DB *sql.DB
}

func NewFileRepository(db *sql.DB) *FileRepository {
	return &FileRepository{DB: db}
}

func (fr *FileRepository) UploadFile(file *models.File) error {
	_, err := fr.DB.Exec("INSERT INTO files (message_id, url, type) VALUES ($1, $2, $3)",
		file.MessageID, file.URL, file.Type)
	if err != nil {
		return err
	}
	return nil
}

func (fr *FileRepository) GetFileByID(fileID int) (*models.File, error) {
	var file models.File
	row := fr.DB.QueryRow("SELECT id, message_id, url, type FROM files WHERE id = $1", fileID)
	if err := row.Scan(&file.ID, &file.MessageID, &file.URL, &file.Type); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &file, nil
}
