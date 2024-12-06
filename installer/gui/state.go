package gui

import "github.com/pilowl/lethalpacker/installer/modder"

type State struct {
	GamePath string
	Mods     []modder.Mod
}

func BuildInitialState(gamePath string, mods []modder.Mod) *State {
	return &State{
		GamePath: gamePath,
		Mods:     mods,
	}
}
