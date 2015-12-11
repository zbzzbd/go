package driver

import (
  "fmt"
   "reflect"
   "strconv"
   "time"
)

type ValueConverter interface{
	//convertValue converts a value to a driver value
	ConvertValue(v interface{}) (Value,error)

}

type Valuer interface {
	//Value returns a driver Value
	Value() (Value,error)
}

/*
  Bool is a ValueConverter that converts input values to bools

  the  conversion rules are: boolean art returned unchanged
*/
var Bool boolType
type boolType struct {}
var _ ValueConverter =boolType{} // _这里也表示忽略么？


func (boolType)String() string {   //方法中未曾使用receiver 参数，可以省略
	return "Bool"
}

func  (boolType) ConverValue(src interface{}) (value,error) {
	
	switch s:= src.(type){   //src.(type) 这句话什么意思？其中的type又代表什么意思？
		case bool: return s,nil
		case string: 
		b,err :=strconv.ParseBool(s)
        if err !=nil{
        	return nil ,fmt.Errorf("sql/driver: couldn't convert %q into type bool", s)
        }
        return b,nil
    case [] byte:
    	b,err := strconv.ParseBool(string(s))
    	if err !=nil{
    		return nil , fmt.Errorf("sql/driver:couldn't vonvert %q into type bool ", s)
    	}
    	return b,nil
	}

	sv:=reflect.ValueOf(src)
	switch sv.Kind() {
	case reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32,reflect.Int64:
		uv :=sv.Int()
		if uv ==1 || uv ==0 {
			return uv==1,nil
		} 
		return nil,fmt.Errorf("sql/driver:couldn't convert %v (%T)", src,src)

	}
}
