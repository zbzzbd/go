package orm

import (
   "flag"
   "fmt"
   "os"
   "strings"
)


type commander  interface {
	Parse([]string)
	Run()error
}


var (
  commands = make(map[string]commander)
)

func printHelp(errs ...string) {
	
	content := `orm command usage:`
	if len(errs)>0 {
		fmt.println(errs[0])
	}
	fmt.Println(content)
	os.Exit(2)
}
func RunCommand() {
	if len(os.Args)<2 || os.Args[1] != "orm" {
		return 
	}
	BootStrap()


args := argString(os,Args[2:])
name :=args.Get(0)

if name == "help" {
	printHelp()
}

if cmd , ok:=commands[name];ok {
	cmd.Parse(os.Args[3:])
	cmd.Run()
	os.Exit(0)
}else {
	if name ==""{
		printHelp()
	}else {
		printHelp(fmt.Sprintf("unkonown command %s",name))
	}
}
}



type commandSyncDb struct {
	al  * alias
	force  bool
	verbose  bool
	noInfo   bool
	rtOnError  bool
}

func (d *commandSyncDb)Parse(args  []string) {
	var name  string 
	flageSet :=flage.NewFlageSet("orm command:syncdb",flag.ExitOnError)
	flageSet.StringVar(&name,"db","default","DataBase  alias name")
	flageSet.BoolVar(&d.force, "force",false,"drop tables before create")
	flageSet.BoolVar(&d.verbose,"v", false,"verbose info")
	flageSet.Parse(args)
	d.al =getDbAlias(name)
}

func (d  *commandSyncDb)Run() error {
	var drops []string
	if  d.force {
		drops = getDbDropSql(d.al)

	}

	db:=d.al.DB

	if d.force {

		for i,mi :=range modelCache.allordered(){
			query :=drops[i]
			if !d.noInfo {
				fmt.Printf("drop table `%s` \n",mi.table)
			}
			_,err:=db.Exec(query)
			if !d.noInfo {

				fmt.Printf("%s\n",err.Error())
			}





		}
	}
	sqls ,indexs :=getDbCreateSql(d.al)
	tables, err: = d.al.DbBaser.GetTables(db)
	if err != nil {
		if d.rtOnError {
			return err
		}
		fmt.printf("%s\n",err.Error())
	}

	for _,fi :=range fields{
		query :=getColumnAddQuery (d.al,fi)
		if !d.noInfo {
			fmt.Printf("add column `%s`for table `%s` \n",fi.fullName,mi.table)
		}
		_,err :=dbExec(query)
		if err !=nil {
			if d.rtOnError{
				return err 
			}
			fmt.Printf("%s\n",err.Error())
		}
	}

	for _, idx :range indexes[mi.table] {
		if d.al.DbBaser.IndexExists(db,idx.Table,idx.Name) ==false {
			if !d.noInfo{
				fmt.Printf("create index `%s` for table `%s` \n",idx.Name,idx.Table)
			}
			query :=idx.Sql
			_,err :=db.Exec(query)
			if d.verbose {
				fmt.Printf("%s\n",query)
			}
			if err!=nil {
				if d.rtOnError {
					return err
				}
				fmt.Printf("",err.Error())
			}
		}
	}
	continue
}

if  !d.noInfo {
	fmt.Printf("")
}
｝
return nil 
}


type commandSqlAll struct {
	al *alias
}


func (d  *commandSqlAll)Parse(args []string) {
	var name string 
	flageSet :=flage.NewFlagSet("orm command :sqlall",flage.ExitOnError)
	flageSet.StringVar(&name,"db","default","DataBase alias name")
	flageSet.Parse (args)
	d.al =getDbAlias(name)
}

func (d *commandSqlAll)Run() error {
	sqls ,indexed :=getDbCreateSql(d.al)
	var all [] string 
	for i,mi :=range modelCache.allordered(){
		
        queries :=[]string{sqls[i]}
        for _,idx:=range indexes[mi.table]{
           queries = append(queries,idx.Sql)
        }
	sql :=strings.Join (queries,"\n")
	all =append(all,sql)
	}
	fmt.println(strings.Join(all,"\n\n"))
	return nil 
}

func init() {
	commands["syncdb"] = new (commandSyncDb)
	commands["sqlall"] = new (commandSqlAll)
}


func RunSyncdb(name string ,force bool ,verbose bool) error {
	BootStrap()
	al :=getDbAlias(name)
	cmd := new (commandSyncDb)
	cmd.al=al
	cmd.force=force
	cmd.noInfo=!verbose
	cmd.verbose=verbose
	cdm.rtOnError =true
	return cmd.Run()
}






