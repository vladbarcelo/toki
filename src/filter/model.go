package filter

import (
	"encoding/json"
	"os"
	"regexp"
)

type Filter struct {
	Cols             []string
	Keys             []key
	RawSearchQuery   string `json:"rawSearchQuery"`
	RawRegexContent  string
	Interval         timeInterval
	searchQueryRegex regexp.Regexp
}

type key struct {
	Key string
	Val interface{}
}

type timeInterval struct {
	From float64
	To   float64
}

func NewFilter() *Filter {
	f, _ := os.ReadFile("options.json")
	if len(f) > 0 {
		var options Filter
		json.Unmarshal(f, &options)
		options.parseRawSearchQuery()
		return &options
	}

	return &Filter{
		Cols: make([]string, 0),
	}
}
