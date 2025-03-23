package findfiles_test

import (
	"os"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"

	"findfiles"
)

func TestFindFiles_checkTxtFileInTestdataFolder(t *testing.T) {
	t.Parallel()

	fsys := os.DirFS("testdata")
	want := []string{
		"file.txt",
		"subfolder/subfolder.txt",
		"subfolder2/another.txt",
		"subfolder2/file.txt",
	}
	got := findfiles.FindFiles(fsys, "txt")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFindFiles_checkGoFileInMemory(t *testing.T) {
	t.Parallel()

	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}

	got := findfiles.FindFiles(fsys, "go")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}
