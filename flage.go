package main 
import  (
"fmt"

"path/filepath"
"os"
"errors"

)


func GetRunningFilePath() string {
	fmt.Printf("Dir--->%s\n%s\n%s\n", os.Args[0],filepath.Dir("."))
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	fmt.Printf("DIR---->%v\n", dir)
	if err != nil {
   errors.New("cuowu")
		return ""
	}

	return dir + "/"
}


 func main() {
	

 path:=GetRunningFilePath()
 

fmt.Println(path)
}
