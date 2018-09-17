package main

import "fmt"

//structure definition
type Name struct{
	firstName,lastName string
	salary int
}

func (name Name) fullName() string {
	return fmt.Sprintf("%s %s",name.firstName,name.lastName)
}

func main(){
	name := Name{
		firstName:"Annie",
		lastName:"Paul",
		salary:1200
		}	// braces should be closed on the previous line itself, else we get syntax error

	// Other ways of assigning values to struct type	
	name1:=Name{"Daisy","Vaught",1100}
	var name2 Name
	name2.firstName="Rhythm"
	name2.lastName="Blake"
	fmt.Println(name.fullName())
	fmt.Println(name1.fullName())
	fmt.Println(name2.fullName())

}