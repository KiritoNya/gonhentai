package raw

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"
)

// DoujinRaw contains the json returned by the API for doujinshi
type DoujinRaw struct {
	info map[string]json.RawMessage
	img  map[string]json.RawMessage
	tags []*TagRaw
}

type User struct {
	Id   int `validate:"-"`
	Info struct {
		Name  string `validate:"presence,min=2,max=32"`
		Email string `validate:"email,required"`
	}
}

// NewDoujinRawUrl creates a new Doujinshi raw object with the url
func NewDoujinRawUrl(clientHttp *http.Client, url string) (*DoujinRaw, error) {
	var rw map[string]json.RawMessage

	// Do request
	resp, err := clientHttp.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the reader
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal
	err = json.Unmarshal(content, &rw)
	if err != nil {
		return nil, err
	}

	// Create object
	r, err := NewDoujinRaw(rw)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// NewDoujinRaw creates a new Doujinshi raw object with the raw message
func NewDoujinRaw(data map[string]json.RawMessage) (*DoujinRaw, error) {
	var r DoujinRaw

	// Fill info
	r.info = data

	// Fill img
	err := json.Unmarshal(data["images"], &r.img)
	if err != nil {
		return nil, err
	}

	// Fill tags
	err = r.getTags()
	if err != nil {
		return nil, err
	}

	// Fill infoRaw
	delete(data, "images")
	delete(data, "tags")
	r.info = data

	return &r, nil
}

// Id is a function that returns the id of the doujinshi raw
func (dr *DoujinRaw) Id() (id int, err error) {

	// Unmarshal data
	err = json.Unmarshal(dr.info["id"], &id)
	if err != nil {

		// Check if id is a string
		var idString string
		err2 := json.Unmarshal(dr.info["id"], &idString)
		if err2 != nil {
			// It isn't a string, it's a real error
			return 0, err2
		}

		// Convert string to int
		id, err = strconv.Atoi(idString)
		if err2 != nil {
			return 0, err
		}
	}

	return id, nil
}

// MediaId is a function that returns the media id of the doujinshi raw
func (dr *DoujinRaw) MediaId() (mediaId int, err error) {

	var mediaIdString string
	// Unmarshal data
	err = json.Unmarshal(dr.info["media_id"], &mediaIdString)
	if err != nil {
		return 0, err
	}

	// Convert string to int
	mediaId, err = strconv.Atoi(mediaIdString)
	if err != nil {
		return 0, err
	}

	return mediaId, nil
}

// Title is a function that returns the title of the doujinshi raw
func (dr *DoujinRaw) Title() (title TitleRaw, err error) {
	// Unmarshal data
	err = json.Unmarshal(dr.info["title"], &title.Data)
	if err != nil {
		return TitleRaw{}, err
	}

	return title, nil
}

// CoverImage is a function that returns the cover of the doujinshi raw
func (dr *DoujinRaw) CoverImage() (cover ImageRaw, err error) {
	// Unmarshal data
	err = json.Unmarshal(dr.img["cover"], &cover.Data)
	if err != nil {
		return ImageRaw{}, err
	}
	return cover, nil
}

// Thumbnail is a function that returns the thumbnail of the doujinshi raw
func (dr *DoujinRaw) Thumbnail() (thumb ImageRaw, err error) {
	// Unmarshal data
	err = json.Unmarshal(dr.img["thumbnail"], &thumb.Data)
	if err != nil {
		return ImageRaw{}, err
	}
	return thumb, nil
}

// Pages is a function that returns all pages of the doujinshi raw
func (dr *DoujinRaw) Pages() (pages []*ImageRaw, err error) {
	var pageImagesRaw []json.RawMessage

	// Unmarshal
	err = json.Unmarshal(dr.img["pages"], &pageImagesRaw)
	if err != nil {
		return nil, err
	}

	// Foreach page
	for _, img := range pageImagesRaw {
		var ir ImageRaw
		var err error

		// Unmarshal
		err = json.Unmarshal(img, &ir.Data)
		if err != nil {
			return nil, err
		}

		pages = append(pages, &ir)
	}

	return pages, nil
}

// Scanlator is a function that returns the scanlator of the doujinshi raw
func (dr *DoujinRaw) Scanlator() (scanlator string, err error) {
	// Unmarshal data
	err = json.Unmarshal(dr.info["scanlator"], &scanlator)
	if err != nil {
		return "", err
	}

	return scanlator, nil
}

// UploadDate is a function that returns the upload date of the doujinshi raw
func (dr *DoujinRaw) UploadDate() (uploadTime time.Time, err error) {
	// Parse upload date
	var uploadUnixDate int64
	err = json.Unmarshal(dr.info["upload_date"], &uploadUnixDate)
	if err != nil {
		return time.Time{}, err
	}
	uploadTime = time.Unix(uploadUnixDate, 0)

	return uploadTime, nil
}

// Parodies is a function that returns the parodies list of the doujinshi raw
func (dr *DoujinRaw) Parodies() (parodies []*TagRaw, err error) {
	return dr.filterTags("parody")
}

// Characters is a function that returns the characters list of the doujinshi raw
func (dr *DoujinRaw) Characters() (characters []*TagRaw, err error) {
	return dr.filterTags("character")
}

// Tags is a function that returns the tags list of the doujinshi raw
func (dr *DoujinRaw) Tags() (tags []*TagRaw, err error) {
	return dr.filterTags("tag")
}

// Artists is a function that returns the artists list of the doujinshi raw
func (dr *DoujinRaw) Artists() (artists []*TagRaw, err error) {
	return dr.filterTags("artist")
}

// Groups is a function that returns the groups list of the doujinshi raw
func (dr *DoujinRaw) Groups() (groups []*TagRaw, err error) {
	return dr.filterTags("group")
}

// Languages is a function that returns the languages list of the doujinshi raw
func (dr *DoujinRaw) Languages() (languages []*TagRaw, err error) {
	return dr.filterTags("language")
}

// Categories is a function that returns the categories list of the doujinshi raw
func (dr *DoujinRaw) Categories() (categories []*TagRaw, err error) {
	return dr.filterTags("category")
}

// NumPages is a function that returns the number of pages of the doujinshi raw
func (dr *DoujinRaw) NumPages() (numPages int, err error) {
	err = json.Unmarshal(dr.info["num_pages"], &numPages)
	if err != nil {
		return 0, err
	}

	return numPages, nil
}

// NumFavorites is a function that returns the number of favorites of the doujinshi raw
func (dr *DoujinRaw) NumFavorites() (numFavorites int, err error) {
	err = json.Unmarshal(dr.info["num_favorites"], &numFavorites)
	if err != nil {
		return 0, err
	}

	return numFavorites, nil
}

// All is a function that returns all info of the doujinshi
func (dr *DoujinRaw) All() (doujinMap map[string]interface{}, err error) {
	doujinMap = make(map[string]interface{})

	// Get Id of doujinshi
	id, err := dr.Id()
	if err != nil {
		return nil, err
	}

	// Get MediaId of doujinshi
	mediaId, err := dr.MediaId()
	if err != nil {
		return nil, err
	}

	// Get title of doujinshi
	titleRaw, err := dr.Title()
	if err != nil {
		return nil, err
	}

	title, err := titleRaw.All()
	if err != nil {
		return nil, err
	}

	// Get cover of doujinshi
	coverRaw, err := dr.CoverImage()
	if err != nil {
		return nil, err
	}

	cover, err := coverRaw.All()
	if err != nil {
		return nil, err
	}

	// Get cover of doujinshi
	thumbRaw, err := dr.Thumbnail()
	if err != nil {
		return nil, err
	}

	thumb, err := thumbRaw.All()
	if err != nil {
		return nil, err
	}

	// Get pages of doujinshi
	pagesRaw, err := dr.Pages()
	if err != nil {
		return nil, err
	}

	var pages []map[string]interface{}
	for _, pageRaw := range pagesRaw {
		page, err := pageRaw.All()
		if err != nil {
			return nil, err
		}

		pages = append(pages, page)
	}

	// Get scanlator of doujinshi
	scanlator, err := dr.Scanlator()
	if err != nil {
		return nil, err
	}

	// Get uploadDate of doujinshi
	uploadDate, err := dr.UploadDate()
	if err != nil {
		return nil, err
	}

	// Get parodies of doujinshi
	parodiesRaw, err := dr.Parodies()
	if err != nil {
		return nil, err
	}

	var parodies []map[string]interface{}
	for _, parodyRaw := range parodiesRaw {
		parody, err := parodyRaw.All()
		if err != nil {
			return nil, err
		}

		parodies = append(parodies, parody)
	}

	// Get characters of doujinshi
	charactersRaw, err := dr.Characters()
	if err != nil {
		return nil, err
	}

	var characters []map[string]interface{}
	for _, characterRaw := range charactersRaw {
		character, err := characterRaw.All()
		if err != nil {
			return nil, err
		}

		characters = append(characters, character)
	}

	// Get tags of doujinshi
	tagsRaw, err := dr.Tags()
	if err != nil {
		return nil, err
	}

	var tags []map[string]interface{}
	for _, tagRaw := range tagsRaw {
		tag, err := tagRaw.All()
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	// Get artists of doujinshi
	artistsRaw, err := dr.Tags()
	if err != nil {
		return nil, err
	}

	var artists []map[string]interface{}
	for _, artistRaw := range artistsRaw {
		artist, err := artistRaw.All()
		if err != nil {
			return nil, err
		}

		artists = append(artists, artist)
	}

	// Get groups of doujinshi
	groupsRaw, err := dr.Tags()
	if err != nil {
		return nil, err
	}

	var groups []map[string]interface{}
	for _, groupRaw := range groupsRaw {
		group, err := groupRaw.All()
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)
	}

	// Get languages of doujinshi
	languagesRaw, err := dr.Tags()
	if err != nil {
		return nil, err
	}

	var languages []map[string]interface{}
	for _, langRaw := range languagesRaw {
		lang, err := langRaw.All()
		if err != nil {
			return nil, err
		}

		languages = append(languages, lang)
	}

	// Get categories of doujinshi
	categoriesRaw, err := dr.Tags()
	if err != nil {
		return nil, err
	}

	var categories []map[string]interface{}
	for _, categoryRaw := range categoriesRaw {
		category, err := categoryRaw.All()
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	// Get number of pages of doujinshi
	numPages, err := dr.NumPages()
	if err != nil {
		return nil, err
	}

	// Get number of favorites of doujinshi
	numFavorites, err := dr.MediaId()
	if err != nil {
		return nil, err
	}

	doujinMap["Id"] = id
	doujinMap["MediaId"] = mediaId
	doujinMap["Title"] = title
	doujinMap["Cover"] = cover
	doujinMap["Thumbnail"] = thumb
	doujinMap["Pages"] = pages
	doujinMap["Scanlator"] = scanlator
	doujinMap["UploadDate"] = uploadDate
	doujinMap["Parodies"] = parodies
	doujinMap["Characters"] = characters
	doujinMap["Tags"] = tags
	doujinMap["Artists"] = artists
	doujinMap["Groups"] = groups
	doujinMap["Languages"] = languages
	doujinMap["Categories"] = categories
	doujinMap["NumPages"] = numPages
	doujinMap["NumFavorites"] = numFavorites

	return doujinMap, nil
}

// getTags is a function that returns the tags list of the doujinshi raw
func (dr *DoujinRaw) getTags() error {
	var tagsRaw []json.RawMessage

	// Unmarshal
	err := json.Unmarshal(dr.info["tags"], &tagsRaw)
	if err != nil {
		return err
	}

	for _, tagRaw := range tagsRaw {
		var t TagRaw

		err := json.Unmarshal(tagRaw, &t.Data)
		if err != nil {
			return err
		}

		dr.tags = append(dr.tags, &t)
	}

	return nil
}

// filterTags is a function that returns the filtered tags list
func (dr *DoujinRaw) filterTags(filter string) (filtered []*TagRaw, err error) {
	// Check if there is tags in the object
	if dr.tags == nil {
		err := dr.getTags()
		if err != nil {
			return nil, err
		}
	}

	for _, tag := range dr.tags {
		tagType, err := tag.Type()
		if err != nil {
			return nil, err
		}

		if tagType == filter {
			filtered = append(filtered, tag)
		}
	}

	return filtered, nil
}
