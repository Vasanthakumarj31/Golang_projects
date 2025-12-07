package main  
import "fmt"
func main(){
  for{
    fmt.Println("This  is infinite loop")
    var a int
    fmt.Scan(&a)
    if(a == 1){
      break
    }
  }
}
