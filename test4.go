package main
import ("fmt")
func main(){
  var y int
  fmt.Println("enter any number")
  fmt.Scanln(&y)
  for i := 2; i < y; i++ {
    if (y%i == 0) {
    fmt.Println("The given number is non-prime")
    break
    } else {
    fmt.Println("The given number is prime")
    break
  }
}
}
