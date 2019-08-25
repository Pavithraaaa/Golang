package main
import "fmt"

func main(){
var p int
fmt.Println("enter a number")
fmt.Scanln(&p)
for i := 1 ; i <= 10 ; i ++{
  fmt.Printf("%v*%v=%v\n",p,i,p*i)

}
}
