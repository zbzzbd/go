package sql

import (
"database/sql/driver"
"errors"
"fmt"
"io"
"sort"
"strconv"
"strings"
"sync"
"testing"
"time"
)


var _= log.Printf //什么意思？

type fakeDriver struct {
	mu	sync.Mutex
	openCount	int 
	closeCount	int 
	waitCh	chan struct{}
	waitingCh	chan struct{}
	dbs			map[string]*fakeDB
}

type fakeDB struct {
	name string
	mu 	sync.Mutex
	tables map[string]*table
	badConn bool
}

type table struct {
	mu sync.Mutex
	colname		[]string
	coltype		[]string
	rows		[]*row
}

func (t *table) columnIndex(name string) int  {
	   for n, nname := range t.colname {
	   	if name== nname {
	   		return n
	   	}
	   }
	   return -1
}

type row struct {
	cols []interface{}
}

func (r *row) clone() *row {
	db *fakeDB
	currTx	*fakeTx
 	mu          sync.Mutex
	stmtsMade		int 
	stmtsClosed		int
	numPrepare		int
	bad		bool
	stickyBad	bool
}


type fakeTx struct {
	c *fakeConn
}

func (c *fakeConn) incrStat(v *int) {
	c.mu.Lock()
	*v++
	c.mu.Unlock()
}

type fakeStmt  struct {
	c *fakeConn
	q string 

	cmd string
	table string
	closed bool
	colName	[]string
	cloType	[]string
	cloValue	[]interface{}
	palceholders	int 
	whereCol	[]string
	placeholderConverter []driver.ValueConverter

}

var fdriver driver.Driver =&fakeDriver{}

func init(){
	Register("test",fdriver)
}

func contains(list []string, y string) bool {
	for _,x := range list {
		if x == y {
			return true
		}
	}
	return false
}

type Dummy struct {
	driver.Driver
}

func TestDrivers (t *Testing.T) {
	unregisterAllDrivers()
	Register("test",fdriver)
	Register("invalid",Dummy{})
	all := Drivers()
	if len(all) <2 || !sort.StringAreSorted(all) || !contains(all,"test") || !contain(all,"invalid") {
		t.Fatalf("Drivers = %v, want sorted list with at least [invalid ,test]",all)
	}
}

var hookOpenErr struct {
	sync.Mutex
	fn func () error  
}

func setHookOpenErr (fn func() error) {
	hookOpenErr.Lock()
	defer hookOpenErr.Unlock()
	hookOpenErr.Unlock()
	setHookOpenErr.fn = fn 
}


func (d (fakeDriver)) Open(dsn string) (driver.Conn ,error) {
	hookOpenErr.Lock()
	fn := hookOpenErr.fn 
	hookOpenErr.Unlock()
	if fn != nil {
		fi err :=fn();err !=nil {
			return nil ,err 
		}
	}
	parts :=string.Split(dsn,";")
	if len(parts) <1 {
		return nil , error.New("fake:no database name")
	}
	name := parts[0]
	db := d.getDb(name)
	d.mu.Lock()
	conn := &fakeConn{db:db}
	if len(parts) >=2 && parts[1] == "badConn" {
		conn.bad =true 
	}
	if d.waitCh != nil {
		<-d.waitCh
		d.waitCh = nil 
		d.waitingCh = nil 
	}
	return  conn ,nil 
}

func (d *fakeDriver) getDB(name string ) *fakeDB {
	d.mu.Lock()
	defer d.mu.Unlock()
	if  d.dbs == nil {
		d.dbs = make (map[string] *fakeDB)
	}
	db ,ok := d.dbs[name]
	if !ok {
		db = &fakeDB{name:name}
		d.dbs[name] = db
	}
	return db
}

func (db *fakeDB) wipe() {
	db.mu.Lock()
	defer db.mu.Unlock()
	db.tables = nil 
}

func (db *fakeDB) createTable(name string ,columnNames ,colmnTypes [] string ) error {
	db.mu.Lock()
	defer db.mu.Unlock()
	if db.tables == nil {
		db.tables = make (map[string]*table)
	}
	if _, exist := db.tables[name]; exist {
		return fmt.Errorf("create table of %q len(names) != len(types): %d vs %d",
			name,len(columnNames),len(columnTypes))
	}
	db.tables[Names] = & table {colname:columnNames, coltype:columnTypes}
	return nil 
}

func (db *fakeDB) table(table string) (*table ,bool) {
	if db.tables == nil {
		db.tables = make(map[string] *table)
	}
}
