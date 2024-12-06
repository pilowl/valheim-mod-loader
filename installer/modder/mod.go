package modder

type Mod struct {
	Name           string
	RelPath        string
	Active         bool
	ClientSide     bool
	ChildFilePaths []string
}

var DLLMappings map[string]Mod = map[string]Mod{
	"LateCompanyV1.0.16.dll": {
		Name:       "Late Company",
		Active:     true,
		ClientSide: false,
	},
	"BoomboxController.dll": {
		Name:       "Boombox Controller",
		Active:     true,
		ClientSide: false,
	},
	"BrutalCompanyPlus.dll": {
		Name:       "Brutal Company Plus",
		Active:     false,
		ClientSide: false,
	},
	"FasterItemDropship.dll": {
		Name:       "Faster Item Drop",
		Active:     true,
		ClientSide: false,
	},
	"MoreCompany.dll": {
		Name:       "More Company",
		Active:     true,
		ClientSide: false,
	},
	"YippeeMod.dll": {
		Name:           "Yippee",
		Active:         true,
		ClientSide:     true,
		ChildFilePaths: []string{"BepInEx/plugins/yippeesound"},
	},
	"HDLethalCompany.dll": {
		Name:           "HD Lethal Company",
		Active:         false,
		ClientSide:     true,
		ChildFilePaths: []string{"BepInEx/plugins/HDLethalCompany/hdlethalcompany"},
	},
	"SkinwalkerMod.dll": {
		Name:       "Skinwalker Mod",
		Active:     true,
		ClientSide: false,
	},
}
