package parser

import (
	"strconv"
	"strings"

	"github.com/titanous/json5"
)

func (p *Parser) ParseContent(content string) []Line {
	rawLines := strings.Split(content, "\n")

	lines := make([]Line, 0, len(rawLines))

	for index, rawLine := range rawLines {
		// rawLine = strings.ReplaceAll(rawLine, "c", "")
		line := Line{Raw: rawLine, Index: strconv.Itoa(index)}
		json5.Unmarshal([]byte(rawLine), &line.Parsed)
		lines = append(lines, line)
	}

	return lines
}
