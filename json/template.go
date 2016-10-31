package json

import (
	"bytes"
	"encoding/json"
	"text/template"
)

type JsonTemplate interface {
	SetTemplate(interface{})
}

// Encodes a struct into json and then Parses it as a go/template. It then
// decodes it back into a complete struct.
func Parse(out, in interface{}) error {
	var (
		buf = new(bytes.Buffer)
	)

	// encode interface to json
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		return err
	}

	// parse the json as a template
	tmpl, err := template.New("").Parse(buf.String())
	if err != nil {
		return err
	}

	buf.Reset()

	// execute template to out buffer
	err = tmpl.Execute(buf, in)
	if err != nil {
		return err
	}

	// finally decode the executed template
	return json.NewDecoder(buf).Decode(out)
}

func OExecute(t JsonTemplate) error {
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

/*
// Reads a json file and parses fields
func ReadTemplate(t Template, path string) error {
	err := Read(t, path)
	if err != nil {
		return err
	}
	return Execute(t)
}

*/
