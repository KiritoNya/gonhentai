package nhentai

// Tag is a struct that contains all the information related to the tags of a doujinshi
type Tag struct {
	id    int `validate:"min=0"`
	name  string
	url   string
	count int `validate:"min=1"`
}

// Group is a specific type of Tag
type Group Tag

// Artist is a specific type of Tag
type Artist Tag

// Language is a specific type of Tag
type Language Tag

// Category is a specific type of Tag
type Category Tag

// Parody is a specific type of Tag
type Parody Tag

// Character is a specific type of Tag
type Character Tag
