package nhentai

import (
	"encoding/json"
	"errors"
	"github.com/KiritoNya/nhentai/internal/pkg/raw"
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
	raw          json.RawMessage
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

// GetUrl is a function that gets the url of the doujinshi and assign them to the doujinshi object
func (d *Doujinshi) GetUrl() error {

	// Validate id
	if !validateNhentaiId(d.Id) {
		return errors.New("Id not valid")
	}

	d.Url = DoujinBaseUrl + strconv.Itoa(d.Id)
	return nil
}

// GetRelated is a function that gets the related doujinshi and assign it to the doujinshi object
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
	var tempMap map[string]json.RawMessage

	// Unmarshal
	err := json.Unmarshal(b, &tempMap)
	if err != nil {
		return err
	}

	// Create raw object
	rd, err := raw.NewDoujinRaw(tempMap)
	if err != nil {
		return err
	}

	// Get data
	doujinData, err := rd.All()
	if err != nil {
		return err
	}

	// Assign values
	d.Id = doujinData["Id"].(int)
	d.MediaId = doujinData["MediaId"].(int)
	d.Scanlator = doujinData["Scanlator"].(string)
	d.UploadDate = doujinData["UploadDate"].(time.Time)
	d.NumPages = doujinData["NumPages"].(int)
	d.NumFavorites = doujinData["NumFavorites"].(int)

	// Assign title values
	var title Title
	titleMap := doujinData["Title"].(map[string]string)
	title.English = titleMap["English"]
	title.Japanese = titleMap["Japanese"]
	title.Pretty = titleMap["Pretty"]
	d.Title = &title

	// Assign cover values
	var c Cover
	coverMap := doujinData["Cover"].(map[string]interface{})
	c.Ext = coverMap["Ext"].(string)
	c.Width = coverMap["Width"].(int)
	c.Heigth = coverMap["Height"].(int)
	d.CoverImage = &c

	// Assign thumbnail values
	var t Thumbnail
	thumbMap := doujinData["Thumbnail"].(map[string]interface{})
	t.Ext = thumbMap["Ext"].(string)
	t.Width = thumbMap["Width"].(int)
	t.Heigth = thumbMap["Height"].(int)
	d.Thumbnail = &t

	// Assign pages values
	pagesMap := doujinData["Pages"].([]map[string]interface{})
	for numPage, pageMap := range pagesMap {
		var p Page
		p.Ext = pageMap["Ext"].(string)
		p.Width = pageMap["Width"].(int)
		p.Heigth = pageMap["Height"].(int)
		p.Num = numPage
		d.Pages = append(d.Pages, &p)
	}

	// Assign parodies values
	parodiesMap := doujinData["Parodies"].([]map[string]interface{})
	for _, parodyMap := range parodiesMap {
		var p Parody
		p.Id = parodyMap["Id"].(int)
		p.Name = parodyMap["Name"].(string)
		p.Url = parodyMap["Url"].(string)
		p.Count = parodyMap["Count"].(int)

		d.Parodies = append(d.Parodies, &p)
	}

	// Assign characters values
	charactersMap := doujinData["Characters"].([]map[string]interface{})
	for _, characterMap := range charactersMap {
		var c Character
		c.Id = characterMap["Id"].(int)
		c.Name = characterMap["Name"].(string)
		c.Url = characterMap["Url"].(string)
		c.Count = characterMap["Count"].(int)

		d.Characters = append(d.Characters, &c)
	}

	// Assign tags values
	tagsMap := doujinData["Tags"].([]map[string]interface{})
	for _, tagMap := range tagsMap {
		var t Tag
		t.Id = tagMap["Id"].(int)
		t.Name = tagMap["Name"].(string)
		t.Url = tagMap["Url"].(string)
		t.Count = tagMap["Count"].(int)

		d.Tags = append(d.Tags, &t)
	}

	// Assign artists values
	artistsMap := doujinData["Artists"].([]map[string]interface{})
	for _, artistMap := range artistsMap {
		var a Artist
		a.Id = artistMap["Id"].(int)
		a.Name = artistMap["Name"].(string)
		a.Url = artistMap["Url"].(string)
		a.Count = artistMap["Count"].(int)

		d.Artists = append(d.Artists, &a)
	}

	// Assign groups values
	groupsMap := doujinData["Groups"].([]map[string]interface{})
	for _, groupMap := range groupsMap {
		var g Group
		g.Id = groupMap["Id"].(int)
		g.Name = groupMap["Name"].(string)
		g.Url = groupMap["Url"].(string)
		g.Count = groupMap["Count"].(int)

		d.Groups = append(d.Groups, &g)
	}

	// Assign language values
	languagesMap := doujinData["Languages"].([]map[string]interface{})
	for _, languageMap := range languagesMap {
		var l Language
		l.Id = languageMap["Id"].(int)
		l.Name = languageMap["Name"].(string)
		l.Url = languageMap["Url"].(string)
		l.Count = languageMap["Count"].(int)

		d.Languages = append(d.Languages, &l)
	}

	// Assign categories values
	categoriesMap := doujinData["Categories"].([]map[string]interface{})
	for _, categoryMap := range categoriesMap {
		var c Category
		c.Id = categoryMap["Id"].(int)
		c.Name = categoryMap["Name"].(string)
		c.Url = categoryMap["Url"].(string)
		c.Count = categoryMap["Count"].(int)

		d.Categories = append(d.Categories, &c)
	}

	return nil
}
