package services

import "github.com/BFFTeam/YeahMusic/internal/models"

type AdminService struct{}

func (s *AdminService) AddSong(song models.Song) error {
	return nil
}

func (s *AdminService) UpdateSong(song models.Song) error {
	return nil
}

func (s *AdminService) DeleteSong(songID int) error {
	return nil
}
