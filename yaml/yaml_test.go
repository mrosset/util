package yaml

import (
	"github.com/mrosset/util/json"
	"github.com/mrosset/via/pkg"
	"log"
	"os"
	"testing"
)

//const testfile = "/home/strings/gocode/src/github.com/mrosset/via/plans/config.json"
const testfile = "/home/strings/gocode/src/github.com/mrosset/via/plans/core/gcc.json"

var (
	p = new(via.Plan)
)

func init() {
	err := json.Read(p, testfile)
	if err != nil {
		log.Fatal(err)
	}
}

func testWrite(t *testing.T) {
	err := Write(p, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWritePretty(t *testing.T) {
	err := WritePretty(p, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}
