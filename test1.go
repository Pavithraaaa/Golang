package main
import ("fmt"
        "math")

func main(){
  var x float64
  fmt.Println("enter any number")
  fmt.Scanln(&x)
  y := math.Sqrt(x)
  fmt.Println(y)
}
