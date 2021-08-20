package nhentai_test

import (
	"github.com/KiritoNya/nhentai"
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
