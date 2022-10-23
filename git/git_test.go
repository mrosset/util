package git

import (
	"github.com/mrosset/util/file"
	"os"
	"path/filepath"
	"testing"
)

const (
	// TODO: use in memory repository instead?
	GIT_URI = "https://github.com/mrosset/via-test"
)

func TestClone(t *testing.T) {
	t.Parallel()
	var (
		uri  = GIT_URI
		gitd = "testdata/repo"
	)
	defer os.RemoveAll(gitd)
	if err := Clone(gitd, uri); err != nil {
		t.Fatal(err)
	}
	if !file.Exists(filepath.Join(gitd, "README.md")) {
		t.Error("README.md does not exist")
	}
}

func TestCloneBranch(t *testing.T) {
	t.Parallel()
	var (
		uri    = GIT_URI
		gitd   = "testdata/branch"
		expect = "x86_64-via-linux-gnu-release"
	)
	defer os.RemoveAll(gitd)
	if err := CloneBranch(gitd, uri, expect); err != nil {
		t.Fatal(err)
	}
	got, err := Branch(gitd)
	if err != nil {
		t.Fatal(err)
	}
	if got != expect {
		t.Errorf("Expect branch %s got %s", expect, got)
	}
}
