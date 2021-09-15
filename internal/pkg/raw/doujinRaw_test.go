package raw_test

import (
	"encoding/json"
	"github.com/KiritoNya/gonhentai/internal/pkg/raw"
	"testing"
)

func TestDoujinRaw_Id(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	id, err := dr.Id()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Doujin.Id != id {
		t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.Doujin.Id, id)
	}
}

func TestDoujinRaw_MediaId(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	mediaId, err := dr.MediaId()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Doujin.MediaId != mediaId {
		t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.Doujin.MediaId, mediaId)
	}
}

func TestDoujinRaw_Title(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	titles, err := dr.Title()
	if err != nil {
		t.Fatal(err)
	}

	var titleEnglish, titleJapanese, titlePretty string
	titlesMap := make(map[string]string)
	err = json.Unmarshal(titles.Data["english"], &titleEnglish)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(titles.Data["japanese"], &titleJapanese)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(titles.Data["pretty"], &titlePretty)
	if err != nil {
		t.Fatal(err)
	}
	titlesMap["English"] = titleEnglish
	titlesMap["Japanese"] = titleJapanese
	titlesMap["Pretty"] = titlePretty

	// Check map length
	if len(titlesMap) != len(OutputTest.Doujin.Titles) {
		t.Fatalf("\nExpected: '%d %s'\nObtained: '%d %s'", len(OutputTest.Doujin.Titles), "title", len(titlesMap), "title")
	}

	for key, title := range titlesMap {
		if OutputTest.Doujin.Titles[key] != title {
			t.Errorf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Doujin.Titles[key], title)
		}
	}
}

func TestDoujinRaw_CoverImage(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	cover, err := dr.CoverImage()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.CoverImage, cover.Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_Thumbnail(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	thumbnail, err := dr.Thumbnail()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.Thumbnail, thumbnail.Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_Pages(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	pages, err := dr.Pages()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.Pages, pages[0].Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_Scanlator(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	scanlator, err := dr.Scanlator()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Doujin.Scanlator != scanlator {
		t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Doujin.Scanlator, scanlator)
	}
}

func TestDoujinRaw_UploadDate(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	uploadDate, err := dr.UploadDate()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Doujin.UploadDate != uploadDate {
		t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Doujin.UploadDate, uploadDate)
	}
}

func TestDoujinRaw_Parodies(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	parodies, err := dr.Parodies()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.Parodies, parodies[0].Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_Characters(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	characters, err := dr.Characters()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.Characters, characters[0].Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_Tags(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	tags, err := dr.Tags()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.Tags, tags[0].Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_Artists(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	artists, err := dr.Artists()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.Artists, artists[0].Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_Groups(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	groups, err := dr.Groups()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.Groups, groups[0].Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_Languages(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	languages, err := dr.Languages()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.Languages, languages[0].Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_Categories(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	categories, err := dr.Categories()
	if err != nil {
		t.Fatal(err)
	}

	isEqual, msg, err := checkResult(OutputTest.Doujin.Categories, categories[0].Data)
	if err != nil {
		t.Fatal(err)
	}

	if !isEqual {
		t.Fatal(msg)
	}
}

func TestDoujinRaw_NumPages(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	numPages, err := dr.NumPages()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Doujin.NumPages != numPages {
		t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.Doujin.NumPages, numPages)
	}
}

func TestDoujinRaw_NumFavorites(t *testing.T) {
	dr, err := raw.NewDoujinRawUrl(ClientHttp, InputTest.Doujin.Url)
	if err != nil {
		t.Fatal(err)
	}

	numFavorites, err := dr.NumFavorites()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Doujin.NumFavorites != numFavorites {
		t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.Doujin.NumFavorites, numFavorites)
	}
}
