package nhentai_test

import (
	"github.com/KiritoNya/nhentai"
	"testing"
)

func TestPageImage_GetUrl(t *testing.T) {
	var p nhentai.Page

	p.Ext = "jpg"
	err := p.GetUrl(mediaId, numPage)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("PageImage_Url: ", p.Url)
	t.Log("PageImage_GetUrl: [OK]")
}

func TestCover_GetUrl(t *testing.T) {
	var c nhentai.Cover
	c.Ext = "jpg"

	err := c.GetUrl(mediaId)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("CoverImage_Url", c.Url)
	t.Log("CoverImage_GetUrl [OK]")
}

func TestThumbnail_GetUrl(t *testing.T) {
	var thumb nhentai.Thumbnail
	thumb.Ext = "jpg"

	err := thumb.GetUrl(mediaId, numPage)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ThumbnailImage_Url", thumb.Url)
	t.Log("ThumbnailImage_GetUrl [OK]")
}

func TestImage_GetSize(t *testing.T) {
	var p nhentai.Page
	p.Url = pageUrl

	err := p.GetSize()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Image_Size", p.Size)
	t.Log("Image_GetSize [OK]")
}

func TestImage_GetData(t *testing.T) {
	var p nhentai.Page
	p.Url = pageUrl

	err := p.GetData()
	if err != nil {
		t.Fatal(err)
	}

	if p.Data == nil {
		t.Fatal("Image data is empty")
	}

	t.Log("Image_GetData [OK]")
}

func TestImage_GenerateName(t *testing.T) {
	var p nhentai.Page
	p.Ext = "jpg"

	err := p.GenerateName("image", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Image_Name:", p.Name)
	t.Log("GetImageName [OK]")
}

func TestImage_Save(t *testing.T) {
	var p nhentai.Page
	p.Url = pageUrl
	p.Ext = ".jpg"

	err := p.Save("./img.jpeg", 0644)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Image_Save [OK]")
}
