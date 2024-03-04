package styles

import "github.com/charmbracelet/lipgloss"

// colors
var (
	WhiteColor    = lipgloss.Color("#869496")
	BlackColor    = lipgloss.Color("#14272f")
	DarkBlueColor = lipgloss.Color("#173540")
	GreyColor     = lipgloss.Color("#676767")
	BlueColor     = lipgloss.Color("#4689cc")
	YellowColor   = lipgloss.Color("#ae8a2d")
	RedColor      = lipgloss.Color("#cc241d")
)

var (
	DefaultWhiteText = lipgloss.NewStyle().
				Foreground(WhiteColor)

	DefaultKeyStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(DarkBlueColor).
			Background(GreyColor)

	InfoValueStyle = lipgloss.NewStyle().
			Foreground(BlackColor).
			Background(BlueColor)

	WarnValueStyle = lipgloss.NewStyle().
			Foreground(DarkBlueColor).
			Background(YellowColor)

	ErrorValueStyle = lipgloss.NewStyle().
			Foreground(DarkBlueColor).
			Background(RedColor)

	FatalValueStyle = lipgloss.NewStyle().
			Foreground(DarkBlueColor).
			Background(GreyColor)

	TitleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	InfoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return TitleStyle.Copy().BorderStyle(b)
	}()
)
