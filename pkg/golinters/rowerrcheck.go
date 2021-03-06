package golinters

import (
	"github.com/jingyugao/rowserrcheck/passes/rowserr"
	"golang.org/x/tools/go/analysis"

	"github.com/hitzhangjie/go-readability/pkg/golinters/goanalysis"
	"github.com/hitzhangjie/go-readability/pkg/lint/linter"
)

func NewRowsErrCheck() *goanalysis.Linter {
	analyzer := rowserr.NewAnalyzer()
	return goanalysis.NewLinter(
		"rowserrcheck",
		"checks whether Err of rows is checked successfully",
		[]*analysis.Analyzer{analyzer},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo).
		WithContextSetter(func(lintCtx *linter.Context) {
			pkgs := lintCtx.Settings().RowsErrCheck.Packages
			analyzer.Run = rowserr.NewRun(pkgs...)
		})
}
