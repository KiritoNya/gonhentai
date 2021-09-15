package gonhentai

// TagsType is the data struct that describes the type of tag
type TagsType string

const (
	// Tag is the tag of doujinshi
	Tag TagsType = "tag"

	// Group is the group of doujinshi
	Group TagsType = "group"

	// Artist is the artist of doujinshi
	Artist TagsType = "artist"

	// Language is the language of doujinshi
	Language TagsType = "language"

	// Category is the category of doujinshi
	Category TagsType = "category"

	// Parody is the parody of doujinshi
	Parody TagsType = "parody"

	// Character is the character of doujinshi
	Character TagsType = "character"
)

// TagInfo is a struct that contains all the information related to the tags of a doujinshi
type TagInfo struct {
	Id    int
	Name  string
	Url   string
	Count int
	Type  TagsType
}
