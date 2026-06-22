package web

import (
	"io/fs"
	"testing"
)

func TestWrapDistFSOpenTrimsLeadingSlash(t *testing.T) {
	assets, err := fsSubDirEntries(distFS, "dist/assets")
	if err != nil {
		t.Fatalf("read embedded assets: %v", err)
	}
	if len(assets) == 0 {
		t.Fatal("embedded assets directory is empty")
	}

	name := assets[0]
	wrapped := &wrapDistFS{FS: distFS}
	file, err := wrapped.Open("/" + name)
	if err != nil {
		t.Fatalf("open asset with leading slash: %v", err)
	}
	_ = file.Close()
}

func fsSubDirEntries(fsys interface {
	ReadDir(string) ([]fs.DirEntry, error)
}, dir string) ([]string, error) {
	entries, err := fsys.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	names := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			names = append(names, entry.Name())
		}
	}
	return names, nil
}
