package main

import (
	"fmt"
	"math"		
	)

func main(){
	
	/* https://golang.org/pkg/math/ visit for all math functions */
	var a,b float64=2,3		// The dataType of a and b automatically gets assigned as "int"
	var c=math.Max(a,b)	//this will throw error as Max expects float64 arguments. Change a,b definition to float64 for running it.
	fmt.Println(c)
}