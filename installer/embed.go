package installer

import _ "embed"

var (
	//go:embed res/logo.jpg
	logoBytes []byte

	//go:embed pack/valgaym_mod_pack.zip
	modZipBytes []byte

	//go:embed res/music.mp3
	musicBytes []byte
)
