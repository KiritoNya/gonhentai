package gonhnetai

import "golang.org/x/net/html"

// User is a struct that contains all of the user's information
type User struct {
	Id           int `validate:"min=0"`
	Username     string
	Url          string
	ProfileImage *Avatar
	IsSuperUser  bool
	IsStaff      bool
	html         *html.Node
}
