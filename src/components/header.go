package components

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/vladbarcelo/toki/src/helpers"
	"github.com/vladbarcelo/toki/src/styles"
)

func RenderHeader(
	textInputComponent string,
	viewportWidth int,
) string {
	title := styles.TitleStyle.Render(textInputComponent)
	line := strings.Repeat("─", helpers.Max(0, viewportWidth-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}
