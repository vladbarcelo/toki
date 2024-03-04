package pages

import (
	"fmt"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/vladbarcelo/toki/src/components"
	"github.com/vladbarcelo/toki/src/store"
)

func RenderDefaultPage(s *store.Store) string {
	if !s.Ready {
		return "\n  Initializing..."
	}
	return fmt.Sprintf("%s\n%s\n%s",
		components.RenderHeader(s.TextInput.View(), s.Viewport.Width),
		s.Viewport.View(),
		components.RenderFooter(
			len(s.Content),
			len(s.FilteredContent),
			s.Viewport.ScrollPercent(),
			s.Viewport.Width,
		),
	)
}

func UpdateDefaultPage(s *store.Store, msg tea.Msg) tea.Cmd {
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(
			key.WithKeys("shift+up"),
		)):
			s.SelectedLine -= 1
			s.Viewport.SetContent(components.RenderContent(s.FilteredContent, s.Filter.Cols, s.SelectedLine))
			cmds = append(cmds, viewport.Sync(s.Viewport))
		case key.Matches(msg, key.NewBinding(
			key.WithKeys("shift+down"),
		)):
			s.SelectedLine += 1
			s.Viewport.SetContent(components.RenderContent(s.FilteredContent, s.Filter.Cols, s.SelectedLine))
			cmds = append(cmds, viewport.Sync(s.Viewport))
		case key.Matches(msg, key.NewBinding(
			key.WithKeys("shift+right"),
		)):
			s.Viewport.MouseWheelEnabled = true
			s.Viewport.SetContent(components.RenderDetailedContent(s.Content, s.FilteredContent, s.SelectedLine))
			cmds = append(cmds, viewport.Sync(s.Viewport))
		case key.Matches(msg, key.NewBinding(
			key.WithKeys("ctrl+c", "esc"),
		)):
			return tea.Quit
		case key.Matches(msg, key.NewBinding(
			key.WithKeys("enter", "shift+left"),
		)):
			s.Viewport.MouseWheelEnabled = false
			s.SelectedLine = 0
			s.ReadContentFromFile()
			s.FilteredContent = s.Filter.Filter(s.Content, s.TextInput.Value())
			s.Viewport.SetContent(components.RenderContent(s.FilteredContent, s.Filter.Cols, s.SelectedLine))
			cmds = append(cmds, viewport.Sync(s.Viewport))
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(components.RenderHeader(s.TextInput.View(), s.Viewport.Width))
		footerHeight := lipgloss.Height(components.RenderFooter(len(s.Content), len(s.FilteredContent), s.Viewport.ScrollPercent(), s.Viewport.Width))
		verticalMarginHeight := headerHeight + footerHeight

		if !s.Ready {
			s.Viewport = components.CreateViewport(
				msg.Height,
				msg.Width,
				verticalMarginHeight,
				headerHeight,
				components.RenderContent(s.FilteredContent, s.Filter.Cols, s.SelectedLine),
			)

			// Render the viewport one line below the header.
			s.Viewport.YPosition = headerHeight + 1

			s.Ready = true
		} else {
			s.Viewport.Width = msg.Width
			s.Viewport.Height = msg.Height - verticalMarginHeight
		}

		s.TextInput.Width = s.Viewport.Width / 2

		cmds = append(cmds, viewport.Sync(s.Viewport))
	}

	// Handle keyboard and mouse events in the viewport

	s.Viewport, cmd = s.Viewport.Update(msg)
	cmds = append(cmds, cmd)

	s.TextInput, cmd = s.TextInput.Update(msg)
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}
