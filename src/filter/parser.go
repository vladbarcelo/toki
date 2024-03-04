package filter

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/titanous/json5"
)

func (f *Filter) parseRawSearchQuery() {
	// parse columns
	f.Cols = make([]string, 0)
	colsRegex := *regexp.MustCompile(`^\[.*?\]`)
	rawCols := colsRegex.FindString(f.RawSearchQuery)
	if rawCols != "" {

		for _, col := range strings.Split(rawCols[1:len(rawCols)-1], ",") {
			f.Cols = append(f.Cols, strings.Trim(col, " "))
		}
	}

	// parse key:value
	f.Keys = make([]key, 0)
	kvRegex := *regexp.MustCompile(`{.*}`)
	rawKV := kvRegex.FindString(f.RawSearchQuery)
	if rawKV != "" {
		var parsedKV map[string]interface{}
		json5.Unmarshal([]byte(rawKV), &parsedKV)

		for k, v := range parsedKV {
			f.Keys = append(f.Keys, key{k, v})
		}
	}

	// parse interval, case 1: ~last5m
	f.Interval = timeInterval{
		From: 0,
		To:   0,
	}
	timeIntervalRegex := *regexp.MustCompile(`~ ?last[0-9]+[smh]`)
	rawTimeInterval := timeIntervalRegex.FindString(f.RawSearchQuery)
	if rawTimeInterval != "" {
		rawTimeInterval = strings.ReplaceAll(rawTimeInterval, "~", "")
		rawTimeInterval = strings.ReplaceAll(rawTimeInterval, " ", "")
		rawTimeInterval = strings.ReplaceAll(rawTimeInterval, "last", "")

		var quant time.Duration

		switch rawTimeInterval[len(rawTimeInterval)-1] {
		case 's':
			quant = time.Second
		case 'm':
			quant = time.Minute
		case 'h':
			quant = time.Hour
		}

		rawDuration, _ := strconv.Atoi(rawTimeInterval[:len(rawTimeInterval)-1])

		f.Interval = timeInterval{
			From: float64(time.Now().Add(-time.Duration(rawDuration) * quant).UnixMilli()),
		}
	}

	// parse regex or use the entire input as a regex
	searchQueryRegex := *regexp.MustCompile(`\|= ?".*?"`)
	searchQuery := searchQueryRegex.FindString(f.RawSearchQuery)
	if searchQuery != "" {
		searchQuery = strings.ReplaceAll(searchQuery, "\"", "")
		searchQuery = strings.ReplaceAll(searchQuery, "|", "")
		searchQuery = strings.ReplaceAll(searchQuery, "=", "")
		f.RawRegexContent = strings.Trim(searchQuery, " ")
		f.searchQueryRegex = *regexp.MustCompile(fmt.Sprintf("%s", f.RawRegexContent))
	}

	if len(searchQuery) == 0 && rawKV == "" && rawCols == "" {
		f.RawRegexContent = f.RawSearchQuery
		f.searchQueryRegex = *regexp.MustCompile(fmt.Sprintf("%s", f.RawRegexContent))
	}

	optionsToSave := Filter{
		RawSearchQuery: f.RawSearchQuery,
	}

	savedOptions, _ := json.Marshal(optionsToSave)
	err := os.WriteFile("options.json", savedOptions, 0644)
	if err != nil {
		panic(err)
	}
}
