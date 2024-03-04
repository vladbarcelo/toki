package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/vladbarcelo/toki/src/helpers"
)

func RenderLine(viewportWidth int) string {
	line := strings.Repeat("─", helpers.Max(0, viewportWidth))
	return lipgloss.JoinHorizontal(lipgloss.Center, line)
}
