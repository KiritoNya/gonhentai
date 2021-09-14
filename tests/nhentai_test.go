package nhentai_test

import "github.com/KiritoNya/nhentai"

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
	QueryFilter nhentai.QueryFilter
	QueryOption nhentai.QueryOptions
}{
	SearchQuery: "Blend s",
	SearchTag:   0,
	QueryFilter: nhentai.QueryFilter{
		ToDelete: []nhentai.Filter{
			{
				Id:   0,
				Name: "yaoi",
				Type: nhentai.Tag,
			},
		},
		ToFilter: []nhentai.Filter{
			{
				Id:   0,
				Name: "maika sakuranomiya",
				Type: nhentai.Character,
			},
		},
	},
	QueryOption: nhentai.QueryOptions{Page: "1"},
}

var OutputTest = struct {
	SearchCustom int
}{
	SearchCustom: 22,
}
