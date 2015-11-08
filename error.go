package  main 
import "errors"
import "fmt"

//定义一个函数f1,返回2个参数类型分别为 int,error

func f1(arg int) (int ,error) {
	if arg ==42 {
		return -1,errors.New("can't work with 42")//初始化一个错误，errors.new() 来实现
	}
	return arg+3, nil
}

//定义一个结构体
type argError  struct {
	arg  int 
	prob  string
}

//定义一个结构体的方法为error, 返回string 类型
func (e *argError) Error() string{
	return fmt.Sprintf("%d-%s", e.arg,e.prob)
}

//定义一个函数f2,f返回参数类型为int,error
func f2(arg int ) (int,error) {
	if arg ==42{
		return -1,&argError{arg,"can't work with it "}//错误信息，直接初始化一个对象argerror
	}
	return  arg+3,nil
}

func main() {
	
   for _,i:= range []int{7,42}{
   	 if r,e :=f1(i);e!=nil {   //if表达式可以传递3个参数
   	 	fmt.Println("f1 failed",e)
   	 }else {
   	 	fmt.Println("f1 worked:", r)
   	 }
   }
  for _,i:=range []int{7,42}{
  	if r,e:=f2(i);e != nil {
        fmt.Println("f2 failed:",e)
  		}else {
         fmt.Println("f2 worked:",r)
  		}
  }
  _,e:=f2(42)
  if ae,ok:=e.(*argError);ok {
  	fmt.Println(ae.arg)
  	fmt.Println(ae.prob)
  }
   

}
