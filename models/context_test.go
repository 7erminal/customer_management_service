package models

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"testing"

	"github.com/beego/beego/v2/client/orm"
)

type testDriver struct{}

func (testDriver) Open(string) (driver.Conn, error) { return testConn{}, nil }

type testConn struct{}

func (testConn) Prepare(string) (driver.Stmt, error) { return nil, errors.New("unexpected prepare") }
func (testConn) Close() error                        { return nil }
func (testConn) Begin() (driver.Tx, error)           { return nil, errors.New("unexpected transaction") }
func (testConn) Ping(context.Context) error          { return nil }
func (testConn) QueryContext(context.Context, string, []driver.NamedValue) (driver.Rows, error) {
	return testRows{}, nil
}

type testRows struct{}

func (testRows) Columns() []string         { return []string{"v"} }
func (testRows) Close() error              { return nil }
func (testRows) Next([]driver.Value) error { return io.EOF }
func TestCriticalModelQueriesRespectCancellation(t *testing.T) {
	sql.Register("trace-test-driver", testDriver{})
	if err := orm.RegisterDriver("trace-test-driver", orm.DRMySQL); err != nil {
		t.Fatal(err)
	}
	if err := orm.RegisterDataBase("default", "trace-test-driver", ""); err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	cases := []struct {
		name string
		call func() error
	}{
		{"insert user", func() error { _, err := AddUsersWithContext(ctx, &Users{}); return err }},
		{"read customer", func() error { _, err := GetCustomerByIdWithContext(ctx, 1); return err }},
		{"insert account", func() error { _, err := AddAccountsWithContext(ctx, &Accounts{}); return err }},
		{"read invite", func() error { _, err := GetUserInvitesByIdWithContext(ctx, 1); return err }},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if err := tc.call(); !errors.Is(err, context.Canceled) {
				t.Fatalf("want context canceled, got %v", err)
			}
		})
	}
}
