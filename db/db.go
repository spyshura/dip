package db

import "database/sql"

func A() {
	sql.Open("go-mssqldb", "dipdb")
}
