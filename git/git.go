package git

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"os"
	"path/filepath"
)

// Clone remote URL into directory.
func Clone(dir string, uri string) error {
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:               uri,
		Progress:          os.Stdout,
		Depth:             1,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	return err
}

// gitref returns git branch reference
func gitref(branch string) plumbing.ReferenceName {
	return plumbing.ReferenceName(
		fmt.Sprintf("refs/heads/%s", branch),
	)
}

// CloneBranch clone remove URL with branch to directory
func CloneBranch(dir, uri, branch string) error {
	_, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:           uri,
		Progress:      os.Stdout,
		Depth:         1,
		ReferenceName: gitref(branch),
	})
	return err
}

// Branch returns the currently checked out branch for a git directory
// FIXME: this will probably fail with a detached head
func Branch(path string) (string, error) {
	r, err := git.PlainOpen(path)
	if err != nil {
		return "", err
	}
	head, err := r.Head()
	if err != nil {
		return "", err
	}
	return filepath.Base(head.Name().String()), nil
}
