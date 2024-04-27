package repositories

import (
	"chat_app/internal/models"
	"database/sql"
)

type GroupRepository struct {
	DB *sql.DB
}

func NewGroupRepository(db *sql.DB) *GroupRepository {
	return &GroupRepository{DB: db}
}

func (gr *GroupRepository) CreateGroup(group *models.Group) error {
	_, err := gr.DB.Exec("INSERT INTO groups (name) VALUES ($1)", group.Name)
	if err != nil {
		return err
	}
	return nil
}

func (gr *GroupRepository) GetGroupByID(groupID int) (*models.Group, error) {
	var group models.Group
	row := gr.DB.QueryRow("SELECT id, name FROM groups WHERE id = $1", groupID)
	if err := row.Scan(&group.ID, &group.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return &group, nil
}
