package main
import ("fmt")
func main(){
  fmt.Println("Enter any number between 1 to 10")
  var i int
  fmt.Scanln(&i)
  switch i {
case 1: fmt.Println("One")
case 2: fmt.Println("Two")
case 3: fmt.Println("Three")
case 4: fmt.Println("Four")
case 5: fmt.Println("Five")
case 6: fmt.Println("Six")
case 7: fmt.Println("Seven")
case 8: fmt.Println("Eight")
case 9: fmt.Println("Nine")
case 10: fmt.Println("Ten")
default:
   fmt.Println("Not supported")
}
}
