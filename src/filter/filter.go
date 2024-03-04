package filter

import (
	"github.com/vladbarcelo/toki/src/parser"
)

func (f *Filter) Filter(content []parser.Line, rawQuery string) []parser.Line {
	f.RawSearchQuery = rawQuery
	f.parseRawSearchQuery()

	var filteredContent = make([]parser.Line, 0)

	for _, line := range content {
		if len(f.RawSearchQuery) == 0 {
			filteredContent = append(filteredContent, line)
		} else {
			lineToAppend := f.filterLine(line)

			if lineToAppend != nil {
				newLine := *lineToAppend

				if len(f.Cols) != 0 {
					newLine = f.filterColumnsForLine(newLine)
				}

				filteredContent = append(filteredContent, newLine)
			}
		}
	}

	return filteredContent
}

func (f *Filter) filterLine(line parser.Line) *parser.Line {
	// Filter out raw lines if cols are specified
	if len(f.Cols) != 0 && line.Parsed == nil {
		return nil
	}

	// if there is at least some search query
	// and some detected regex
	// and the line doesn't match the regex, skip it
	if len(f.RawSearchQuery) != 0 && f.RawRegexContent != "" && !f.searchQueryRegex.MatchString(line.Raw) {
		return nil
	}

	// Filter by time interval
	if line.Parsed != nil && f.Interval.From != 0 {
		if lineTime, ok := line.Parsed["time"]; ok {
			if lineTime.(float64) < f.Interval.From {
				return nil
			}
		}
	}

	// filter by key:value
	if line.Parsed != nil && len(f.Keys) != 0 {
		for _, key := range f.Keys {
			if val, ok := line.Parsed[key.Key]; ok {
				if val != key.Val {
					return nil
				}
			} else {
				return nil
			}
		}
	}

	return &line
}

func (f *Filter) filterColumnsForLine(line parser.Line) parser.Line {
	cols := f.Cols
	data := line.Parsed
	line.Parsed = make(map[string]interface{}, 0)
	for _, col := range cols {
		if val, ok := data[col]; ok {
			line.Parsed[col] = val
		}
	}

	return line
}
