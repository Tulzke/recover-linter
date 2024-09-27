package recover_linter

import (
	"github.com/golangci/plugin-module-register/register"
	"github.com/tulzke/recover-linter/internal"
	"golang.org/x/tools/go/analysis"
)

const linterName = "recoverlint"

func init() {
	register.Plugin(linterName, NewPlugin)
}

func NewPlugin(_ any) (register.LinterPlugin, error) {
	return Plugin{}, nil
}

type Plugin struct{}

func (Plugin) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		NewAnalyzer(),
	}, nil
}

func (Plugin) GetLoadMode() string {
	return register.LoadModeTypesInfo
}

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: linterName,
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
