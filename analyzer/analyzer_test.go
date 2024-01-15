package analyzer

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tj/assert"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	assert.Nil(t, err)

	testdata := filepath.Join(filepath.Dir(wd), "testdata")
	testAnalyzerResult := analysistest.Run(t, testdata, Analyzer, "todo.go")
	assert.Equal(t, 1, len(testAnalyzerResult))
}
