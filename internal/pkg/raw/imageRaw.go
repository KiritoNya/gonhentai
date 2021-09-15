package raw

import "encoding/json"

// ImageRaw is the data struct that describes a image raw
type ImageRaw struct {
	Data map[string]json.RawMessage
}

// Ext  is a function that returns the type of the image raw
func (ir *ImageRaw) Ext() (ext string, err error) {
	//Unmarshal data
	err = json.Unmarshal(ir.Data["t"], &ext)
	if err != nil {
		return "", nil
	}

	return normalizeExt(ext)
}

// Width  is a function that returns the width of the image raw
func (ir *ImageRaw) Width() (width int, err error) {
	return ir.size("w")
}

// Height  is a function that returns the height of the image raw
func (ir *ImageRaw) Height() (height int, err error) {
	return ir.size("h")
}

// All  is a function that returns all info of the image
func (ir *ImageRaw) All() (imageMap map[string]interface{}, err error) {
	imageMap = make(map[string]interface{})

	// Get Ext of image
	ext, err := ir.Ext()
	if err != nil {
		return nil, err
	}

	// Get Width of image
	width, err := ir.Width()
	if err != nil {
		return nil, err
	}

	// Get height of image
	height, err := ir.Height()
	if err != nil {
		return nil, err
	}

	// Fill map
	imageMap["Ext"] = ext
	imageMap["Width"] = width
	imageMap["Height"] = height

	return imageMap, nil
}

// size is a function that returns width or height of the image raw
func (ir *ImageRaw) size(dim string) (s int, err error) {
	err = json.Unmarshal(ir.Data[dim], &s)
	if err != nil {
		return 0, nil
	}
	return
}
