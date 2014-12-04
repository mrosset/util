package json

import (
	"bytes"
	"encoding/json"
	"text/template"
)

type Template interface {
	SetTemplate(interface{})
}

func Execute(t Template) error {
	in := new(bytes.Buffer)
	// encode interface to json
	t.SetTemplate(t)
	err := json.NewEncoder(in).Encode(t)
	if err != nil {
		return err
	}
	tmpl, err := template.New("").Parse(in.String())
	if err != nil {
		return err
	}

	out := new(bytes.Buffer)
	// execute template to out buffer
	err = tmpl.Execute(out, t)
	if err != nil {
		return err
	}
	// finally decode the executed template
	return json.NewDecoder(out).Decode(t)
}

// Reads a json file and parses fields
func ReadTemplate(t Template, path string) error {
	err := Read(t, path)
	if err != nil {
		return err
	}
	return Execute(t)
}
