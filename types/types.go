package types

import "go/token"

// Message contains a message and position.
type Message struct {
	Pos     token.Position
	Message string
}

type GoDoxSettings struct {
	FormatRules []GoDoxFormatRule
}

type GoDoxFormatRule struct {
	Keyword           string
	RegularExpression string
}
