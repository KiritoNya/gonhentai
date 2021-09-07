package raw_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//const UrlTest string = "https://nhentai.net/api/gallery/354862"
const UrlTest string = "https://nhentai.net/api/gallery/371687"

var ClientHttp = http.DefaultClient

/* __________________________________________________________________________ */

/* ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ */
/* SECTION                          Test Inputs                                 */
/* ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ */

var InputTest = struct {
	Doujin struct {
		Url string
	}
	Image string
	Tag   string
	Title string
}{
	Doujin: struct {
		Url string
	}{
		Url: "https://nhentai.net/api/gallery/371687",
	},
	Image: `{
			"t":"j",
            "w":1280,
            "h":1863
         }`,
	Tag: `{
         "id": 9162,
         "type": "tag",
         "name": "masturbation",
         "url": "/tag/masturbation/",
         "count": 11535
      }`,
	Title: `{
      "english":"[Kabuki Shigeyuki] Fetish Girl [English] [QuarantineScans]",
      "japanese":"[\u9999\u5439\u8302\u4e4b] \u30d5\u30a7\u30c6\u30a3\u30c3\u30b7\u30e5\u30fb\u30ac\u30fc\u30eb [\u82f1\u8a33]",
      "pretty":"Fetish Girl"
   }`,
}

/* __________________________________________________________________________ */

/* ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ */
/* SECTION                          Test Outputs                                 */
/* ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ */

var OutputTest = struct {
	Doujin struct {
		Id           int
		MediaId      int
		Titles       map[string]string
		CoverImage   string
		Thumbnail    string
		Pages        string
		Scanlator    string
		UploadDate   time.Time
		Parodies     string
		Characters   string
		Tags         string
		Artists      string
		Groups       string
		Languages    string
		Categories   string
		NumPages     int
		NumFavorites int
	}
	Image struct {
		Ext    string
		Width  int
		Height int
	}
	Tags struct {
		Id    int
		Type  string
		Name  string
		Url   string
		Count int
	}
	Title struct {
		English  string
		Japanese string
		Pretty   string
	}
}{
	Doujin: struct {
		Id           int
		MediaId      int
		Titles       map[string]string
		CoverImage   string
		Thumbnail    string
		Pages        string
		Scanlator    string
		UploadDate   time.Time
		Parodies     string
		Characters   string
		Tags         string
		Artists      string
		Groups       string
		Languages    string
		Categories   string
		NumPages     int
		NumFavorites int
	}{
		Id:      371687,
		MediaId: 2000657,
		Titles: map[string]string{
			"English":  "(Mega Akihabara Doujinsai 3) [Eclipse (Rougetu)] Prisma Sanshimai to Chaldea Ikaseya Oji-san (Fate/Grand Order, Fate/kaleid liner Prisma Illya)",
			"Japanese": "(メガ秋葉原同人祭 第3回) [えくりぷす (朧月)] プリズマ三姉妹とカルデアイかせ屋おじさん (Fate/Grand Order、Fate/kaleid liner プリズマ☆イリヤ)",
			"Pretty":   "Prisma Sanshimai to Chaldea Ikaseya Oji-san",
		},
		CoverImage: `{"t": "p","w": 350,"h": 495}`,
		Thumbnail:  `{"t": "p","w": 250,"h": 354}`,
		Pages:      `{"t":"p","w":1280,"h":1810}`,
		Scanlator:  "",
		UploadDate: time.Unix(1630688695, 0),
		Parodies: `{
				"id":24886,
				"type":"parody",
				"name":"fate kaleid liner prisma illya",
				"url":"/parody/fate-kaleid-liner-prisma-illya/",
				"count":639
		}`,
		Characters: `{
				"id":15651,
				"type":"character",
				"name":"miyu edelfelt",
				"url":"/character/miyu-edelfelt/",
				"count":253
		}`,
		Tags: `{
				"id":6817,
				"type":"tag",
				"name":"unusual pupils",
				"url":"/tag/unusual-pupils/",
				"count":9597
		}`,
		Artists: `{
				"id":12117,
				"type":"artist",
				"name":"rougetu",
				"url":"/artist/rougetu/",
				"count":69
		}`,
		Groups: `{
			"id":21696,
			"type":"group",
			"name":"eclipse",
			"url":"/group/eclipse/",
			"count":61
		}`,
		Languages: `{
			"id":6346,
			"type":"language",
			"name":"japanese",
			"url":"/language/japanese/",
			"count":215493
		}`,
		Categories: `{
			"id":    33172,
			"type":  "category",
			"name":  "doujinshi",
			"url":   "/category/doujinshi/",
			"count": 271193
		}`,
		NumPages:     34,
		NumFavorites: 0,
	},
	Image: struct {
		Ext    string
		Width  int
		Height int
	}{
		"jpg",
		1280,
		1863,
	},
	Tags: struct {
		Id    int
		Type  string
		Name  string
		Url   string
		Count int
	}{
		9162,
		"tag",
		"masturbation",
		"/tag/masturbation/",
		11535,
	},
	Title: struct {
		English  string
		Japanese string
		Pretty   string
	}{
		"[Kabuki Shigeyuki] Fetish Girl [English] [QuarantineScans]",
		"[香吹茂之] フェティッシュ・ガール [英訳]",
		"Fetish Girl",
	},
}

// checkResult is a function that checks whether the expected result and the obtained result are equal
func checkResult(expected string, obtained map[string]json.RawMessage) (result bool, msg string, err error) {
	var testMap map[string]interface{}
	var testMap2 map[string]interface{}

	err = json.Unmarshal([]byte(expected), &testMap)
	if err != nil {
		return false, "", err
	}

	data, err := json.Marshal(obtained)
	if err != nil {
		return false, "", err
	}

	err = json.Unmarshal(data, &testMap2)
	if err != nil {
		return false, "", err
	}

	// Check lenght
	if len(testMap) != len(testMap2) {
		return false, fmt.Sprintf("\nExpected: '%d' attributes\nObtained: '%d' attributes", len(testMap), len(testMap2)), nil
	}

	// Check values
	for key, mapAttr := range testMap {
		if testMap2[key] != mapAttr {
			return false, fmt.Sprintf("\nExpected: '%v'\nObtained: '%v'", mapAttr, testMap2[key]), nil
		}
	}

	return true, "", nil
}
