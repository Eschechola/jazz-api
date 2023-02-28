package repositories

import (
	"math/rand"
	"sort"
	"strconv"

	"golang.org/x/exp/slices"
	entities "jazz-api.com/src/jazz-api.com/entities"
)

var albums = []entities.Album{
	{Id: "1", Title: "Atravessei Sao Paulo", Artist: "Eduardo Duzz", Price: 99.99},
	{Id: "2", Title: "Cadente", Artist: "Eduardo Duzz", Price: 109.99},
	{Id: "3", Title: "Patricia", Artist: "Eduardo Duzz", Price: 89.99},
}

func sortAlbums() {
	sort.Slice(albums, func(current, next int) bool {
		return albums[current].Id < albums[next].Id
	})
}

func getAlbumIndex(id string) int {
	return slices.IndexFunc(albums,
		func(album entities.Album) bool {
			return album.Id == id
		})
}

func GetAlbums() []entities.Album {
	return albums
}

func GetSortedAlbums() []entities.Album {
	sortAlbums()
	return albums
}

func GetAlbumById(id string) entities.Album {
	var albumIndex = getAlbumIndex(id)

	return albums[albumIndex]
}

func AddAlbum(album entities.Album) entities.Album {
	var newId = rand.Intn(1000)

	var newAlbum = entities.Album{
		Id:     strconv.Itoa(newId),
		Title:  album.Title,
		Artist: album.Artist,
		Price:  album.Price,
	}

	albums = append(albums, newAlbum)

	return newAlbum
}

func UpdateAlbum(id string, title string, artist string, price float64) {
	var album = GetAlbumById(id)

	album.Title = title
	album.Artist = artist
	album.Price = price

	RemoveAlbum(album.Id)
	AddAlbum(album)
}

func RemoveAlbum(id string) []entities.Album {
	var albumIndex = getAlbumIndex(id)
	return append(albums[:albumIndex], albums...)
}
