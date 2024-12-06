package file

import (
	"archive/zip"
	"bytes"
	"io"
)

func UnzipFiles(zipBytes []byte) ([]File, error) {
	bReader := bytes.NewReader(zipBytes)
	zipReader, err := zip.NewReader(bReader, int64(len(zipBytes)))
	if err != nil {
		return nil, err
	}

	files := make([]File, 0, len(zipReader.File))
	for _, file := range zipReader.File {
		archiveFile, err := file.Open()
		if err != nil {
			return nil, err
		}
		defer archiveFile.Close()

		var newFile File

		newFile.Data, err = io.ReadAll(archiveFile)
		if err != nil {
			return nil, err
		}
		newFile.RelPath = file.Name

		files = append(files, newFile)
	}

	return files, nil
}
