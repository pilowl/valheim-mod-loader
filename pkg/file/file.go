package file

import (
	"path/filepath"
)

type File struct {
	RelPath string
	Data    []byte
}

func (f File) GetFileName() string {
	return filepath.Base(f.RelPath)
}

func (f File) GetParentFolder() string {
	return filepath.Base(filepath.Dir(f.RelPath))
}

func (f File) GetParentFolders() string {
	return UnifySlashes(filepath.Dir(f.RelPath))
}

func (f File) GetExt() string {
	return filepath.Ext(f.RelPath)
}
