package nhentai

type Image struct {
	title  string `validate:"omitempty" json:",omitempty"`
	url    string `validate:"omitempty,nhentai_img_url" json:",omitempty"`
	size   int64  `validate:"omitempty,min=0" json:",omitempty"`
	heigth int    `validate:"omitempty,min=0" json:"h,omitempty"`
	width  int    `validate:"omitempty,min=0" json:"w,omitempty"`
	ext    string `validate:"omitempty,eq=jpg,eq=png,eq=gif" json:"t,omitempty"`
	data   []byte
}
