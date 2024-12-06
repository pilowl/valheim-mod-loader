package installer

import (
	"bytes"
	"io"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/effects"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/pilowl/lethalpacker/installer/gui"
	"github.com/pilowl/lethalpacker/installer/modder"
	"github.com/pilowl/lethalpacker/installer/reg"
)

type Installer struct {
	gui *gui.GUI
}

const WINDOW_TITLE = "Valgaym Lethal Company Mod Pack"

func New() *Installer {
	lethalCompanyInstallPath := reg.GetLethalCompanyPath()
	modLoader, err := modder.NewLoaderFromZip(modZipBytes)
	if err != nil {
		log.Print(err)
		return nil
	}
	state := gui.BuildInitialState(lethalCompanyInstallPath, modLoader.GetMods())

	gui := gui.Init(WINDOW_TITLE, logoBytes, state, func(state *gui.State) error {
		return modLoader.InstallMods(state.Mods, state.GamePath)
	})
	return &Installer{
		gui: gui,
	}
}

func (i *Installer) Run() {
	go music()
	i.gui.Run()
}

func music() {
	musicReader := bytes.NewReader(musicBytes)

	streamer, format, err := mp3.Decode(io.NopCloser(musicReader))
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	err = speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	if err != nil {
		return
	}

	ctrl := &beep.Ctrl{Streamer: beep.Loop(-1, streamer), Paused: false}
	volume := &effects.Volume{
		Streamer: ctrl,
		Base:     2,
		Volume:   -5,
		Silent:   false,
	}

	speaker.Play(beep.Seq(volume, beep.Callback(func() {})))

	for {
	}
}
