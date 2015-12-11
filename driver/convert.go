package driver

import (
"errors"
"fmt"
"reflect"
"strconv"
)

var errNilPtr = errors.New("destination pointer is nil")




func driverArgs(ds *driverStmt, args[] interface{}) ([]driver.Value,error) {
	dargs :=make([]driver.Value,len(args))
	var si driver.Stmt
	if ds != nil {
		di =ds.si
	}
	cc ,ok :=si.(driver.ColumnConverter)  //这句话什么意思？ 
}
