package internal

import "go/ast"

func CollectAllGoStatements(files []*ast.File) []*ast.GoStmt {
	result := make([]*ast.GoStmt, 0)
	for _, file := range files {
		ast.Inspect(file, func(node ast.Node) bool {
			gostmt, ok := node.(*ast.GoStmt)
			if !ok {
				return true
			}

			result = append(result, gostmt)
			return true
		})
	}

	return result
}
