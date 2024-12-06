package gui

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/pilowl/lethalpacker/installer/modder"
)

type GUI struct {
	app    fyne.App
	window fyne.Window
	logo   []byte
	state  *State

	installHandler InstallHandler
}

func Init(windowTitle string, logo []byte, state *State, instalHandller InstallHandler) *GUI {
	app := app.New()
	app.Settings().SetTheme(&Theme{app.Settings().Theme()})

	window := app.NewWindow(windowTitle)
	window.Resize(fyne.NewSize(700, 800))
	window.SetFixedSize(true)

	return &GUI{
		app:            app,
		window:         window,
		state:          state,
		logo:           logo,
		installHandler: instalHandller,
	}
}

func (gui *GUI) Run() {
	gui.Build()
	gui.window.ShowAndRun()
}

func (gui *GUI) Build() {
	image := canvas.NewImageFromReader(bytes.NewReader(gui.logo), "logo")
	image.FillMode = canvas.ImageFillContain

	gui.window.SetContent(container.New(layout.NewGridLayout(0), image, gui.createForm()))

}

func (gui *GUI) createForm() *fyne.Container {
	textBox := widget.NewEntry()
	textBox.SetText(gui.state.GamePath)

	btnSelectPath := widget.NewButton("Select Lethal Company Path", func() {
		dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if err != nil || uri == nil {
				return
			}
			selectedPath := strings.ReplaceAll(uri.Path(), `/`, `\`)
			textBox.SetText(selectedPath)
			gui.state.GamePath = selectedPath
		}, gui.window).Show()
	})
	btnUploadMods := widget.NewButton("Install Mods", func() {
		resultTitle := "Installation Result"
		err := gui.installHandler(gui.state)
		switch err {
		case nil:
			dialog.NewInformation(resultTitle, "Great success!", gui.window).Show()
		default:
			dialog.NewError(fmt.Errorf("Something went wrong: %s", err.Error()), gui.window).Show()
		}
	})
	modSelector := gui.createModSelector()

	vContainer := container.NewGridWrap(fyne.NewSize(500, 35), textBox, btnSelectPath, btnUploadMods)

	newVContainerWithScroll := container.NewGridWithRows(2, vContainer, container.NewStack(modSelector))

	contCentered := container.NewCenter(newVContainerWithScroll)
	return contCentered
}

func (gui *GUI) createModSelector() *container.Scroll {
	modChecks := gui.createChecks(gui.state.Mods)
	scrollWidget := container.NewVScroll(container.NewGridWithColumns(2, modChecks...))
	return scrollWidget
}

func (gui *GUI) createChecks(mods []modder.Mod) []fyne.CanvasObject {
	sort.Slice(mods, func(a, b int) bool {
		if mods[a].Active && !mods[b].Active {
			return true
		}
		return mods[a].Name < mods[b].Name
	})
	checks := make([]fyne.CanvasObject, 0, len(mods))
	for i, mod := range mods {
		modName := mod.Name
		if mod.ClientSide {
			modName = fmt.Sprintf("%s (Client-side)", modName)
		}
		newCheck := widget.NewCheck(modName, checkChangeStateFunc(mods, i))
		newCheck.SetChecked(mod.Active)
		checks = append(checks, newCheck)
	}
	return checks
}

func checkChangeStateFunc(mods []modder.Mod, i int) func(b bool) {
	return func(b bool) {
		mods[i].Active = b
	}
}
