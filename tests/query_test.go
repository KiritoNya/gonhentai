package tests_test

import (
	"encoding/json"
	"fmt"
	"github.com/KiritoNya/nhentai"
	"os"
	"testing"
)

func TestRecentDoujinshi(t *testing.T) {

	// Call function
	dc, err := nhentai.RecentDoujinshi(nhentai.QueryOptions{Page: "1"})
	if err != nil {
		t.Fatal(err)
	}

	// Marshal data
	data, err := json.MarshalIndent(dc, " ", "\t")
	if err != nil {
		t.Fatal(err)
	}

	data2, err := os.ReadFile("data/getRecentDoujinshi.test.json")
	if err != nil {
		t.Fatal(err)
	}

	// Read test data file
	if string(data2) != string(data) {
		t.Fatal(`Test data and obtained dat doesn't match`)
	}

	fmt.Println(string(data))
}

func TestSearch(t *testing.T) {
	// Call function
	dc, err := nhentai.Search("ishtar", nhentai.QueryOptions{Page: "1"})
	if err != nil {
		t.Fatal(err)
	}

	// Marshal data
	data, err := json.MarshalIndent(dc, " ", "\t")
	if err != nil {
		t.Fatal(err)
	}

	data2, err := os.ReadFile("data/search.test.json")
	if err != nil {
		t.Fatal(err)
	}

	// Read test data file
	if string(data2) != string(data) {
		t.Fatal(`Test data and obtained dat doesn't match'`)
	}
}

func TestSearchTag(t *testing.T) {
	// Call function
	dc, err := nhentai.SearchTag(25663, nhentai.QueryOptions{Page: "1"})
	if err != nil {
		t.Fatal(err)
	}

	// Marshal data
	data, err := json.MarshalIndent(dc, " ", "\t")
	if err != nil {
		t.Fatal(err)
	}

	data2, err := os.ReadFile("data/search.test.json")
	if err != nil {
		t.Fatal(err)
	}

	// Read test data file
	if string(data2) != string(data) {
		t.Fatal(`Test data and obtained dat doesn't match'`)
	}
}
