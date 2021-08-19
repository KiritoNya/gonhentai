package nhentai

type Tag struct {
	id    int `validate:"min=0"`
	name  string
	url   string
	count int `validate:"min=1"`
}

type Group Tag
type Artist Tag
type Language Tag
type Category Tag
type Parody Tag
type Character Tag
