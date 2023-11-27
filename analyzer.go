package recover_linter

import (
	"github.com/tulzke/recover-linter/internal"
	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "recoverlint",
		Doc:  "A linter that checks for recover in goroutines",
		Run:  run,
	}
}

func run(pass *analysis.Pass) (any, error) {
	statements := internal.CollectAllGoStatements(pass.Files)
	_ = statements

	diagnostics := internal.Inspect(statements)

	for _, d := range diagnostics {
		pass.Report(d)
	}

	return nil, nil
}
