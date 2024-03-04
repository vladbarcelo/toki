package components

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
)

func CreateViewport(
	windowHeight int,
	windowWidth int,
	verticalMarginHeight int,
	headerHeight int,
	contentComponent string,
) viewport.Model {
	viewportComponent := viewport.New(windowWidth, windowHeight-verticalMarginHeight)
	viewportComponent.YPosition = headerHeight
	viewportComponent.HighPerformanceRendering = true
	viewportComponent.SetContent(contentComponent)
	viewportComponent.MouseWheelEnabled = false
	viewportComponent.KeyMap = viewport.KeyMap{
		PageDown: key.NewBinding(
			key.WithKeys("pgdown"),
			key.WithHelp("pgdn", "page down"),
		),
		PageUp: key.NewBinding(
			key.WithKeys("pgup"),
			key.WithHelp("pgup", "page up"),
		),
		HalfPageUp: key.NewBinding(
			key.WithKeys("ctrl+u"),
			key.WithHelp("ctrl+u", "½ page up"),
		),
		HalfPageDown: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "½ page down"),
		),
		Up: key.NewBinding(
			key.WithKeys("up"),
			key.WithHelp("↑", "up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down"),
			key.WithHelp("↓", "down"),
		),
	}

	return viewportComponent
}
