package nhentai_test

import (
	"encoding/json"
	"github.com/KiritoNya/nhentai"
	"os"
	"strconv"
	"testing"
)

const dojinshiId int = 354862

func TestNewDoujinshiId(t *testing.T) {
	doujin, err := nhentai.NewDoujinshiId(dojinshiId)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Doujinshi: ", doujin)
	t.Log("NewDoujinshiId [OK]")
}

func TestNewDoujinshiUrl(t *testing.T) {

	doujinUrl := nhentai.DoujinPrefix + strconv.Itoa(dojinshiId)

	doujin, err := nhentai.NewDoujinshiUrl(doujinUrl)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Doujinshi: ", doujin)
	t.Log("NewDoujinshiUrl [OK]")
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

	t.Log("DoujinshiUnmarshalJSON: [OK]")
}
