package file

import (
	"io/fs"
	"path/filepath"
	"strings"
)

func GetAllFilesInFolder(rootPath string, filePath string) ([]File, error) {
	files := make([]File, 0)
	err := filepath.WalkDir(rootPath+filePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !d.Type().IsRegular() {
			return nil
		}

		fileBytes, err := Read(path)
		if err != nil {
			return err
		}

		path = UnifySlashes(path)

		files = append(files, File{
			RelPath: strings.TrimLeft(path, rootPath),
			Data:    fileBytes,
		})

		return nil
	})
	if err != nil {
		return nil, err
	}

	return files, nil
}
