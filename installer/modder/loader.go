package modder

import (
	"errors"
	"os"

	"github.com/pilowl/lethalpacker/pkg/file"
	"github.com/pilowl/lethalpacker/pkg/logger"
)

type ModLoader struct {
	files []file.File
}

func NewLoaderFromZip(zipBytes []byte) (*ModLoader, error) {
	files, err := file.UnzipFiles(zipBytes)
	if err != nil {
		return nil, err
	}

	return &ModLoader{
		files: files,
	}, err
}

func (ml *ModLoader) GetMods() []Mod {
	mods := make([]Mod, 0)
	for _, file := range ml.files {
		if !isPluginsFolder(file) {
			continue
		}

		if file.GetExt() != ".dll" {
			continue
		}

		if definedMod, ok := DLLMappings[file.GetFileName()]; ok {
			definedMod.RelPath = file.RelPath
			mods = append(mods, definedMod)
			continue
		}

		mods = append(mods, Mod{
			Name:       file.GetFileName(),
			RelPath:    file.RelPath,
			Active:     false,
			ClientSide: false,
		})
	}
	return mods
}

const (
	FOLDER_BEPINEX        = `BepInEx`
	FOLDER_BEPINEX_CONFIG = `BepInEx/config/`
)

func (ml *ModLoader) InstallMods(mods []Mod, lethalCompanyPath string) error {
	log := logger.NewLogger()

	lethalCompanyPath = file.UnifySlashes(lethalCompanyPath)
	if exists := file.Exists(lethalCompanyPath); !exists {
		return errors.New("Failed to verify Lethal Company installation existance")
	}

	excludedFiles := make(map[string]interface{})

	for _, mod := range mods {
		if !mod.Active {
			excludedFiles[mod.RelPath] = struct{}{}
			for _, relatedFile := range mod.ChildFilePaths {
				excludedFiles[relatedFile] = struct{}{}
			}
		}
	}

	var (
		bepinexAbsolutePath = lethalCompanyPath + FOLDER_BEPINEX

		cachedConfigFiles = make([]file.File, 0)

		err error
	)

	if file.Exists(bepinexAbsolutePath) {
		if file.Exists(lethalCompanyPath + FOLDER_BEPINEX_CONFIG) {
			if cachedConfigFiles, err = file.GetAllFilesInFolder(lethalCompanyPath, FOLDER_BEPINEX_CONFIG); err != nil {
				log.With("bepinex_config_path", lethalCompanyPath+FOLDER_BEPINEX_CONFIG).WithError(err)
				return errors.New("Failed to backup config files in BepinEx/Config foder")
			}
		}
		if err := os.RemoveAll(bepinexAbsolutePath); err != nil {
			log.With("bepinex_path", bepinexAbsolutePath).WithError(err).Error("Faield to delete bepinex folder")
			return errors.New("Failed to remove existing BepInEx folder. Maybe it is open right now or something?")
		}
	}

	for _, f := range ml.files {
		if _, excluded := excludedFiles[f.RelPath]; excluded {
			continue
		}
		if err := file.Create(lethalCompanyPath+f.GetParentFolders(), f.GetFileName(), f.Data); err != nil {
			log.With("path", lethalCompanyPath+f.GetParentFolders()).
				With("file_name", f.GetFileName()).
				WithError(err).
				Error("Failed to create file")
		}
	}

	for _, f := range cachedConfigFiles {
		if err := file.Create(lethalCompanyPath+f.GetParentFolders(), f.GetFileName(), f.Data); err != nil {
			log.With("path", lethalCompanyPath+f.GetParentFolders()).
				With("file_name", f.GetFileName()).
				WithError(err).
				Error("Failed to create ffiel")
		}
	}

	return nil
}

func isPluginsFolder(file file.File) bool {
	return file.GetParentFolder() == "plugins"
}
