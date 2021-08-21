package nhentai

type Page struct {
	Number       int    `validate:"min=0"`
	Url          string `validate:"doujin_page_url"`
	Image        *Image
	NextPage     *Page
	PreviousPage *Page
}
