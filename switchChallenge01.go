/* RZFeeser | Alta3 Research
   switch - case and default  */

package main

import (
    "fmt"
    "strings"
//    "runtime"
)

func main() {

	var goVer string

	fmt.Println("enter version")
	fmt.Scanf("%s", &goVer)

           // init gover                
	   switch   {
    case strings.HasPrefix(goVer, "go1.17"):                 // if matches "go1.17", do the code below then STOP
        fmt.Println("You are using the latest version of GoLang")
    case strings.HasPrefix(goVer, "go1.16"), strings.HasPrefix(goVer,  "go1.15"):       // if matches "go1.16", "go1.16.5", OR "go1.15" 
        fmt.Println("You are using an older, but acceptable version of GoLang")
    default:                       // in all other cases run the code below
        fmt.Println("I cannot make a recommendation.")
    }
}
