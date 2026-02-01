package services

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"YeahMusic/internal/models"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrForbidden     = errors.New("forbidden")
	ErrAlreadyExists = errors.New("already exists")
)

type Store struct {
	mu sync.RWMutex

	userID     int64
	sessionID  int64
	artistID   int64
	albumID    int64
	trackID    int64
	playlistID int64

	usersByID    map[int64]*models.User
	usersByEmail map[string]*models.User

	sessionsByToken map[string]*models.Session

	artistsByID map[int64]*models.Artist
	albumsByID  map[int64]*models.Album
	tracksByID  map[int64]*models.Track

	playlistsByID  map[int64]*models.Playlist
	playlistTracks map[int64]map[int64]*models.PlaylistTrack
}

func NewStore() *Store {
	return &Store{
		usersByID:       map[int64]*models.User{},
		usersByEmail:    map[string]*models.User{},
		sessionsByToken: map[string]*models.Session{},
		artistsByID:     map[int64]*models.Artist{},
		albumsByID:      map[int64]*models.Album{},
		tracksByID:      map[int64]*models.Track{},
		playlistsByID:   map[int64]*models.Playlist{},
		playlistTracks:  map[int64]map[int64]*models.PlaylistTrack{},
	}
}

func (s *Store) nextID(ptr *int64) int64 {
	return atomic.AddInt64(ptr, 1)
}

func (s *Store) Now() time.Time { return time.Now() }

// ----- Users -----

func (s *Store) CreateUser(u *models.User) (*models.User, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.usersByEmail[u.Email]; ok {
		return nil, ErrAlreadyExists
	}
	u.ID = s.nextID(&s.userID)
	u.CreatedAt = s.Now()
	s.usersByID[u.ID] = u
	s.usersByEmail[u.Email] = u
	return u, nil
}

func (s *Store) GetUserByEmail(email string) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.usersByEmail[email]
	if !ok {
		return nil, ErrNotFound
	}
	return u, nil
}

func (s *Store) GetUserByID(id int64) (*models.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	u, ok := s.usersByID[id]
	if !ok {
		return nil, ErrNotFound
	}
	return u, nil
}

// ----- Sessions -----

func (s *Store) CreateSession(userID int64, token string, ttl time.Duration) *models.Session {
	s.mu.Lock()
	defer s.mu.Unlock()

	sec := &models.Session{
		ID:        s.nextID(&s.sessionID),
		UserID:    userID,
		Token:     token,
		CreatedAt: s.Now(),
		ExpiresAt: s.Now().Add(ttl),
	}
	s.sessionsByToken[token] = sec
	return sec
}

func (s *Store) GetSession(token string) (*models.Session, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	sec, ok := s.sessionsByToken[token]
	if !ok {
		return nil, ErrNotFound
	}
	return sec, nil
}

func (s *Store) DeleteSession(token string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.sessionsByToken, token)
}

func (s *Store) CleanupExpiredSessions(now time.Time) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	n := 0
	for tok, sec := range s.sessionsByToken {
		if !sec.ExpiresAt.After(now) {
			delete(s.sessionsByToken, tok)
			n++
		}
	}
	return n
}

// ----- Catalog -----

func (s *Store) AddArtist(a *models.Artist) *models.Artist {
	s.mu.Lock()
	defer s.mu.Unlock()
	a.ID = s.nextID(&s.artistID)
	s.artistsByID[a.ID] = a
	return a
}

func (s *Store) AddAlbum(a *models.Album) *models.Album {
	s.mu.Lock()
	defer s.mu.Unlock()
	a.ID = s.nextID(&s.albumID)
	s.albumsByID[a.ID] = a
	return a
}

func (s *Store) AddTrack(t *models.Track) *models.Track {
	s.mu.Lock()
	defer s.mu.Unlock()
	t.ID = s.nextID(&s.trackID)
	s.tracksByID[t.ID] = t
	return t
}

func (s *Store) ListArtists() []*models.Artist {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]*models.Artist, 0, len(s.artistsByID))
	for _, a := range s.artistsByID {
		out = append(out, a)
	}
	return out
}

func (s *Store) ListAlbums(artistID int64) []*models.Album {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := []*models.Album{}
	for _, a := range s.albumsByID {
		if artistID == 0 || a.ArtistID == artistID {
			out = append(out, a)
		}
	}
	return out
}

func (s *Store) ListTracks(albumID int64) []*models.Track {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := []*models.Track{}
	for _, t := range s.tracksByID {
		if albumID == 0 || t.AlbumID == albumID {
			out = append(out, t)
		}
	}
	return out
}

// ----- Playlists -----

func (s *Store) CreatePlaylist(userID int64, title string) *models.Playlist {
	s.mu.Lock()
	defer s.mu.Unlock()

	p := &models.Playlist{
		ID:        s.nextID(&s.playlistID),
		UserID:    userID,
		Title:     title,
		CreatedAt: s.Now(),
	}
	s.playlistsByID[p.ID] = p
	if _, ok := s.playlistTracks[p.ID]; !ok {
		s.playlistTracks[p.ID] = map[int64]*models.PlaylistTrack{}
	}
	return p
}

func (s *Store) ListPlaylists(userID int64) []*models.Playlist {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := []*models.Playlist{}
	for _, p := range s.playlistsByID {
		if p.UserID == userID {
			out = append(out, p)
		}
	}
	return out
}

func (s *Store) AddTrackToPlaylist(userID, playlistID, trackID int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.playlistsByID[playlistID]
	if !ok {
		return ErrNotFound
	}
	if p.UserID != userID {
		return ErrForbidden
	}
	if _, ok := s.tracksByID[trackID]; !ok {
		return ErrNotFound
	}
	if _, ok := s.playlistTracks[playlistID]; !ok {
		s.playlistTracks[playlistID] = map[int64]*models.PlaylistTrack{}
	}
	s.playlistTracks[playlistID][trackID] = &models.PlaylistTrack{
		PlaylistID: playlistID,
		TrackID:    trackID,
		AddedAt:    s.Now(),
	}
	return nil
}

func (s *Store) RemoveTrackFromPlaylist(userID, playlistID, trackID int64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	p, ok := s.playlistsByID[playlistID]
	if !ok {
		return ErrNotFound
	}
	if p.UserID != userID {
		return ErrForbidden
	}
	if _, ok := s.playlistTracks[playlistID]; !ok {
		return ErrNotFound
	}
	if _, ok := s.playlistTracks[playlistID][trackID]; !ok {
		return ErrNotFound
	}
	delete(s.playlistTracks[playlistID], trackID)
	return nil
}
