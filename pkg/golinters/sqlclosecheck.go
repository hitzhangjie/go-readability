package golinters

import (
	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis"

	"github.com/hitzhangjie/go-readability/pkg/golinters/goanalysis"
)

func NewSQLCloseCheck() *goanalysis.Linter {
	analyzers := []*analysis.Analyzer{
		analyzer.NewAnalyzer(),
	}

	return goanalysis.NewLinter(
		"sqlclosecheck",
		"Checks that sql.Rows and sql.Stmt are closed.",
		analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
