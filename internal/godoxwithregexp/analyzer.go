package godoxwithregexp

import (
	"bufio"
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"path/filepath"
	"regexp"
	"unicode"
	"unicode/utf8"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"

	"github.com/SilverdewBaker/godox-with-regexp/types"
)

// Analyzer is an Analyzer for checking keyword comment format.
var Analyzer = &analysis.Analyzer{
	Name:     "godoxwithregexp",
	Doc:      "Checks the format of keyword comments",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

var FormatRules = []types.GoDoxFormatRule{
	{
		Keyword:           "TODO",
		RegularExpression: "^TODO\\(\\d{8}\\)\\s+.+$",
	},
}

// run is the main logic of the keyword comment format analyzer.
func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	formatRules := FormatRules
	inspector.Preorder(nil, func(node ast.Node) {
		var messages []types.Message
		for _, file := range pass.Files {
			for _, c := range file.Comments {
				for _, ci := range c.List {
					messages = append(messages, getMessagesFormat(ci, pass.Fset, formatRules)...)
				}
			}
		}

		for _, message := range messages {
			pass.Reportf(token.Pos(message.Pos.Offset), message.Message)
		}
	})

	return nil, nil
}

func getMessagesFormat(comment *ast.Comment, fset *token.FileSet, formatRules []types.GoDoxFormatRule) []types.Message {
	commentText := extractComment(comment.Text)
	b := bufio.NewReader(bytes.NewBufferString(commentText))
	var comments []types.Message

	for lineNum := 0; ; lineNum++ {
		line, _, err := b.ReadLine()
		if err != nil {
			break
		}

		const minimumSize = 4
		sComment := bytes.TrimSpace(line)
		if len(sComment) < minimumSize {
			continue
		}

		for _, formatRule := range formatRules {
			kw := formatRule.Keyword
			formatPattern := formatRule.RegularExpression

			if lkw := len(kw); !(bytes.EqualFold([]byte(kw), sComment[0:lkw]) &&
				!hasAlphanumRuneAdjacent(sComment[lkw:])) {
				continue
			}

			// check the format
			if formatPattern != "" && isFormatted(formatPattern, string(sComment)) {
				continue
			}

			pos := fset.Position(comment.Pos())
			// trim the comment
			const commentLimit = 40
			if len(sComment) > commentLimit {
				sComment = []byte(fmt.Sprintf("%.40s...", sComment))
			}

			comments = append(comments, types.Message{
				Pos: pos,
				Message: fmt.Sprintf(
					"%s:%d: Line does not match the expected format: %s, %q",
					filepath.Clean(pos.Filename),
					pos.Line+lineNum,
					formatPattern,
					sComment,
				),
			})

			break
		}
	}

	return comments
}

func isFormatted(regularExpression, input string) bool {
	regex := regexp.MustCompile(regularExpression)
	if regex.MatchString(input) {
		return true
	} else {
		return false
	}
}

func extractComment(commentText string) string {
	switch commentText[1] {
	case '/':
		commentText = commentText[2:]
		if len(commentText) > 0 && commentText[0] == ' ' {
			commentText = commentText[1:]
		}
	case '*':
		commentText = commentText[2 : len(commentText)-2]
	}

	return commentText
}

func hasAlphanumRuneAdjacent(rest []byte) bool {
	if len(rest) == 0 {
		return false
	}

	switch rest[0] { // most common cases
	case ':', ' ', '(':
		return false
	}

	r, _ := utf8.DecodeRune(rest)

	return unicode.IsLetter(r) || unicode.IsNumber(r) || unicode.IsDigit(r)
}
