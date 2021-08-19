package nhentai

type Image struct {
	title  string `validate:"omitempty"`
	url    string `validate:"omitempty,nhentai_img_url"`
	size   int64  `validate:"omitempty,min=0"`
	heigth int    `validate:"omitempty,min=0"`
	width  int    `validate:"omitempty,min=0"`
	ext    string `validate:"omitempty,eq=jpg,eq=png,eq=gif"`
	data   []byte
}
