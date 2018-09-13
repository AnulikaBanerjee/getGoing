/* decision making*/

package main

import . "fmt"
const VALUE int = 10	//constants are declared outside functions. It's good practice to name them in all-caps.
func main(){

		a,b:=2,2
		if a>b {
		Println("a is bigger")
		// Go is strict on newlines. else should be written right after the curly brace, otherwise it throws error. 
		}else if b>a {
		Println("b is bigger")
		}else {
		Println("a and b are equal")
		}
		
		switch a {
			case 1 : Println("a=1")		//break is not required
			case 3 : Println("a=3")
			default : Println("a is not 1 or 3")
		}
		
		var x interface{}
		switch i := a.(type){				//error: cannot type switch on non-interface value a (type int). Change "a" to "x"
			case nil : Println("Type of x is",i)
			case float64 : Println("float64")
			case int : Println("int")
		}
	
		/*there's a select statement for decision making on channels. */
}

