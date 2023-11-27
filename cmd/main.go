package main

import (
	recover_linter "github.com/tulzke/recover-linter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(recover_linter.NewAnalyzer())
}
