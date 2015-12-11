package orm

import (
"database/sql"
"errors"
"fmt"
"reflect"
"strings"
"time"
)

const (

format_Date ="2006-01-02"
format_DateTime = "2006-01-02  15:04:05"
)

var (

  operators = map[string]bool {

  	"exact":   true
  	"iexact":  true
    "contains":   true
    "icontains" :    true
    "gt":   true
    "gte" :   true
    "lt"  :  true
    "lte"  :true
    "eq"  : true
    "nq"  :true
    "startswith"  :true
    "endwith" : true
    "istartswith" :true
    "in" : true
    "between" :true
    "isnull"  : true
  }

)

var (
 ErrMissPK =errors.New("missed pk value")
)

 //check dbBase  implements dbBaser interface 
var _ dbBaser = new (dbBase)

func (d *dbBase)collectValues(mi  *modelInfo, ind  reflect.Value, cols  []string ,skipAuto bool ,insert bool ,names *[]string ,tz  *time.Location)  (values []interface{}) {

	var columns []string

	if names != nil {

		columns=*names
	}
    for _,column := range cols {
       var fi *fieldInfo
       if fi,_ = mi.fields.GetByAny(column); fi !=nil {
       	column= fi.column
       }else {
       	panic(fmt.Errorf("wrong db field/colunm name `%s` for model `%s`",column ,mi.fullName))//使用panic抛出异常
       }
        if fi.dbcol ==false || fi.auto && skipAuto {
        	continue
        }
        value,err := d.collectFieldValues(mi,fi,ind ,insert,tz)
        if err != nil {
        	return nil ,err
        }

        if names !=nil {
        	columns = append(columns,column)
        }
        values = append(values,value)

    }
    if names !=nil {
    	*names =columns
    }
    return
}

func collectFieldValue(mi *modelInfo, fi *fieldInfo, ind reflect.Value,insert bool , tz * time.Location) (interface{},error) {

 var value interface {}
 if fi.pk {
   _,value,_ =getExistpk(mi,ind)
 }else {
 	field := ind.Field(fi.fieldIndex)
 	if fi.isFeilder {
 		f := field.Addr().Interface().(Fielder)
 		value = f.RawValue() 
 	}else {

         switch  fi.fieldType{
         case  TypeBooleanField :
         	if nb,ok:= field.Interface().(sql.NullBool);ok {
         		value = nil 
         		if nb.Valid {
         			value = nb.Bool
         		}
         	}else if field.Kind() == reflect.Ptr {
         		if field.IsNil (){
         			value = nil 
         		}else {
         			value =field.Elem().Bool()
         		}
         	}else {
         		value = field.Bool()
         	}
        case  Typecharfield,TypeTextField :
        	if ns ,ok = field.interface().(sql.NullString);ok {   //这是什么用法？
             
        	}


         }

 	}
 }
	
}