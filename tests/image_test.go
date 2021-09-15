package gonhentai_test

import (
	"log"
	"testing"
)

func TestPageImage_GetUrl(t *testing.T) {
	var p gonhentai.Page

	p.Ext = "jpg"
	err := p.GetUrl(mediaId, numPage)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("PageImage_Url: ", p.Url)
	t.Log("PageImage_GetUrl: [OK]")
}

func TestCover_GetUrl(t *testing.T) {
	var c gonhentai.Cover
	c.Ext = "jpg"

	err := c.GetUrl(mediaId)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("CoverImage_Url", c.Url)
	t.Log("CoverImage_GetUrl [OK]")
}

func TestThumbnail_GetUrl(t *testing.T) {
	var thumb gonhentai.Thumbnail
	thumb.Ext = "jpg"

	err := thumb.GetUrl(mediaId, numPage)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ThumbnailImage_Url", thumb.Url)
	t.Log("ThumbnailImage_GetUrl [OK]")
}

func TestImage_GetSize(t *testing.T) {
	var p gonhentai.Page
	p.Url = pageUrl

	err := p.GetSize()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Image_Size", p.Size)
	t.Log("Image_GetSize [OK]")
}

func TestImage_GetData(t *testing.T) {
	var p gonhentai.Page
	p.Url = pageUrl
	p.Ext = ".jpg"

	gonhentai.UseProgressBar = true

	err := p.GetData()
	if err != nil {
		log.Fatal(err)
	}

	if p.Data == nil {
		t.Fatal("Image data is empty")
	}

	t.Log("Image_GetData [OK]")
}

func TestImage_GenerateName(t *testing.T) {
	var p gonhentai.Page
	p.Ext = "jpg"

	err := p.GenerateName("image", "")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Image_Name:", p.Name)
	t.Log("GetImageName [OK]")
}

func TestImage_Save(t *testing.T) {
	var p1 gonhentai.Page
	p1.Url = pageUrl
	p1.Ext = ".jpg"

	gonhentai.UseProgressBar = true

	err := p1.Save("./img.jpg", 0644)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Image_Save [OK]")
}
