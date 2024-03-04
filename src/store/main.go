package store

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/vladbarcelo/toki/src/filter"
	"github.com/vladbarcelo/toki/src/parser"
)

type Store struct {
	Page            string
	Content         []parser.Line
	FilteredContent []parser.Line
	Ready           bool
	Viewport        viewport.Model
	TextInput       textinput.Model
	Filter          *filter.Filter
	Parser          *parser.Parser
	FileSize        int64
	SelectedLine    int
}

func NewStore() *Store {
	ti := textinput.New()
	ti.Placeholder = "[columns] -> {key:value} |= \"search string\""
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 50

	parser := parser.NewParser()

	filter := filter.NewFilter()

	ti.SetValue(filter.RawSearchQuery)

	store := &Store{
		TextInput:    ti,
		Parser:       parser,
		Filter:       filter,
		FileSize:     0,
		SelectedLine: 0,
		Page:         "default",
	}

	store.ReadContentFromFile()

	return store
}

func (s *Store) ReadContentFromFile() {
	filePath := os.Args[1]
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("could not stat file %s:", filePath), err)
		os.Exit(1)
	}

	// File has no data or is not updated, skipping
	if fileInfo.Size() == s.FileSize {
		return
	}

	rawContent, _ := os.ReadFile(filePath)
	if err != nil {
		fmt.Println(fmt.Sprintf("could not load file %s:", filePath), err)
		os.Exit(1)
	}

	s.Content = s.Parser.ParseContent(string(rawContent))

	s.FilteredContent = s.Filter.Filter(s.Content, s.Filter.RawSearchQuery)

	s.FileSize = fileInfo.Size()
}
