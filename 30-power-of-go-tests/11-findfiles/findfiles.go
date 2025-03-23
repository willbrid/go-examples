package findfiles

import (
	"io/fs"
	"path/filepath"
)

func FindFiles(fsys fs.FS, extName string) (paths []string) {
	var extension string = "." + extName

	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == extension {
			paths = append(paths, path)
		}
		return nil
	})

	return paths
}
