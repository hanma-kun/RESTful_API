package database

import (
	"github.com/hanma-kun/REST/models"
	"sync"
)

var albums = []models.Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

var albumMutex sync.RWMutex

// GetAlbums retrieves all albums from the database.
func GetAlbums() []models.Album {
	albumMutex.RLock()
	defer albumMutex.RUnlock()

	return albums
}

// GetAlbumByID retrieves an album by its ID from the database.
func GetAlbumByID(id string) (models.Album, bool) {
	albumMutex.RLock()
	defer albumMutex.RUnlock()

	for _, a := range albums {
		if a.ID == id {
			return a, true
		}
	}
	return models.Album{}, false
}

// AddAlbum adds a new album to the database.
func AddAlbum(newAlbum models.Album) {
	albumMutex.Lock()
	defer albumMutex.Unlock()

	albums = append(albums, newAlbum)
}