/* RZFeeser | Alta3 Research
   errors.New constructs a basic error value with given message */

package main

import (
    "errors" // preview of using errors
    "fmt"
)

//name of funct   //input                      //output
func rollchar(firstName string, lastName string) (string, error) {
    if lastName == "Turnip" || lastName == "Radish" || lastName == "Potato" {
        return firstName + " the " + lastName, errors.New("Vegetables are not suitable last names for heroes.")
    }
    return firstName + " the " + lastName, nil
}

func main() {
    fmt.Println("Welcome to the Character Generator")

    var name string
    var descriptor string

    fmt.Print("enter name: ")
    fmt.Scanf("%s", &name)
    fmt.Print("enter descriptor: ")
    fmt.Scanf("%s", &descriptor)

    playerChar, err := rollchar(name, descriptor)


    if err != nil {
        fmt.Println("Error while spawning your requested character.")
    } else {
        fmt.Println(playerChar, "has been generated.")
    }
}
