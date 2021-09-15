package raw_test

import (
	"encoding/json"
	"github.com/KiritoNya/gonhentai/internal/pkg/raw"
	"testing"
)

func TestTitleRaw_English(t *testing.T) {
	var tl raw.TitleRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Title), &tl.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	title, err := tl.English()
	if err != nil {
		t.Fatal(err)
	}

	// Check result
	if title != OutputTest.Title.English {
		t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Title.English, title)
	}
}

func TestTitleRaw_Japanese(t *testing.T) {
	var tl raw.TitleRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Title), &tl.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	title, err := tl.Japanese()
	if err != nil {
		t.Fatal(err)
	}

	// Check result
	if title != OutputTest.Title.Japanese {
		t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Title.Japanese, title)
	}
}

func TestTitleRaw_Pretty(t *testing.T) {
	var tl raw.TitleRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Title), &tl.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	title, err := tl.Pretty()
	if err != nil {
		t.Fatal(err)
	}

	// Check result
	if title != OutputTest.Title.Pretty {
		t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Title.Pretty, title)
	}
}

func TestTitleRaw_All(t *testing.T) {
	var tl raw.TitleRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Title), &tl.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	titleMap, err := tl.All()
	if err != nil {
		t.Fatal(err)
	}

	for lang, title := range titleMap {
		switch lang {
		case "English":
			if title != OutputTest.Title.English {
				t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Title.English, title)
			}
		case "Japanese":
			if title != OutputTest.Title.Japanese {
				t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Title.Japanese, title)
			}
		case "Pretty":
			if title != OutputTest.Title.Pretty {
				t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Title.Pretty, title)
			}
		}
	}
}
