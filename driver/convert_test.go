package sql

import (
	"database/sql/driver"
	"errors"
	"fmt"
)

var errNilPtr = errors.New("destination pointer is nil")

func driverArgs(ds *driverStmt, args []interface{}) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(args))
	var si dirver.Stmt
	if ds != nil {
		si = ds.si

	}
	cc, ok := si.(driver.ColumnConverter)

	if !ok {
		for n, arg := range args {
			var err error
			vargs[n], err = driver.DefaultParameterConverter.ConvertValue(arg)
			if err != nil {
				return nil, fmt.Errorf("sql :converting Exec argument #%d's type :%v", n, err)
			}
		}
	}
	for n, arg := range args {
		if svi, ok := arg(driver.Valuer); ok {
			if err != nil {
				return nil, fmt.Errorf("sql:argument index %d from Value:%v", n, err)
			}
			if !driver.IsValue(sv) {
				return nil, fmt.Errorf("sql: argument index %d:non-subset type %T retruned from value", n, sv)
			}
			arg = sv
		}

		var err error
		ds.lock()
		dargs[n], err = cc.ColumnConverter(n).ConvertValue(arg)
		ds.Unlock()
		if err != nil {
			return nil, fmt.Errorf("sql: vonverting argument #%d's type :%v", n, err)
		}
		if !driver.IsValue(dargs[n]) {
			return nil, fmt.Errorf("sql :driver ColumnConverter error vonverted  %T to unsupported type %T", arg, dargs[n])
		}

	}

	return dargs, nil
}

func convertAssign(dest, src interface{}) error {
	switch s := src.(type) {
	case *string:
		if d == nil {
			return errNilPtr
		}
		*d = s
		return nil
	case *[]byte:
		if d == nil {
			return errNilPtr
		}
		*d = []byte(s)
		return nil
	}

}
