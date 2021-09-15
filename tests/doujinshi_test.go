package gonhentai_test

import (
	"encoding/json"
	"github.com/KiritoNya/gonhentai"
	"os"
	"strconv"
	"testing"
)

func TestNewDoujinshiId(t *testing.T) {
	doujin, err := gonhentai.NewDoujinshiId(doujinshiId)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Doujinshi: ", doujin)
	t.Log("NewDoujinshiId [OK]")
}

func TestNewDoujinshiUrl(t *testing.T) {

	doujinUrl := gonhentai.DoujinBaseUrl + strconv.Itoa(doujinshiId)

	doujin, err := gonhentai.NewDoujinshiUrl(doujinUrl)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Doujinshi: ", doujin)
	t.Log("NewDoujinshiUrl [OK]")
}

func TestDoujinshi_GetUrl(t *testing.T) {
	var d gonhentai.Doujinshi
	d.Id = doujinshiId

	err := d.GetUrl()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("URL:", d.Url)
	t.Log("Doujinshi_GetUrl [OK]")
}

func TestDoujinshi_GetRelated(t *testing.T) {

	// Make doujinshi object
	d, err := gonhentai.NewDoujinshiId(doujinshiId)
	if err != nil {
		t.Fatal(err)
	}

	// Get related doujinshi
	err = d.GetRelated()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Related Doujinshi:", d.Related)
	t.Log("Doujinshi_GetRelated [OK]")
}

func TestDoujinshi_GetComments(t *testing.T) {
	// Make doujinshi object
	d, err := gonhentai.NewDoujinshiId(doujinshiId)
	if err != nil {
		t.Fatal(err)
	}

	// Get related doujinshi comments
	err = d.GetComments()
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Comments:", d.Comments)
	t.Log("Doujinshi_GetComments [OK]")
}

func TestDoujinshi_Save(t *testing.T) {
	// Make doujinshi object
	d, err := gonhentai.NewDoujinshiId(370060)
	if err != nil {
		t.Fatal(err)
	}

	// Save doujinshi
	err = d.Save(pathTemplate, 0644)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Doujinshi_Save [OK]")
}

func TestDoujinshi_UnmarshalJSON(t *testing.T) {
	var d gonhentai.Doujinshi

	content, err := os.ReadFile("./doujinshi.test.json")
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(content, &d)
	if err != nil {
		t.Fatal(err)
	}

	data, err := json.MarshalIndent(d, " ", "\t")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(data))
	t.Log("DoujinshiUnmarshalJSON: [OK]")
}
