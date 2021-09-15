package gonhentai_test

import "github.com/KiritoNya/gonhentai"

// Test input const
const (
	numPage               int    = 36
	numPageIncorrect      int    = -1
	mediaId               int    = 1886630
	mediaIdIncorrect      int    = 1255555555
	doujinshiId           int    = 354862
	dojinshiIdIncorrect   int    = 1255447877
	pageUrl               string = "https://i.nhentai.net/galleries/1886630/36.jpg"
	pageUrlIncorrect      string = "https://i.nhentai.net/galleries/1886630/360.jpg"
	pathTemplate          string = "/home/<username>/{{.Doujinshi.Id}} - {{.Doujinshi.Title.Pretty}}/{{.Page.Num}}.{{.Page.Ext}}"
	pathTemplateIncorrect string = "/home/<username>/{{.Doujinshi.Fake}} - {{.Doujinshi.Title.Pretty}}/{{.Page.Num}}.{{.Page.Ext}}"
	imageName             string = "img.jpg"
)

var InputTests = struct {
	SearchQuery string
	SearchTag   int
	QueryFilter gonhentai.QueryFilter
	QueryOption gonhentai.QueryOptions
}{
	SearchQuery: "Blend s",
	SearchTag:   29859,
	QueryFilter: gonhentai.QueryFilter{
		ToDelete: []gonhentai.Filter{
			{
				Id:   0,
				Name: "yaoi",
				Type: gonhentai.Tag,
			},
		},
		ToFilter: []gonhentai.Filter{
			{
				Id:   0,
				Name: "maika sakuranomiya",
				Type: gonhentai.Character,
			},
		},
	},
	QueryOption: gonhentai.QueryOptions{Page: "1"},
}

var OutputTest = struct {
	SearchCustom int
}{
	SearchCustom: 20,
}
