package golinters

import (
	"github.com/sonatard/noctx"
	"golang.org/x/tools/go/analysis"

	"github.com/hitzhangjie/go-readability/pkg/golinters/goanalysis"
)

func NewNoctx() *goanalysis.Linter {
	analyzers := []*analysis.Analyzer{
		noctx.Analyzer,
	}

	return goanalysis.NewLinter(
		"noctx",
		"noctx finds sending http request without context.Context",
		analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
