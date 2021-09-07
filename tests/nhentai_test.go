package tests_test

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
