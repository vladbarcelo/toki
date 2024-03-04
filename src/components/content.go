package components

import (
	"fmt"
	"strings"
	"time"

	"github.com/vladbarcelo/toki/src/parser"
	"github.com/vladbarcelo/toki/src/styles"
)

var levels = map[float64]string{
	10: "trace",
	20: "debug",
	30: "info",
	40: "warn",
	50: "error",
	60: "fatal",
}

func RenderContent(content []parser.Line, enabledCols []string, selectedLine int) string {

	if len(content) > 1000 {
		content = content[0:1000]
	}

	res := "\n"

	for index, line := range content {
		rawLine := ""
		if line.Parsed == nil {
			line.Raw = strings.ReplaceAll(line.Raw, "c", "")
			rawLine += fmt.Sprintf("%s", line.Raw)
		} else {
			if len(enabledCols) != 0 {
				for _, col := range enabledCols {
					value := line.Parsed[col]
					rawLine += renderColumn(index%2 == 0, col, value)
				}
			} else {
				for col, value := range line.Parsed {
					rawLine += renderColumn(index%2 == 0, col, value)
				}
			}
		}
		lineIndex := line.Index
		if selectedLine == index {
			lineIndex = "> " + lineIndex
		}
		res += fmt.Sprintf("%6s %s\n", lineIndex, rawLine)
	}

	return res
}

func renderColumn(dim bool, col string, value interface{}) string {
	switch col {
	default:
		value = fmt.Sprintf("%s", value)
	case "level":
		value = renderLevelValue(dim, value)
	case "time":
		value = renderTimeValue(dim, int64(value.(float64)))
	}
	return fmt.Sprintf(" %s: %s ", styles.DefaultKeyStyle.Faint(dim).Render(" "+col+" "), styles.DefaultWhiteText.Faint(dim).Render(value.(string)))
}

func renderLevelValue(dim bool, value interface{}) string {
	parsedValue := "unknown"

	switch v := value.(type) {
	case string:
		parsedValue = v
	case float64:
		parsedValue = levels[v]
	}

	switch parsedValue {
	case "info":
		return styles.InfoValueStyle.Faint(dim).Render(" " + parsedValue + " ")
	case "warn":
		return styles.WarnValueStyle.Faint(dim).Render(" " + parsedValue + " ")
	case "error":
		return styles.ErrorValueStyle.Faint(dim).Render(" " + parsedValue + " ")
	case "fatal":
		return styles.FatalValueStyle.Faint(dim).Render(" " + parsedValue + " ")
	default:
		return styles.DefaultWhiteText.Faint(dim).Render(" " + parsedValue + " ")
	}
}

func renderTimeValue(dim bool, value int64) string {
	t := time.UnixMilli(value)
	return styles.DefaultWhiteText.Faint(dim).Render(fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second()))
}
