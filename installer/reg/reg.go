package reg

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/windows/registry"
)

const (
	REG_STEAM_PATH             = `SOFTWARE\WOW6432Node\Valve\Steam`
	REG_STEAM_INSTALL_PATH_KEY = "InstallPath"

	REL_LIBRARY_FOLDERS_VDF_PATH = "\\steamapps\\libraryfolders.vdf"

	LETHAL_COMPANY_STEAM_ID            = 1966720
	REL_STEAM_APPS_LETHAL_COMPANY_PATH = "\\steamapps\\common\\Lethal Company\\"
)

func GetLethalCompanyPath() string {
	// Checking registry for steam installation
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, REG_STEAM_PATH, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	installPath, _, err := k.GetStringValue(REG_STEAM_INSTALL_PATH_KEY)
	if err != nil {
		log.Fatal(err)
	}

	// Reading libraryfolder.vdf to search for Lethal Company install path
	f, err := os.Open(installPath + REL_LIBRARY_FOLDERS_VDF_PATH)
	if err != nil {
		return ""
	}
	defer f.Close()

	var lastPath string

	reader := bufio.NewReader(f)
	for {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}

			log.Fatal(err)
			break
		}

		line := string(bytes)

		// Another steam game install path found, remember it
		if strings.Contains(line, `"path"`) {
			sp := strings.Split(line, `"`)
			if len(sp) < 4 {
				break
			}
			lastPath = sp[3]
		}

		// Lethal Company app id found, return last remembered path
		if strings.Contains(line, fmt.Sprintf(`"%d"`, LETHAL_COMPANY_STEAM_ID)) {
			lethalCompanyPath := fmt.Sprintf(`%s\steamapps\common\Lethal Company\`, lastPath)
			lethalCompanyPath = strings.ReplaceAll(lethalCompanyPath, `\\`, `\`)

			return lethalCompanyPath
		}
	}

	return ""
}
