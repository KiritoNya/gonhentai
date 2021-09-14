package nhentai

type TagsType string

const (
	Tag        TagsType = "tag"
	Group      TagsType = "group"
	Artist     TagsType = "artist"
	Language   TagsType = "language"
	Category   TagsType = "category"
	ParodyType TagsType = "parody"
	Character  TagsType = "character"
)

// TagInfo is a struct that contains all the information related to the tags of a doujinshi
type TagInfo struct {
	Id    int
	Name  string
	Url   string
	Count int
	Type  TagsType
}
