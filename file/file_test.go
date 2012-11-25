package file

import (
	"os"
	"path"
	"testing"
)

var (
	existsFiles   = []string{"../util.go", "file.go"}
	notExistFiles = []string{"aaaaaaaa", "bbbbbbbbb"}
	testFile      = "testdata/pass.md5"
)

func TestExists(t *testing.T) {
	for _, f := range existsFiles {
		exists := Exists(f)
		if !exists {
			t.Errorf("expect to find %s got %v", f, exists)
		}
		t.Logf("%s -> %v", f, exists)
	}
	for _, f := range notExistFiles {
		exists := Exists(f)
		if exists {
			t.Errorf("expect not to find %s got %v", f, exists)
		}
		t.Logf("%s -> %v", f, exists)
	}
}

func TestExpand(t *testing.T) {
	expect := os.Getenv("HOME")
	if got := Path("$HOME").Expand(); got != expect {
		t.Errorf("expected %s got %s", expect, got)
	}
}

func TestJoin(t *testing.T) {
	expect := path.Join(os.Getenv("HOME"), "test")
	if got := Path("$HOME").Add("test"); got != expect {
		t.Errorf("expected %s got %s", expect, got)
	}
}

func TestCat(t *testing.T) {
	err := Cat(os.Stdout, "file.go")
	if err != nil {
		t.Error(err)
	}
}

func TestMd5(t *testing.T) {
	var (
		expect = "4528E6A7BB9341C36C425FAF40EF32C3"
	)
	got, err := Md5(testFile)
	if err != nil {
		t.Error(err)
	}
	if expect != got {
		t.Errorf("expected %s got %s", expect, got)
	}
}

func TestMd5Path(t *testing.T) {
	var (
		expect = "8EE1066C618D504704A7C188C873BBBB"
		got    = Md5Path(testFile)
	)
	if expect != got {
		t.Errorf("expected %s got %s", expect, got)
	}
}
