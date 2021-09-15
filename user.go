package gonhentai

import (
	"errors"
	"golang.org/x/net/html"
	"strconv"
	"strings"
)

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

// NewUser is a function that creates a new User object
func NewUser(url string) (*User, error) {
	var u User
	u.Url = url

	// Get html of user page
	html, err := u.getHtml()
	if err != nil {
		return nil, err
	}

	u.html = html
	return &u, nil
}

// GetId is a function that extract Id from the Url and assign it to the User object
func (u *User) GetId() (err error) {
	// Check url
	if !validateUserUrl(u.Url) {
		return errors.New("Url not setted")
	}

	idString := strings.Split(u.Url, "/")[4]
	idInt, err := strconv.Atoi(idString)
	if err != nil {
		return err
	}

	u.Id = idInt
	return nil
}

// getHtml is a function that do the request and add html to the object
func (u *User) getHtml() (*html.Node, error) {

	// Check url
	if u.Url == "" {
		return nil, errors.New("Url not setted")
	}

	// Do request
	resp, err := ClientHttp.Get(u.Url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse html response
	html, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	return html, nil
}
