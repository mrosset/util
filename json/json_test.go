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
}

var (
	foo = struct {
	}{}
	testStruct = &Test{
		Name:        "plan",
		Loooooooong: "Long",
		Short:       "Short",
		Version:     "1.0",
		UrlExpect:   "http://ftp.gnu.org/gnu/plan-1.0.tar.gz",
		Url:         "http://ftp.gnu.org/gnu/{{.Name}}-{{.Version}}.tar.gz",
	}
	testFile = "testdata/test.json"
)

func TestWrite(t *testing.T) {
	err := Write(testStruct, testFile)
	if err != nil {
		t.Fatal(err)
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
