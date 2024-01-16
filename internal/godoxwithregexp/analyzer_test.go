package godoxwithregexp

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
	repoDir := filepath.Dir(filepath.Dir(wd))
	testdata := filepath.Join(repoDir, "testdata")
	testAnalyzerResult := analysistest.Run(t, testdata, Analyzer, "basic")
	assert.Equal(t, 1, len(testAnalyzerResult))
}
