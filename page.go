package nhentai

import (
	"errors"
	"strconv"
)

type Page struct {
	Number int    `validate:"min=0"`
	Url    string `validate:"doujin_page_url"`
	Image  *Image
}

// GetUrl is a function that generate the url page and assign its to the page object
func (p *Page) GetUrl(doujinId int) error {

	// Check validity of the number of page
	if p.Number <= 0 {
		return errors.New("Page number invalid")
	}

	doujinString := strconv.Itoa(doujinId)

	p.Url = DoujinPrefix + doujinString + "/" + strconv.Itoa(p.Number)
	return nil
}
