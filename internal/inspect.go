package internal

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func Inspect(statements []*ast.GoStmt) []analysis.Diagnostic {
	diagnostics := make([]analysis.Diagnostic, 0)
	for _, statement := range statements {
		if statement.Call == nil {
			continue
		}

		expr := statement.Call.Fun
		var hasRecover bool
		switch v := expr.(type) {
		case *ast.FuncLit:
			if v == nil || v.Body == nil {
				break
			}

			for _, stmt := range v.Body.List {
				hasRecover = isNodeHasRecoverInDefer(stmt)
				if hasRecover {
					break
				}
			}
		case *ast.Ident:
			if v == nil || v.Obj == nil || v.Obj.Decl == nil {
				break
			}

			declNode, ok := v.Obj.Decl.(ast.Node)
			if ok {
				hasRecover = isNodeHasRecoverInDefer(declNode)
			}
		}

		if hasRecover {
			continue
		}

		diagnostics = append(diagnostics, analysis.Diagnostic{
			Pos:     statement.Pos(),
			Message: fmt.Sprintf("goroutine without recover. panic here can crash entire application"),
		})
	}

	return diagnostics
}
