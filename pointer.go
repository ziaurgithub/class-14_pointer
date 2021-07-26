package main
import "fmt"

func new(a *int){
fmt.Println(a)
*a=*a+20
}


func main(){
var x int
var y *int

fmt.Println("x value is", x)
fmt.Println("x memory address is", &x)

fmt.Println("y value is", y)
fmt.Println("y memory address is", &y)

x=10
y=&x
fmt.Println("x is", x)
fmt.Println("y is", &y)

fmt.Println("y dereferencing value is", *y)

new(&x)
fmt.Println(x)
}

