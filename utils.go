package nhentai

import (
	"bytes"
	"errors"
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
