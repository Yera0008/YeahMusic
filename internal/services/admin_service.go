package services

import "YeahMusic/internal/models"

type AdminService struct {
	store *Store
}

func NewAdminService(store *Store) *AdminService {
	return &AdminService{store: store}
}

func (a *AdminService) Ping() *models.Admin {
	return &models.Admin{ID: 1, Email: "admin@local", Role: "admin"}
}
