package components

import (
	"fmt"
	"strconv"

	"github.com/TylerBrock/colorjson"
	"github.com/vladbarcelo/toki/src/parser"
)

func RenderDetailedContent(content []parser.Line, filteredContent []parser.Line, selectedLine int) string {
	indexStr := filteredContent[selectedLine].Index

	index, _ := strconv.Atoi(indexStr)

	line := content[index]

	if line.Parsed == nil {
		return fmt.Sprintf("\n%s", line.Raw)
	} else {
		f := colorjson.NewFormatter()
		f.Indent = 2

		res, _ := f.Marshal(line.Parsed)
		return fmt.Sprintf("\n%s", res)
	}
}
