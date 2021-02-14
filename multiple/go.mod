module example.com/mutiple

go 1.15

require (
	example.com/extral v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
)

replace example.com/extral => ../extral
