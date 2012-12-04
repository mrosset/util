package walk

// #include "fts.h"
// #include "stdlib.h"
/*
FTS *new_fts(char *root) {
	char *fts_args[] = { root, NULL };
	return fts_open(fts_args, FTS_FOLLOW | FTS_PHYSICAL | FTS_NOSTAT | FTS_XDEV, NULL);
}
*/
import "C"
import (
	"os"
	"path/filepath"
)

func Walk(root string) (files []string, err error) {
	walkFn := func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	}
	err = filepath.Walk(root, walkFn)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func WalkC(root string) (files []string, err error) {
	fts := C.new_fts(C.CString(root))
	defer C.fts_close(fts)
	for ent := C.fts_read(fts); ent != nil; ent = C.fts_read(fts) {
		files = append(files, C.GoString(ent.fts_path))
	}
	return
}
