# getGoing

I'm creating programs on Golang as I learn.
Each file has comments to guide on rules and common errors.


# Go Notes and Basics

- Package named main with a main function is the only executable package in Go. Others are helper package.
```
$ go build main.go // creates a main.exe file in the same path.  ./main.exe can be used to run it.

$go run main.go helper.go // it’s valid to run main with other dependent file
```
- Variables can not be initialized outside func

- Variables can not be initialized twice

- Go won’t tolerate anything unused in the code, like unused imports or variables.

 - Go slices can be initialized in 3 ways:

	type person struct{
		name string
		age int
	}

     a)  person1:= person{“Amy”,26}
    b)  person1:= person{“name”:“Amy” , “age” :26}
    c)  person1.name=”Amy”
        person1.age=21

- If a function works on a pointer to a type, we can pass the type directly (not the pointer to it) and go will take the pointer to it by itself. 

- When we create a slice, internally go creates an array with the values and a data structure to store details of the struct which are: pointer to head, capacity and length. Capacity is always 2^n size as it grows or shrinks. Length is actual count of values held.

- Go is a “pass by value” language. Which means unless pointers come into play, everything in go is passed by value i.e. only a copy is passed.

- If we modify a struct via a function by passing the struct (not pointer), it won’t work as it’ll be a pass by value and update works on the copy of struct. On the contrary, modification on a slice via a function would work although it’s still a pass by value, but what is being passed is not the array, but the data structure of the slice.

- This kind of behavior is what separates types in Go as value types and reference types. Value types are int, float, bool, string, struct. Reference types are slice, map, channel, pointer and func.

### Commonly used Formatting verbs


| type/verb | verb/type |
| ------ | ------ |
|bool	|	%t|
|int	|	%d|
|string	|	%s|
|float	|	%g|
|chan	|	%p|
|pointer	|	%p|
|%v	|	value in default format|
|%+v	|	the + adds field names for structs|
|%T	|	Go syntax representation of type of the value|
