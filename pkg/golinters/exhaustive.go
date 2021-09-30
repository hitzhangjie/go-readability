package golinters

import (
	"github.com/nishanths/exhaustive"
	"golang.org/x/tools/go/analysis"

	"github.com/hitzhangjie/go-readability/pkg/config"
	"github.com/hitzhangjie/go-readability/pkg/golinters/goanalysis"
)

func NewExhaustive(settings *config.ExhaustiveSettings) *goanalysis.Linter {
	a := exhaustive.Analyzer

	var cfg map[string]map[string]interface{}
	if settings != nil {
		cfg = map[string]map[string]interface{}{
			a.Name: {
				exhaustive.CheckGeneratedFlag:             settings.CheckGenerated,
				exhaustive.DefaultSignifiesExhaustiveFlag: settings.DefaultSignifiesExhaustive,
				exhaustive.IgnorePatternFlag:              settings.IgnorePattern,
			},
		}
	}

	return goanalysis.NewLinter(a.Name, a.Doc, []*analysis.Analyzer{a}, cfg).
		WithLoadMode(goanalysis.LoadModeTypesInfo)
}
