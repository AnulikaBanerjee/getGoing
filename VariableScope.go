package main

import "fmt"

var a,b int = 20,30	// global variable

func main(){
	var a,c int= 10,50	//local variables
	d:=sum(a,c)		// Formal paramter
	/* in the above call, "a" is a formal parameter which means, 
	it's a parameter to a function which gives preference to local var over global var.
	Replace line 8 with [var c int = 10] and check the difference.*/
	fmt.Println(d)
}

func sum(a,b int) int{
	return a+b
}