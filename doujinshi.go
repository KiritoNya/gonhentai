package nhentai

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// Doujinshi is the data struct that describes a manga or a doujinshi
type Doujinshi struct {
	Id           int
	Url          string
	MediaId      int
	Title        *Title
	CoverImage   *Cover
	Thumbnail    *Thumbnail
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

// Title is the data struct that describes the doujinshi title
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
	d.Url = DoujinBaseUrl + strconv.Itoa(id)

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

// GetRelated is a function that gets the related doujinshi and assign them to the doujinshi object
func (d *Doujinshi) GetRelated() error {

	type RespJson struct {
		Result []*Doujinshi
	}
	var rj RespJson

	// Check doujinshi id
	if d.Id == 0 {
		return errors.New("Id of doujinshi not setted")
	}

	// Make url
	tmpl, err := templateSolver(
		searchRelatedApi,
		map[string]interface{}{
			"id": d.Id,
		},
	)

	// Check template error
	if err != nil {
		return err
	}

	// Do request
	resp, err := ClientHttp.Get(baseUrlApi + "/" + tmpl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Parse json response body
	err = json.Unmarshal(content, &rj)
	if err != nil {
		return err
	}

	// Assign related doujinshi to doujinshi object
	d.Related = rj.Result
	return nil
}

// GetComments is a function that gets the related comments of doujinshi and assign them to the doujinshi object
func (d *Doujinshi) GetComments() error {

	var c []*Comment

	// Check id of doujinshi object
	if d.Id == 0 {
		return errors.New("Id not setted")
	}

	// Make url
	tmpl, err := templateSolver(commentsApi, map[string]interface{}{"id": d.Id})
	if err != nil {
		return err
	}

	// Do request
	res, err := ClientHttp.Get(baseUrlApi + tmpl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Read response body
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// Parse json response body
	err = json.Unmarshal(content, &c)
	if err != nil {
		return err
	}

	d.Comments = c
	return nil
}

// Save is a function that download all pages of doujinshi in the specified directory. The name of image can be described by template.
func (d *Doujinshi) Save(dirPathTmpl string, perm os.FileMode) error {

	// Check if pages is setted
	if d.Pages == nil {
		return errors.New("Doujinshi pages not setted")
	}

	// Check if numPages is setted
	if d.NumPages == 0 {
		return errors.New("Pages number not setted")
	}

	// Foreach pages
	for pageNum, pag := range d.Pages {

		type Template struct {
			Doujinshi *Doujinshi
			Page      *Page
		}

		// Prepare template data
		var t Template
		t.Doujinshi = d
		t.Page = pag

		// Generate path from template
		tmpl, err := templateSolver(dirPathTmpl, t)
		if err != nil {
			return err
		}

		// Check page.Data
		if pag.Data == nil {

			// Check if there is url
			if pag.Url == "" {
				err := pag.urlService(d.MediaId, strconv.Itoa(pageNum+1), page)
				if err != nil {
					return err
				}
			}

			// Check name
			if pag.Name == "" {
				pag.Name = filepath.Base(tmpl)
			}

			// Get data of image
			err := pag.GetData()
			if err != nil {
				return err
			}
		}

		// Download image
		err = pag.Save(tmpl, perm)
		if err != nil {
			return err
		}
	}

	return nil
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

	for pageNum, img := range pageImagesRaw {
		var p Page
		var err error

		err = json.Unmarshal(img, &p)
		if err != nil {
			return err
		}

		// Set page number
		p.Num = pageNum + 1

		// Normalize image type
		p.Ext, err = normalizeExt(p.Ext)
		if err != nil {
			return err
		}

		// Append
		d.Pages = append(d.Pages, &p)
	}

	// Parse cover image
	var coverImage Cover
	err = json.Unmarshal(imagesRaw["cover"], &coverImage)
	if err != nil {
		return err
	}
	//cover := Cover(coverImage)
	d.CoverImage = &coverImage

	// Parse thumbnail
	var thumbnail Thumbnail
	err = json.Unmarshal(imagesRaw["thumbnail"], &thumbnail)
	if err != nil {
		return err
	}
	//thumbnailObj := Thumbnail(thumbnail)
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

		// Check if id is a string
		var idString string
		err2 := json.Unmarshal(rawDoujin["id"], &idString)
		if err2 != nil {
			// It isn't a string, it's a real error
			return err2
		}

		// Convert string to int
		idInt, err := strconv.Atoi(idString)
		if err2 != nil {
			return err
		}

		// Assign id to id field
		d.Id = idInt
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
