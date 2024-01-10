package types

import "go/token"

// Message contains a message and position.
type Message struct {
	Pos     token.Position
	Message string
}

type GoDoxSettings struct {
	FormatRules []GoDoxFormatRule `yaml:"format-rules,omitempty"`
}

type GoDoxFormatRule struct {
	Keyword           string `yaml:"keyword,omitempty"`
	RegularExpression string `yaml:"regularExpression,omitempty"`
}

type YamlSettings struct {
	LintersSettings struct {
		Custom struct {
			Godoxwithregexp struct {
				Path        string        `yaml:"path,omitempty"`
				Description string        `yaml:"description,omitempty"`
				OriginalURL string        `yaml:"original-url,omitempty"`
				Settings    GoDoxSettings `yaml:"settings,omitempty"`
			} `yaml:"godoxwithregexp,omitempty"`
		} `yaml:"custom,omitempty"`
	} `yaml:"linters-settings,omitempty"`
}
