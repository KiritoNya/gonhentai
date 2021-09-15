package gonhentai_test

import (
	"encoding/json"
	"fmt"
	"github.com/KiritoNya/gonhentai"
	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
	"os"
	"testing"
)

func TestRecentDoujinshi(t *testing.T) {

	// Call function
	dc, err := gonhentai.RecentDoujinshi(InputTests.QueryOption)
	if err != nil {
		t.Fatal(err)
	}

	// Marshal data
	data, err := gonhentai.MarshalIndent(dc, " ", "\t")
	if err != nil {
		t.Fatal(err)
	}

	data2, err := os.ReadFile("data/getRecentDoujinshi.test.json")
	if err != nil {
		t.Fatal(err)
	}

	edits := myers.ComputeEdits(span.URIFromPath("a.txt"), string(data2), string(data))
	diff := fmt.Sprint(gotextdiff.ToUnified("Expected", "Obtained", string(data2), edits))

	// Read test data file
	if diff != "" {
		t.Fatal(diff)
	}
}

func TestSearch(t *testing.T) {
	// Call function
	dc, err := gonhentai.Search(InputTests.SearchQuery, nhentai.QueryOptions{Page: "all"})
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
	_, err := gonhentai.SearchTag(InputTests.SearchTag, InputTests.QueryOption)
	if err != nil {
		t.Fatal(err)
	}

	//fmt.Println(dc)
}

func TestSearchCustom(t *testing.T) {
	qr, err := gonhentai.SearchCustom(InputTests.SearchQuery, InputTests.QueryFilter)
	if err != nil {
		t.Fatal(err)
	}

	for _, res := range qr.Result {
		fmt.Println(res.Id)
	}

	if lenResult := len(qr.Result); lenResult != OutputTest.SearchCustom {
		t.Fatalf("\nExpected: '%d'\nObtained: '%d'", OutputTest.SearchCustom, lenResult)
	}
}
