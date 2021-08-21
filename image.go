package nhentai

type Image struct {
	Title  string `validate:"omitempty" json:",omitempty"`
	Url    string `validate:"omitempty,nhentai_img_url" json:",omitempty"`
	Size   int64  `validate:"omitempty,min=0" json:",omitempty"`
	Heigth int    `validate:"omitempty,min=0" json:"h,omitempty"`
	Width  int    `validate:"omitempty,min=0" json:"w,omitempty"`
	Ext    string `validate:"omitempty,eq=jpg,eq=png,eq=gif" json:"t,omitempty"`
	Data   []byte
}
