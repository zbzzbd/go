package main 
import "fmt"
import "math"
 
 const  s string ="constant"
func main() {
	fmt.Println(s)
    const  n =500000000
    const d =3e20 /n
    fmt.Println(d)
    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))

/* 循环*/
	i:=1
		for i<=3{
			fmt.Println(i)
			i=i+1
        }

     for j:=7;j<=9; j++{
     	fmt.Println(j)
     }
     for{
     	fmt.Println("loop")
     	break
     }
//for配合range可以用于读取slice和map的数据：
      rating :=map[string]float32{"c":5,"go":4.5,"Python":4.5,"c++":2}

     for k,v:=range rating {
        fmt.Println(k,v)
         
     }

     
     //条件
    if 7%2 ==0 {
    	fmt.Println("7 is even")
    }else{
    	fmt.Println("7 is odd")
    }

    if 8%4==0{
        fmt.Println("8 is divisible by 4")
    	} 
    if num:=9; num<0{
    	fmt.Println(num,"is neagetive")
    }else if num <10{
    	fmt.Println(num,"has i digit")
    }else{
    	fmt.Println(num,"has  mutiple digits")
    }

//

}
