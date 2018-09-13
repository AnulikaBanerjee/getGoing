/* variable declaration and initialization*/

package main

import . "fmt"
const VALUE int = 10	//constants are declared outside functions. It's good practice to name them in all-caps.
func main(){
	var a int = 3
	var d int	//have to declare dataType else errors out
	d=5
	var c = 4	// automatic dataType assigned as int
	b := 1		// := is used for short variable declarations+assignment. This is same as var b int = 1
	Println(a+b+c+d+VALUE)

}

