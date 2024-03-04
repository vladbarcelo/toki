package main

// An example program demonstrating the pager component from the Bubbles
// component library.

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/vladbarcelo/toki/src/pages"
	"github.com/vladbarcelo/toki/src/store"
)

// You generally won't need this unless you're processing stuff with
// complicated ANSI escape sequences. Turn it on if you notice flickering.
//
// Also keep in mind that high performance rendering only works for programs
// that use the full size of the terminal. We're enabling that below with
const useHighPerformanceRenderer = true

type Model struct {
	Store *store.Store
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.Store.Page {
	default:
		cmd = pages.UpdateDefaultPage(m.Store, msg)
	}

	return m, cmd
}

func (m Model) View() string {
	switch m.Store.Page {
	default:
		return pages.RenderDefaultPage(m.Store)
	}
}

func initialModel() Model {
	return Model{
		Store: store.NewStore(),
	}
}

func main() {

	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),       // use the full size of the terminal in its "alternate screen buffer"
		tea.WithMouseCellMotion(), // turn on mouse support so we can track the mouse wheel
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("could not run program:", err)
		os.Exit(1)
	}
}
