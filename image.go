package nhentai

import (
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type imageType int

const (
	thumb imageType = iota
	cover
	page
)

type image struct {
	Name   string `validate:"omitempty" json:",omitempty"`
	Url    string `validate:"omitempty,nhentai_img_url" json:",omitempty"`
	Size   int64  `validate:"omitempty,min=0" json:",omitempty"`
	Heigth int    `validate:"omitempty,min=0" json:"h,omitempty"`
	Width  int    `validate:"omitempty,min=0" json:"w,omitempty"`
	Ext    string `validate:"omitempty,eq=jpg,eq=png,eq=gif" json:"t,omitempty"`
	Data   []byte
}

type Cover struct {
	image
}

type Thumbnail struct {
	image
}

type Page struct {
	image
}

type Avatar struct {
	image
}

// GetSize is a function that gets the size of the image and assign it to the object
func (i *image) GetSize() error {

	// Check if the url is setted in the object
	if !validateNhentaiImageUrl(i.Url) {
		return errors.New("Url of the image object not valid")
	}

	// Do the request
	res, err := ClientHttp.Head(i.Url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check if the image exist
	if res.StatusCode == 404 {
		return errors.New("Image not found")
	}

	// Get the size from the header and convert it to int64
	sizeString := res.Header.Get("Content-Length")
	sizeInt, err := strconv.Atoi(sizeString)
	if err != nil {
		return err
	}

	// Assign the size to the image object
	i.Size = int64(sizeInt)
	return nil
}

// GetData is a function that gets the data of image and assign it to the field object Data
func (i *image) GetData() error {

	// Check if the url is setted
	if i.Url == "" {
		return errors.New("Url not setted")
	}

	// Do request
	res, err := ClientHttp.Get(i.Url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check status code
	if res.StatusCode != 200 {
		return errors.New("Url not valid")
	}

	// Read response body
	content, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	i.Data = content
	return nil
}

// GenerateName is a function that generates the image name from the object data and assign it to the field object Name
func (i *image) GenerateName(name, suffix string) error {

	// Check name
	if name == "" {
		return errors.New("Name not valid")
	}

	// Check image extension
	if i.Ext == "" {
		return errors.New("Image extension not setted")
	}

	// Check if the extension is setted
	if i.Ext != "" {
		// Check if it is a valid extension
		if !validateImageType(i.Ext) {
			extWithoutDot := strings.Replace(i.Ext, ".", "", -1)
			if validateImageType(extWithoutDot) {
				i.Ext = extWithoutDot
			} else {
				return errors.New("Image type not valid. Image type must be [jpg|png|gif]")
			}
		}
	}

	i.Name = name + "." + i.Ext + suffix
	return nil
}

// Save is a function that save the image on disk.
func (i *image) Save(path string, perm os.FileMode) error {

	// Check if data field is empty
	if i.Data == nil {
		err := i.GetData()
		if err != nil {
			return err
		}
	}

	// Verify if file already exists
	if _, err := os.Stat(path); err == nil {
		return errors.New(fmt.Sprintf(`File "%s" already exists`, path))
	}

	// Create empty file
	err := os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return err
	}

	// Save file
	err = os.WriteFile(path, i.Data, perm)
	if err != nil {
		return err
	}

	return nil
}

// GetUrl is a function that get the url of thumbnail image and assign it to the Thumbnail object
func (t *Thumbnail) GetUrl(mediaId, numPage int) error {

	// Check if the page number is valid
	if numPage <= 0 {
		return errors.New("Page number not valid")
	}

	return t.urlService(mediaId, strconv.Itoa(numPage), thumb)
}

// GetUrl is a function that get the url of page image and assign it to the Page object
func (pi *Page) GetUrl(mediaId, numPage int) error {

	// Check if the page number is valid
	if numPage <= 0 {
		return errors.New("Page number not valid")
	}

	return pi.urlService(mediaId, strconv.Itoa(numPage), page)
}

// GetUrl is a function that get the url of cover image and assign it to the Cover object
func (c *Cover) GetUrl(mediaId int) error {
	return c.urlService(mediaId, "cover", cover)
}

// GetUrl is a function that gets the url of the image and assign it to the image object
func (i *image) urlService(mediaId int, pageNum string, imgType imageType) error {

	var baseUrl string

	if mediaId <= 0 {
		return errors.New("Media id not valid")
	}

	if imgType != cover {
		if num, _ := strconv.Atoi(pageNum); num <= 0 {
			return errors.New("Number of page not valid")
		}
	}

	if !validateImageType(i.Ext) {
		extWithoutDot := strings.Replace(i.Ext, ".", "", -1)
		if validateImageType(extWithoutDot) {
			i.Ext = extWithoutDot
		} else {
			return errors.New("Image type not valid. Image type must be [jpg|png|gif]")
		}
	}

	switch imgType {
	case thumb:
		baseUrl = ThumbnailBaseUrl
		pageNum = pageNum + "t"
	case cover:
		baseUrl = ThumbnailBaseUrl
	case page:
		baseUrl = ImageBaseUrl
	}

	// Create url of image with template
	url, err := templateSolver(imageCompleteUrl, map[string]interface{}{
		"baseImageUrl": baseUrl,
		"mediaId":      mediaId,
		"numPage":      pageNum,
		"ext":          i.Ext,
	})

	// Check the template error
	if err != nil {
		return err
	}

	// Check if the image exists
	resp, err := ClientHttp.Head(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return errors.New("Image not found, check the input parametres")
	}

	// Assign url to the object
	i.Url = url

	return nil
}
