package nhentai

import (
	"encoding/json"
	"errors"
	"fmt"
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
	coverImage   *Image
	thumbnail    *Image
	pages        []*Page
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

func (t *Title) UnmarshalJSON(b []byte) error {
	var title map[string]json.RawMessage

	err := json.Unmarshal(title["english"], &t.english)
	if err != nil {
		return err
	}

	err = json.Unmarshal(title["japanese"], &t.japanese)
	if err != nil {
		return err
	}

	err = json.Unmarshal(title["pretty"], &t.pretty)
	if err != nil {
		return err
	}

	return nil
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

// Url is a function that return the url of doujinshi
func (d *Doujinshi) Url() string {
	return d.url
}

// MediaId is a function that return the mediaId of doujinshi
func (d *Doujinshi) MediaId() int {
	return d.mediaId
}

func (d *Doujinshi) allInfo() error {
	return nil
}

func (d *Doujinshi) UnmarshalJSON(b []byte) error {
	var rawDoujin map[string]interface{}

	// Unmarshal
	err := json.Unmarshal(b, &rawDoujin)
	if err != nil {
		return err
	}

	// Parse doujin title section
	var t Title
	titleMap := rawDoujin["title"].(map[string]interface{})
	t.english = titleMap["english"].(string)
	t.japanese = titleMap["japanese"].(string)
	t.pretty = titleMap["pretty"].(string)
	d.title = &t

	// Parse doujin image section
	imagesMap := rawDoujin["images"].(map[string]interface{})
	pagesArray := imagesMap["pages"].([]interface{})
	coverImage := imagesMap["cover"].(map[string]interface{})
	thumbImage := imagesMap["thumbnail"].(map[string]interface{})

	// Parse pages section
	for numPage, page := range pagesArray {
		var i Image
		var p Page

		pageMap := page.(map[string]interface{})

		// Get image info
		i.ext, err = normalizeExt(pageMap["t"].(string))
		if err != nil {
			return err
		}

		i.heigth = int(pageMap["h"].(float64))
		i.width = int(pageMap["w"].(float64))

		// Assign image at page
		p.image = &i

		// Set page number
		p.number = numPage

		// Append page to pages
		d.pages = append(d.pages, &p)
	}

	// Parse cover section
	var cover Image
	cover.ext, err = normalizeExt(coverImage["t"].(string))
	if err != nil {
		return err
	}
	cover.width = int(coverImage["w"].(float64))
	cover.heigth = int(coverImage["h"].(float64))
	d.coverImage = &cover

	// Parse thumbnail section
	var thumbnail Image
	thumbnail.ext, err = normalizeExt(thumbImage["t"].(string))
	if err != nil {
		return err
	}
	thumbnail.width = int(thumbImage["w"].(float64))
	thumbnail.heigth = int(thumbImage["h"].(float64))
	d.thumbnail = &thumbnail

	// Parse tags
	tagsMap := rawDoujin["tags"].([]interface{})
	for _, tag := range tagsMap {

		var tg Tag
		tagMap := tag.(map[string]interface{})

		tg.id = int(tagMap["id"].(float64))
		tg.name = tagMap["name"].(string)
		tg.url = BaseUrl + tagMap["url"].(string)
		tg.count = int(tagMap["count"].(float64))

		tagType := tagMap["type"].(string)
		switch tagType {
		case "parody":
			tagParody := Parody(tg)
			d.parodies = append(d.parodies, &tagParody)
		case "character":
			tagCharacter := Character(tg)
			d.characters = append(d.characters, &tagCharacter)
		case "tag":
			d.tags = append(d.tags, &tg)
		case "artist":
			tagArtist := Artist(tg)
			d.artists = append(d.artists, &tagArtist)
		case "group":
			tagGroup := Group(tg)
			d.groups = append(d.groups, &tagGroup)
		case "language":
			tagLanguage := Language(tg)
			d.languages = append(d.languages, &tagLanguage)
		case "category":
			tagCategory := Category(tg)
			d.categories = append(d.categories, &tagCategory)
		default:
			return errors.New("Tag type not found")
		}
	}

	// Parse base doujin info
	mediaId, err := strconv.Atoi(rawDoujin["media_id"].(string))
	if err != nil {
		return err
	}

	d.id = int(rawDoujin["id"].(float64))
	d.mediaId = mediaId
	d.title = &t
	d.scanlator = rawDoujin["scanlator"].(string)
	d.uploadDate = time.Unix(int64(rawDoujin["upload_date"].(float64)), 0)
	d.numPages = int(rawDoujin["num_pages"].(float64))
	d.numFavorites = int(rawDoujin["num_favorites"].(float64))

	data, err := json.Marshal(&d)
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	return nil
}

func (d *Doujinshi) MarshalJSON() error {
	/*j, err := json.Marshal(struct {
		Uuid string
		Name string
	}{
		Uuid: m.uuid,
		Name: m.Name,
	})
	if err != nil {
		return nil, err
	}
	return j, nil*/
}
