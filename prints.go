package main

import "fmt"

func main() {
    var hello_paco string = "Hello, paco!\n"
    fmt.Printf(hello_paco)
    fmt.Println("1 + 2 =", 3); //Pinta las variables que le pases con un espacio en medio
    fmt.Println(len("cinco"))
    var paco = "Paco"
    fmt.Println(paco + " es un robot")
    var pi float64 = 3.14159
    fmt.Printf("%f\n", pi)
    fmt.Printf("%.2f\n", pi)
    fmt.Printf("%T\n", pi); //Pinta el tipo de la variable
    fmt.Println("true && false = ", true && false)
    fmt.Println("true || false = ", true || false)
    fmt.Println("!true =", !true)
}
