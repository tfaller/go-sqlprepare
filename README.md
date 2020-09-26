# go-sqlprepare
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tfaller/go-sqlprepare)](https://pkg.go.dev/github.com/tfaller/go-sqlprepare)

Simple module to prepare multiple SQL statements at once.
Useful at server starts.

```go
var stmtA, stmtB *sql.Stmt

err := sqlprepare.Prepare(db,
		sqlprepare.ToPrepare{Name: "stmtA", Target: &stmtA,
			Query: "SELECT * FROM A"},

		sqlprepare.ToPrepare{Name: "stmtB", Target: &stmtB,
			Query: "INSERT INTO B (a, b, c) VALUES (?, ?, ?)"},
        )

log.Print(err)
```