package nhentai

type Page struct {
	number       int    `validate:"min=0"`
	url          string `validate:"doujin_page_url"`
	image        *Image
	nextPage     *Page
	previousPage *Page
	raw          string `validate:"json"`
}
