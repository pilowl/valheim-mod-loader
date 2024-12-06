package main

import (
	"archive/zip"
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/pilowl/lethalpacker/pkg/file"
	"github.com/pilowl/lethalpacker/pkg/logger"
)

const ZIP_ARCHIVE_NAME = "valgaym_mod_pack.zip"
const ZIP_ARCHIVE_PATH = "installer/pack/"

const MODS_PATH = `mods\`

func main() {
	log := logger.NewLogger()

	files := make([]file.File, 0)

	err := filepath.Walk(MODS_PATH,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if info.IsDir() {
				return nil
			}

			var newFile file.File

			fin, err := os.Open(path)
			if err != nil {
				return err
			}
			defer fin.Close()

			if newFile.Data, err = io.ReadAll(fin); err != nil {
				return err
			}

			newFile.RelPath = file.UnifySlashes(strings.TrimLeft(path, MODS_PATH))
			files = append(files, newFile)
			return nil
		})
	if err != nil {
		log.WithError(err).Error("Failed to walk through the 'mod' folder")
		return
	}

	fzip, err := os.Create(ZIP_ARCHIVE_PATH + ZIP_ARCHIVE_NAME)
	if err != nil {
		log.WithError(err).Error("Failed to create archive")
		return
	}
	defer func() {
		if err := fzip.Close(); err != nil {
			log.WithError(err).Error("Failed to close zip file")
			return
		}
	}()

	zipWriter := zip.NewWriter(fzip)
	defer func() {
		if err := zipWriter.Flush(); err != nil {
			log.WithError(err).Error("Failed to flush zip writer")
			return
		}
		if err := zipWriter.Close(); err != nil {
			log.WithError(err).Error("Failed to close zip writer")
			return
		}
	}()

	for _, file := range files {
		writer, err := zipWriter.Create(file.RelPath)
		if err != nil {
			log.WithError(err).With("path", file.RelPath).Error("Failed to create file")
			return
		}
		if _, err := io.Copy(writer, bytes.NewReader(file.Data)); err != nil {
			log.WithError(err).Error("Failed to create file")
			return
		}
	}

	log.Info("Zip created at " + ZIP_ARCHIVE_PATH + ZIP_ARCHIVE_NAME)
}
