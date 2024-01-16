// +gocover:ignore:file analyzer entrypoint
package main

import (
	"github.com/SilverdewBaker/godox-with-regexp/internal/godoxwithregexp"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(godoxwithregexp.Analyzer)
}
