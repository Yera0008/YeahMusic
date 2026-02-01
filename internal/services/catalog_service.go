package services

import "YeahMusic/internal/models"

type CatalogService struct {
	store *Store
}

func NewCatalogService(store *Store) *CatalogService {
	return &CatalogService{store: store}
}

func (c *CatalogService) ListArtists() []*models.Artist {
	return c.store.ListArtists()
}

func (c *CatalogService) ListAlbums(artistID int64) []*models.Album {
	return c.store.ListAlbums(artistID)
}

func (c *CatalogService) ListTracks(albumID int64) []*models.Track {
	return c.store.ListTracks(albumID)
}
