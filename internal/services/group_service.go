package services

import (
	"chat_app/internal/models"
	"chat_app/internal/repositories"
)

type GroupService struct {
	GroupRepo *repositories.GroupRepository
}

func NewGroupService(groupRepo *repositories.GroupRepository) *GroupService {
	return &GroupService{GroupRepo: groupRepo}
}

func (gs *GroupService) CreateGroup(group *models.Group) error {
	// Membuat grup baru
	err := gs.GroupRepo.CreateGroup(group)
	if err != nil {
		return err
	}

	return nil
}

func (gs *GroupService) GetGroupByID(groupID int) (*models.Group, error) {
	// Dapatkan grup berdasarkan ID
	group, err := gs.GroupRepo.GetGroupByID(groupID)
	if err != nil {
		return nil, err
	}

	return group, nil
}
