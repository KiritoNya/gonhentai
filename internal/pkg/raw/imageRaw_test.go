package raw_test

import (
	"encoding/json"
	"github.com/KiritoNya/nhentai/internal/pkg/raw"
	"testing"
)

func TestImageRaw_Ext(t *testing.T) {
	var ir raw.ImageRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Image), &ir.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	ext, err := ir.Ext()
	if err != nil {
		t.Fatal(err)
	}

	if ext != OutputTest.Image.Ext {
		t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Image.Ext, ext)
	}
}

func TestImageRaw_Width(t *testing.T) {
	var ir raw.ImageRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Image), &ir.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	width, err := ir.Width()
	if err != nil {
		t.Fatal(err)
	}

	if width != OutputTest.Image.Width {
		t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.Image.Width, width)
	}
}

func TestImageRaw_Height(t *testing.T) {
	var ir raw.ImageRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Image), &ir.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	height, err := ir.Height()
	if err != nil {
		t.Fatal(err)
	}

	if height != OutputTest.Image.Height {
		t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.Image.Height, height)
	}
}

func TestImageRaw_All(t *testing.T) {
	var ir raw.ImageRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Image), &ir.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	imageMap, err := ir.All()
	if err != nil {
		t.Fatal(err)
	}

	for attr, value := range imageMap {
		switch attr {
		case "Ext":
			if value != OutputTest.Image.Ext {
				t.Fatalf("\nExpected: '%s'\nObtained: '%s'", OutputTest.Image.Ext, value)
			}
		case "Width":
			if value != OutputTest.Image.Width {
				t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.Image.Width, value)
			}
		case "Height":
			if value != OutputTest.Image.Height {
				t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.Image.Height, value)
			}
		}
	}

}
