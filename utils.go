package gonhnetai

import (
	"bytes"
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
