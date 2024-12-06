package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/sys/windows/registry"
)

func main() {
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\WOW6432Node\Valve\Steam`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()

	installPath, _, err := k.GetStringValue("InstallPath")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Steam installation path is %q\n", installPath)

	f, err := os.Open(fmt.Sprintf("%s\\steamapps\\libraryfolders.vdf", installPath))
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(f)

	var lastPath string

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

		if strings.Contains(line, "\"path\"") {
			sp := strings.Split(line, "\"")
			if len(sp) < 4 {
				break
			}
			lastPath = sp[3]
			fmt.Printf("Found path: %s\n", sp[3])
		}

		if strings.Contains(line, "\"1966720\"") {
			fmt.Println("Found Lethal Company (ID 1966720), exiting.")
			break
		}
	}

	lethalCompanyPath := fmt.Sprintf("%s\\steamapps\\common\\Lethal Company\\", lastPath)
	lethalCompanyPath = strings.ReplaceAll(lethalCompanyPath, "\\\\", "\\")

	fmt.Printf("Lathal Company path: %s\n", lethalCompanyPath)
}
