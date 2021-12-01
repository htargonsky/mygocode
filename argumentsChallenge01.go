/* Alta3 Research | RZFeeser
   CHALLENGE 01 - Display the number of arguments passed to the program */
   
package main
 
import (
    "fmt"
    "os"
)
 
func main() {

    
fmt.Printf("Arg length is %d", len(os.Args[1:]) )
}
