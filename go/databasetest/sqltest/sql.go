package sqltest

import (
	"database/sql"
	"testing"
)

// Open opens a database specified by its database driver name and a
// driver-specific data source name, usually consisting of at least a
// database name and connection information.
//
// Most users will open a database via a driver-specific connection
// helper function that returns a [*DB]. No database drivers are included
// in the Go standard library. See https://golang.org/s/sqldrivers for
// a list of third-party drivers.
//
// Open may just validate its arguments without creating a connection
// to the database. To verify that the data source name is valid, call
// [DB.Ping].
//
// The returned [DB] is safe for concurrent use by multiple goroutines
// and maintains its own pool of idle connections. Thus, the Open
// function should be called just once. It is rarely necessary to
// close a [DB].
func Open(t testing.TB, driverName, dataSourceName string) *sql.DB {
	t.Helper()
	p0, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		t.Fatal(err)
	}
	return p0
}
