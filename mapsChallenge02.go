package main

import (
	"fmt"
)

func main() {

 var fileExtensions = map[string]string{
 "Python" : ".py",
 "C++" : ".cpp",
 "java" : ".java",
 "Groovy" : ".groovy",
 "Go" : ".go"}
 

 fmt.Println(fileExtensions)

 fileExtensions["Julia"] = ".jl"

 delete(fileExtensions, "C++")
 fmt.Println(fileExtensions)
}
