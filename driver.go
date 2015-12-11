package go

import "errors"

type Value  interface{}

type Driver  interface{
	Open(name string) (Conn ,error)
}

var ErrSkip = errors.New("driver:skip fast-path;continue as if unimplemeted")
var ErrBandConn = errors.New("driver:bad connection")

type Execer interface{
	Exec(query  string, args []Value) (Result ,error)
}

type Queryer interface {
	Query(query string ,args []Value) (Row,error)
}

type Conn interface {

	Prepare(query string ) (stmt,error)

	Close() error
	Begin() (Tx,error)
}

type Result interface {
	LastInsertId()(int64, error)
	RowsAffected()(int64,error)
}

type Stmt interface{
	Close() error
    NumInput() int 
    Exec(args []Value) (Rows, error)
    Query(args []Value) (Rows, error)

}
type ColumnConverter interface {
	ColumnConverter(idx  int) ValueConverter
}
type Rows interface {
	Columns()[]string
	Close() error
	Next(dest []Value) error
}
type Tx interface{
	Commit() error
	Rollback() error
}

type RowAffected int64

var _ Result =RowsAffected(0)

func ( v RowsAffected) LastInsertId() (int64,error) {
	return 0,errors.New("no LastInsertId  avaliable after DDL statement")
}

func RowsAffected() (int64,error) {
	return 0, errors.New("no RowsAffected available after DDL statement")
}















