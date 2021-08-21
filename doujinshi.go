package nhentai

import (
	"encoding/json"
	"errors"
	"io"
	"path/filepath"
	"strconv"
	"time"
)

type Doujinshi struct {
	Id           int
	Url          string
	MediaId      int
	Title        *Title
	CoverImage   *Image
	Thumbnail    *Image
	Pages        []*Page
	Scanlator    string
	UploadDate   time.Time
	Parodies     []*Parody
	Characters   []*Character
	Tags         []*Tag
	Artists      []*Artist
	Groups       []*Group
	Languages    []*Language
	Categories   []*Category
	NumPages     int
	NumFavorites int
	Related      []*Doujinshi
	Comments     []*Comment
}

type Title struct {
	English  string
	Japanese string
	Pretty   string
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
	err = json.Unmarshal(content, &d)
	if err != nil {
		return nil, err
	}
	d.Url = DoujinPrefix + strconv.Itoa(id)

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

// UnmarshalJSON is a json parser of doujinshi object
func (d *Doujinshi) UnmarshalJSON(b []byte) error {
	var rawDoujin map[string]json.RawMessage

	// Unmarshal
	err := json.Unmarshal(b, &rawDoujin)
	if err != nil {
		return err
	}

	// Parse title
	err = json.Unmarshal(rawDoujin["title"], &d.Title)
	if err != nil {
		return err
	}

	// Parse images
	var imagesRaw map[string]json.RawMessage
	err = json.Unmarshal(rawDoujin["images"], &imagesRaw)
	if err != nil {
		return err
	}

	// Parse pages
	var pageImagesRaw []json.RawMessage
	err = json.Unmarshal(imagesRaw["pages"], &pageImagesRaw)
	if err != nil {
		return err
	}

	for numPage, image := range pageImagesRaw {
		var i Image
		var p Page
		var err error

		err = json.Unmarshal(image, &i)
		if err != nil {
			return err
		}

		// Normalize image type
		i.Ext, err = normalizeExt(i.Ext)
		if err != nil {
			return err
		}

		p.Image = &i
		p.Number = numPage

		// Append
		d.Pages = append(d.Pages, &p)
	}

	// Parse cover image
	var coverImage Image
	err = json.Unmarshal(imagesRaw["cover"], &coverImage)
	if err != nil {
		return err
	}
	d.CoverImage = &coverImage

	// Parse thumbnail
	var thumbnail Image
	err = json.Unmarshal(imagesRaw["thumbnail"], &thumbnail)
	if err != nil {
		return err
	}
	d.Thumbnail = &thumbnail

	// Parse tags
	var tagsRaw []json.RawMessage
	err = json.Unmarshal(rawDoujin["tags"], &tagsRaw)
	if err != nil {
		return err
	}

	for _, tagRaw := range tagsRaw {
		var tagMap map[string]interface{}
		var tag Tag
		var err error

		err = json.Unmarshal(tagRaw, &tagMap)
		if err != nil {
			return err
		}

		tag.id = int(tagMap["id"].(float64))
		tag.name = tagMap["name"].(string)
		tag.count = int(tagMap["count"].(float64))
		tag.url = tagMap["url"].(string)

		// Filter by tag type
		switch tagMap["type"].(string) {
		case "parody":
			tagParody := Parody(tag)
			d.Parodies = append(d.Parodies, &tagParody)
		case "character":
			tagCharacter := Character(tag)
			d.Characters = append(d.Characters, &tagCharacter)
		case "tag":
			d.Tags = append(d.Tags, &tag)
		case "artist":
			tagArtist := Artist(tag)
			d.Artists = append(d.Artists, &tagArtist)
		case "group":
			tagGroup := Group(tag)
			d.Groups = append(d.Groups, &tagGroup)
		case "language":
			tagLanguage := Language(tag)
			d.Languages = append(d.Languages, &tagLanguage)
		case "category":
			tagCategory := Category(tag)
			d.Categories = append(d.Categories, &tagCategory)
		default:
			return errors.New("Tag type not found")
		}
	}

	// Parse id
	err = json.Unmarshal(rawDoujin["id"], &d.Id)
	if err != nil {
		return err
	}

	// Parse media id
	var mediaIdString string
	err = json.Unmarshal(rawDoujin["media_id"], &mediaIdString)
	if err != nil {
		return err
	}
	d.MediaId, err = strconv.Atoi(mediaIdString)
	if err != nil {
		return err
	}

	// Parse scanlator
	err = json.Unmarshal(rawDoujin["scanlator"], &d.Scanlator)
	if err != nil {
		return err
	}

	// Parse upload date
	var uploadUnixDate int64
	err = json.Unmarshal(rawDoujin["upload_date"], &uploadUnixDate)
	if err != nil {
		return err
	}
	d.UploadDate = time.Unix(uploadUnixDate, 0)

	// Parse number of pages
	err = json.Unmarshal(rawDoujin["num_pages"], &d.NumPages)
	if err != nil {
		return err
	}

	// Parse number of favorites
	err = json.Unmarshal(rawDoujin["num_favorites"], &d.NumFavorites)
	if err != nil {
		return err
	}

	return nil
}
