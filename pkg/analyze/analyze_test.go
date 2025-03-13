package analyze

import (
	"github.com/skirge/yataf/pkg/fetch"
	"testing"
)

func TestAnalyzeContent(t *testing.T) {

	Config := map[string]map[string]int{
		"../../tests/1.test": {
			"Results":     16,
			"Extractions": 16,
		},
		// more to come :P
	}

	for File, ExpectedResults := range Config {
		Content := fetch.GetContentOfLocalFile(File)
		Analyzed := AnalyzeContent(Content, []string{"all"})

		ResultsCount := ExpectedResults["Results"]
		ExtractionsCount := ExpectedResults["Extractions"]

		if len(Analyzed.Result) != ResultsCount {
			t.Errorf("Expected %d results, got %d", ResultsCount, len(Analyzed.Result))
		}

		if len(Analyzed.Result) != ExtractionsCount {
			t.Errorf("Expected %d extractions, got %d", ExtractionsCount, len(Analyzed.Result))
		}
	}
}
