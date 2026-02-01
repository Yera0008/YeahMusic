# YeahMusic – Assignment 4 (Milestone 2)

**Course:** Advanced Programming 1  
**Assignment:** 4 – Project Milestone 2 (Core System Implementation)  
**Team:** **BFF TEAM**  
**Members:** Ayat, Ernar  

---

## Overview

This project is the **Milestone 2 backend implementation** of the YeahMusic system.  
The goal of this milestone is to demonstrate a **working core backend** that follows the architecture and data model defined in **Assignment 3**.

The system is implemented as a **monolithic Go HTTP API**.  
No frontend or UI is included, as it is **not required at this stage**.

---

## Implemented Features

The backend provides the following core functionality:

- User registration and login
- Session-based authentication (token-based)
- Music catalog browsing:
  - Artists
  - Albums
  - Tracks
- Playlist management:
  - Create playlist
  - List user playlists
  - Add track to playlist
  - Remove track from playlist
- Background session cleanup using a goroutine

All data is stored **in-memory** and protected with concurrency-safe access.

---

## Technical Details

- Language: **Go**
- HTTP server: `net/http`
- Data format: **JSON**
- Storage: In-memory (maps + mutexes)
- Concurrency:
  - `sync.RWMutex` for safe concurrent access
  - Background worker (goroutine) for expired session cleanup
- Architecture:
  - `handlers` – HTTP layer
  - `services` – business logic
  - `models` – domain models (based on ERD from Assignment 3)

---

## Project Structure

```
.
├─ go.mod
├─ main.go
├─ README.md
└─ internal
   ├─ handlers
   ├─ models
   └─ services
```

---

## How to Run

Make sure Go is installed.

```bash
go run .
```

If the server starts successfully, you will see:

```
API listening on :8080
```

---

## API Endpoints

### Health Check
```
GET /health
```

### Authentication
```
POST /api/register
POST /api/login
```

### Catalog
```
GET /api/artists
GET /api/albums?artist_id={id}
GET /api/tracks?album_id={id}
```

### Playlists (requires Authorization header)
```
POST   /api/playlists
GET    /api/playlists
POST   /api/playlists/{playlistId}/tracks
DELETE /api/playlists/{playlistId}/tracks/{trackId}
```

**Authorization header format:**
```
Authorization: Bearer <token>
```

---

## Notes

- This is **Milestone 2**, not the final project.
- Full UI, authentication hardening, database integration, and optimization are intentionally not implemented yet.
- The implementation strictly follows the approved design from Assignment 3.

---

## Team

**BFF TEAM**  
Ayat  
Ernar
