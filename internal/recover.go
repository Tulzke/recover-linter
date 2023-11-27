package internal

import "go/ast"

func isNodeHasRecoverInDefer(fun ast.Node) bool {
	var hasRecover bool
	ast.Inspect(fun, func(node ast.Node) bool {
		if hasRecover {
			return false
		}

		if node == nil {
			return false
		}

		deferStatement, ok := node.(*ast.DeferStmt)
		if !ok {
			return true
		}

		hasRecover = isNodeHasRecover(deferStatement.Call.Fun)
		return false
	})

	return hasRecover
}

func isNodeHasRecover(fun ast.Node) bool {
	var hasRecover bool
	ast.Inspect(fun, func(node ast.Node) bool {
		if hasRecover {
			return false
		}

		if node == nil {
			return false
		}

		switch v := node.(type) {
		case *ast.Ident:
			if v.Name == "recover" {
				hasRecover = true
				return false
			}

			return true
		case *ast.GoStmt:
			return false
		default:
			return true
		}
	})

	return hasRecover
}
