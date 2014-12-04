package json

import (
	"testing"
)

type Test struct {
	Name        string
	Loooooooong string
	Short       string
	Version     string
	UrlExpect   string
	Url         string
	Mirror      string
	template    *Test
}

func (t *Test) SetTemplate(i interface{}) error {
	c := *i.(*Test)
	t.template = &c
	return nil
}

var (
	foo = struct {
	}{}
	testStruct = Test{
		Name:        "plan",
		Loooooooong: "Long",
		Short:       "Short",
		Version:     "1.0",
		UrlExpect:   "http://ftp.gnu.org/gnu/plan-1.0.tar.gz",
		Url:         "http://{{.Mirror}}/gnu/{{.Name}}-{{.Version}}.tar.gz",
		Mirror:      "ftp.gnu.org",
	}
	testFile = "testdata/test.json"
)

func TestWrite(t *testing.T) {
	v := testStruct
	err := Write(v, testFile)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTemplate(t *testing.T) {
	v := testStruct
	err := Execute(&v)
	if err != nil {
		t.Error(err)
	}
	got := v.Url
	expect := testStruct.UrlExpect
	if got != expect {
		t.Errorf("expect '%s' got '%s'", expect, got)
	}
	expect = testStruct.Url
	got = v.template.Url
	if expect != got {
		t.Errorf("expect '%s' got '%s'", expect, got)
	}
}

func TestReadTemplate(t *testing.T) {
	var (
		expect = testStruct.UrlExpect
		got    = &Test{}
	)
	err := ReadTemplate(got, testFile)
	if err != nil {
		t.Error(err)
	}
	if expect != got.Url {
		t.Errorf("expect '%s' got '%s'", expect, got.Url)
	}
	expect = testStruct.Url
	if expect != got.template.Url {
		t.Errorf("expect '%s' got '%s'", expect, got.Url)
	}
}

func TestGetFail(t *testing.T) {
	var (
		got = Get(nil, "http://localhost:10000")
	)
	if got == nil {
		t.Logf("expect %v -> got %s", nil, got)
		t.Error(got)
	}
}
