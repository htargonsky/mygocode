package main

import (
	"fmt"
)

type Astro struct {
	name string
	age int
	lastMission string
	isNeeded bool
}
func main() {

	astro1 := Astro{name: "billy the kid", age: 18}


	astro2  := Astro{name: "Sundance kid", age: 40}
//	astro3 := 

	fmt.Println(astro1)
    fmt.Println(astro2)
    //fmt.Println(ast3)
}
