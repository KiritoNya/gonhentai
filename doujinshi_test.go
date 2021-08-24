package nhentai_test

import (
	"encoding/json"
	"github.com/KiritoNya/nhentai"
	"os"
	"strconv"
	"testing"
)

func TestNewDoujinshiId(t *testing.T) {
	doujin, err := nhentai.NewDoujinshiId(dojinshiId)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Doujinshi: ", doujin)
	t.Log("NewDoujinshiId [OK]")
}

func TestNewDoujinshiUrl(t *testing.T) {

	doujinUrl := nhentai.DoujinBaseUrl + strconv.Itoa(dojinshiId)

	doujin, err := nhentai.NewDoujinshiUrl(doujinUrl)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Doujinshi: ", doujin)
	t.Log("NewDoujinshiUrl [OK]")
}

func TestDoujinshi_GetRelated(t *testing.T) {

	// Make doujinshi object
	d, err := nhentai.NewDoujinshiId(dojinshiId)
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
	d, err := nhentai.NewDoujinshiId(dojinshiId)
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
	d, err := nhentai.NewDoujinshiId(370060)
	if err != nil {
		t.Fatal(err)
	}

	// Get related doujinshi comments
	err = d.Save("/home/kiritonya/prova.txt", 0644, nhentai.DefaultImageNameTemplate)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Comments:", d.Comments)
	t.Log("Doujinshi_GetComments [OK]")
}

func TestDoujinshi_UnmarshalJSON(t *testing.T) {
	var d nhentai.Doujinshi

	content, err := os.ReadFile("./tests/doujinshi.test.json")
	if err != nil {
		t.Fatal(err)
	}

	err = json.Unmarshal(content, &d)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Doujinshi: ", d)
	t.Log("DoujinshiUnmarshalJSON: [OK]")
}
