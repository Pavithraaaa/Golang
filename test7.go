package main
import "fmt"
func main(){
 var p string
 v,c := 0,0

 fmt.Println("Enter any string")
 fmt.Scanln(&p)
 for i :=0 ; i < len(p) ; i++ {

 if p[i] == 'a' || p[i] == 'e' || p[i] == 'i' || p[i] == 'o' || p[i] == 'u' {
   v += 1

 } else {
   c += 1
 }

}
fmt.Printf("Vowels=%v\n",v)
fmt.Printf("Consonants=%v\n",c)
}
