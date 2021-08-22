package nhentai

import "golang.org/x/net/html"

type User struct {
	id           int `validate:"min=0"`
	username     string
	url          string
	profileImage *Avatar
	isSuperUser  bool
	isStaff      bool
	raw          string `validate:"json"`
	html         *html.Node
}
