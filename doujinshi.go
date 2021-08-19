package nhentai

import "time"

type Doujinshi struct {
	id           int `validate:"required,sauce"`
	mediaId      int `validate:"omitempty,min=0"`
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
	numPages     int `validate:"min=1"`
	numFavorites int `validate:"min=0"`
	related      []*Doujinshi
	comments     []*Comment
	raw          string `validate:"json"`
}

type Title struct {
	english  string
	japanese string
	pretty   string
}
