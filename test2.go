package main
import ("fmt")
func evenodd(x int)  {
  if ( x % 2 == 0)  {
    fmt.Println("The number is even")
  } else{
    fmt.Println("The number is odd")
}
}

func main(){
  var p int
  fmt.Println("enter a number")
  fmt.Scanln(&p)
  evenodd(p)
}
