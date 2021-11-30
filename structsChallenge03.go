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
// this is our new struct
type nasaMission struct {
        people []Astro
        number int
        message string
}
func main() {

	astro1 := Astro{name: "billy the kid", age: 18, lastMission: "sante Fe train robbery"}


	astro2  := Astro{name: "Sundance kid", age: 40, isNeeded: true, lastMission: "Bolivia"}
//	astro3 := 

	fmt.Println(astro1)
    fmt.Println(astro2)
    //fmt.Println(ast3)
	astros := []Astro{astro1, astro2}

	// display the slice
    fmt.Println(astros)

    // select data from the struct
    fmt.Println(astros[1].lastMission)

    mission1 := nasaMission{people: astros, number: len(astros)}

    fmt.Println(mission1)

    fmt.Printf("%+v", mission1)
}
