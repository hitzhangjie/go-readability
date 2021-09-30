package golinters

import (
	"github.com/kyoh86/exportloopref"
	"golang.org/x/tools/go/analysis"

	"github.com/hitzhangjie/go-readability/pkg/golinters/goanalysis"
)

func NewExportLoopRef() *goanalysis.Linter {
	a := exportloopref.Analyzer

	return goanalysis.NewLinter(
		a.Name,
		a.Doc,
		[]*analysis.Analyzer{a},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
