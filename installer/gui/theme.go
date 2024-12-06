package gui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

var (
	colorPink = color.RGBA{R: 151, G: 18, B: 197, A: 255}
	colorBlue = color.RGBA{R: 15, G: 229, B: 230, A: 255}
)

type Theme struct {
	fyne.Theme
}

func (t Theme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		return color.Black
	case theme.ColorNameInputBackground:
		return color.Black
	case theme.ColorNameInputBorder:
		return colorPink
	case theme.ColorNameForeground:
		return color.White
	case theme.ColorNameButton:
		return color.Black
	case theme.ColorNameMenuBackground:
		return color.Black
	case theme.ColorNameError:
		return color.Black
	case theme.ColorNameOverlayBackground:
		return color.Black
	}
	return t.Theme.Color(name, theme.VariantDark)
}

func (t Theme) Icon(name fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(name)
}

func (t Theme) Font(style fyne.TextStyle) fyne.Resource {
	return theme.DefaultTheme().Font(fyne.TextStyle{})
}

func (t Theme) Size(name fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(name)
}
