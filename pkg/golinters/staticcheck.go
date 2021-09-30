package golinters

import (
	"honnef.co/go/tools/staticcheck"

	"github.com/hitzhangjie/go-readability/pkg/config"
	"github.com/hitzhangjie/go-readability/pkg/golinters/goanalysis"
)

func NewStaticcheck(settings *config.StaticCheckSettings) *goanalysis.Linter {
	cfg := staticCheckConfig(settings)

	analyzers := setupStaticCheckAnalyzers(staticcheck.Analyzers, getGoVersion(settings), cfg.Checks)

	return goanalysis.NewLinter(
		"staticcheck",
		"Staticcheck is a go vet on steroids, applying a ton of static analysis checks",
		analyzers,
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}
