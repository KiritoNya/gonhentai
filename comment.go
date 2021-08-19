package nhentai

import "time"

type Comment struct {
	id        int `validate:"min=0"`
	galleryId int `validate:"min=0"`
	poster    *User
	postDate  time.Time
	body      string
}
