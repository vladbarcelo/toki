package components

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/vladbarcelo/toki/src/helpers"
	"github.com/vladbarcelo/toki/src/styles"
)

func RenderFooter(
	contentLength int,
	filteredContentLength int,
	scrollPercent float64,
	viewportWidth int,
) string {
	lineCountSuffix := ","
	if filteredContentLength > 1000 {
		lineCountSuffix = "\033[31;1;m!\033[0m,"
	}
	info := styles.InfoStyle.Render(fmt.Sprintf("filtered: %d%s total: %d │ %3.f%%", filteredContentLength, lineCountSuffix, contentLength, scrollPercent*100))
	line := strings.Repeat("─", helpers.Max(0, viewportWidth-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}
