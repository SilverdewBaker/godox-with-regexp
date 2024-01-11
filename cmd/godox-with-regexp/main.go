// +gocover:ignore:file analyzer entrypoint
package main

import (
	"golang.org/x/tools/go/analysis"

	"silverdewbaker/todo-with-regexp/analyzer"
)

// This struct must be defined
type analyzerPlugin struct{}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin

// This must be implemented
func New(conf any) ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}
