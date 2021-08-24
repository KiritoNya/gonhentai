package nhentai

import (
	"bytes"
	"errors"
	"strconv"
	"text/template"
)

// templateSolver is a util function that resolve a template
func templateSolver(tmpl string, values interface{}) (string, error) {
	// Create template
	t := template.Must(template.New("template").Parse(tmpl))

	// Execute template
	buf := new(bytes.Buffer)

	if err := t.Execute(buf, values); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// j => jpg, p => png, g => gif
func normalizeExt(ext string) (string, error) {
	switch ext {
	case "j":
		return "jpg", nil
	case "p":
		return "png", nil
	case "g":
		return "gif", nil
	default:
		return "", errors.New("Image type not found")
	}
}

// normalizePageName is a function that generates the name of the images based on the total number of pages of the doujinshi.
//pagTot: 50 => imgName: 01.jpg | pagTot: 9 => imgName: 1.jpg
func normalizePageName(numPage, pageTot int) string {
	var pageResult string

	if pageTot > 9 {
		pageResult += "0"
	}

	if pageTot > 99 {
		pageResult += "0"
	}

	if pageTot > 999 {
		pageResult += "0"
	}

	return pageResult + strconv.Itoa(numPage)
}
