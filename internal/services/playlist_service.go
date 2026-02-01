package services

import "YeahMusic/internal/models"

type PlaylistService struct {
	store *Store
}

func NewPlaylistService(store *Store) *PlaylistService {
	return &PlaylistService{store: store}
}

func (p *PlaylistService) Create(userID int64, title string) *models.Playlist {
	return p.store.CreatePlaylist(userID, title)
}

func (p *PlaylistService) List(userID int64) []*models.Playlist {
	return p.store.ListPlaylists(userID)
}

func (p *PlaylistService) AddTrack(userID, playlistID, trackID int64) error {
	return p.store.AddTrackToPlaylist(userID, playlistID, trackID)
}

func (p *PlaylistService) RemoveTrack(userID, playlistID, trackID int64) error {
	return p.store.RemoveTrackFromPlaylist(userID, playlistID, trackID)
}
