package gonhnetai

import (
	"encoding/json"
	"time"
)

// Comment is a struct that contains all the information of a comment
type Comment struct {
	Id        int `validate:"min=0"`
	GalleryId int `validate:"min=0"`
	Poster    *User
	PostDate  time.Time
	Body      string
}

// UnmarshalJSON is a json parse for the comment object
func (c *Comment) UnmarshalJSON(b []byte) error {
	var rawComment map[string]json.RawMessage

	// Parse comment json
	err := json.Unmarshal(b, &rawComment)
	if err != nil {
		return err
	}

	// Get id of comment
	err = json.Unmarshal(rawComment["id"], &c.Id)
	if err != nil {

		return err
	}

	// Get gallery id
	err = json.Unmarshal(rawComment["gallery_id"], &c.GalleryId)
	if err != nil {
		return err
	}

	// Get poster
	err = json.Unmarshal(rawComment["poster"], &c.Poster)
	if err != nil {
		return err
	}

	// Get post date
	var date int64
	err = json.Unmarshal(rawComment["post_date"], &date)
	if err != nil {
		return err
	}
	c.PostDate = time.Unix(date, 0)

	// Get body
	err = json.Unmarshal(rawComment["body"], &c.Body)
	if err != nil {
		return err
	}

	return nil
}
