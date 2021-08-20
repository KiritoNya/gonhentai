package nhentai

import (
	"errors"
	"io"
	"path/filepath"
	"strconv"
	"time"
)

type Doujinshi struct {
	id           int
	url          string
	mediaId      int
	title        *Title
	pages        *Page
	scanlator    string
	uploadDate   time.Time
	parodies     []*Parody
	characters   []*Character
	tags         []*Tag
	artists      []*Artist
	groups       []*Group
	languages    []*Language
	categories   []*Category
	numPages     int
	numFavorites int
	related      []*Doujinshi
	comments     []*Comment
	raw          []byte
}

type Title struct {
	english  string
	japanese string
	pretty   string
}

// NewDoujinshiId is a constructor of the doujinshi object
func NewDoujinshiId(id int) (*Doujinshi, error) {

	var d Doujinshi

	// Validate id
	if !validateNhentaiId(id) {
		return nil, errors.New("Id not valid")
	}

	// Make url
	url, err := templateSolver(baseUrlApi+galleryApi, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}

	// Do request
	resp, err := ClientHttp.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the reader
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Prepare doujin info
	d.id = id
	d.url = DoujinPrefix + strconv.Itoa(id)
	d.raw = content

	return &d, nil
}

// NewDoujinshiUrl is a constructor of the doujinshi object
func NewDoujinshiUrl(url string) (*Doujinshi, error) {

	// Validate url encoding
	if !validateDoujinUrl(url) {
		return nil, errors.New("Doujinshi url not valid")
	}

	// Normalize url
	if url[len(url)-1] == '/' {
		url = url[:len(url)-1]
	}

	// Extract id from url
	idString := filepath.Base(url)
	id, err := strconv.Atoi(idString)
	if err != nil {
		return nil, err
	}

	// Create object
	doujin, err := NewDoujinshiId(id)
	if err != nil {
		return nil, err
	}

	return doujin, nil
}

// Id is a function that return the id of doujinshi
func (d *Doujinshi) Id() int {
	return d.id
}
