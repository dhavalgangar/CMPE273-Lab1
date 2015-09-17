package main

import "fmt"

func fibbo(num float64) float64 {

   if num == 0.0 {

	return 0.0

   }else if num == 1.0 {

	return 1.0
   }

   return fibbo(num -1.0) + fibbo(num -2.0)

}


func main(){

fmt.Print("Enter the number till which you want fibbonacci series :")

var num float64
fmt.Scanf("%f", &num)

fmt.Print("fib(")
fmt.Print(num)
fmt.Print(") = ")

fmt.Print(fibbo(num))

}

