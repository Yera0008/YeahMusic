package services

import (
	"YeahMusic/internal/models"
	"time"
)

func Seed(store *Store) {
	a1 := store.AddArtist(&models.Artist{Name: "OG Buda", Bio: "Demo artist"})
	a2 := store.AddArtist(&models.Artist{Name: "Ernar Beats", Bio: "Demo artist"})

	al1 := store.AddAlbum(&models.Album{ArtistID: a1.ID, Title: "Скучаю но Работаю", CoverURL: "", ReleaseYear: 2023})
	al2 := store.AddAlbum(&models.Album{ArtistID: a2.ID, Title: "Скучаю но Ещё Работаю", CoverURL: "", ReleaseYear: 2025})

	store.AddTrack(&models.Track{AlbumID: al1.ID, Title: "Ссора Я", DurationSec: 180, AudioURL: "/audio/1.mp3", Lyrics: "la la"})
	store.AddTrack(&models.Track{AlbumID: al1.ID, Title: "Вода", DurationSec: 210, AudioURL: "/audio/2.mp3", Lyrics: "na na"})
	store.AddTrack(&models.Track{AlbumID: al2.ID, Title: "Всё Норм", DurationSec: 200, AudioURL: "/audio/3.mp3", Lyrics: ""})
}

func SessionCleanupWorker(store *Store, every time.Duration, stop <-chan struct{}) {
	t := time.NewTicker(every)
	defer t.Stop()

	for {
		select {
		case now := <-t.C:
			store.CleanupExpiredSessions(now)
		case <-stop:
			return
		}
	}
}
