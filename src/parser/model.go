package parser

type Parser struct {
}

type Line struct {
	Index  string
	Raw    string
	Parsed map[string]interface{}
}

func NewParser() *Parser {
	return &Parser{}
}
