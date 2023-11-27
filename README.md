recover-linter
=============

go-errorlint is a source code linter for Go that can be used to find unsafe goroutines without recover expressions

Panic in goroutine without recover can crash entire application

## Usage
recover-linter accepts a set of package names similar to golint:
```
recover-linter ./...
```