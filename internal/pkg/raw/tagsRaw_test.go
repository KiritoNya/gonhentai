package raw_test

import (
	"encoding/json"
	"github.com/KiritoNya/gonhentai/internal/pkg/raw"
	"testing"
)

func TestTagRaw_Id(t *testing.T) {
	var tr raw.TagRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Tag), &tr.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	id, err := tr.Id()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Tags.Id != id {
		t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Id, id)
	}
}

func TestTagRaw_Type(t *testing.T) {
	var tr raw.TagRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Tag), &tr.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	typ, err := tr.Type()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Tags.Type != typ {
		t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Type, typ)
	}
}

func TestTagRaw_Name(t *testing.T) {
	var tr raw.TagRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Tag), &tr.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	name, err := tr.Name()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Tags.Name != name {
		t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Name, name)
	}
}

func TestTagRaw_Url(t *testing.T) {
	var tr raw.TagRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Tag), &tr.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	url, err := tr.Url()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Tags.Url != url {
		t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Url, url)
	}
}

func TestTagRaw_Count(t *testing.T) {
	var tr raw.TagRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Tag), &tr.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	count, err := tr.Count()
	if err != nil {
		t.Fatal(err)
	}

	if OutputTest.Tags.Count != count {
		t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Count, count)
	}
}

func TestTagRaw_All(t *testing.T) {
	var tr raw.TagRaw

	// Unmarshal
	err := json.Unmarshal([]byte(InputTest.Tag), &tr.Data)
	if err != nil {
		t.Fatal(err)
	}

	// Call method to test
	tagMap, err := tr.All()
	if err != nil {
		t.Fatal(err)
	}

	for attr, value := range tagMap {
		switch attr {
		case "Id":
			if value != OutputTest.Tags.Id {
				t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Id, value)
			}
		case "Type":
			if value != OutputTest.Tags.Type {
				t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Type, value)
			}
		case "Name":
			if value != OutputTest.Tags.Name {
				t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Name, value)
			}
		case "Url":
			if value != OutputTest.Tags.Url {
				t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Url, value)
			}
		case "Count":
			if value != OutputTest.Tags.Count {
				t.Errorf("\nExpected: '%v'\nObtained: '%v'", OutputTest.Tags.Count, value)
			}
		}
	}
}
