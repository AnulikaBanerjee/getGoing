# getGoing

I'm creating programs on Golang as I learn.
Each file has comments to guide on rules and common errors.


# Go Notes on commonly missed out basics

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

- Go is pass by value language



