package nhentai_test

import (
	"github.com/KiritoNya/nhentai"
	"testing"
)

func TestPage_GetUrl(t *testing.T) {
	page := nhentai.Page{
		Number: 10,
		Url:    "https://google.com",
	}

	t.Log(page)
}
