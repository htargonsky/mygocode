/* Alta3 Research | RZFeeser
   Executing system commands  */

package main

import (
    "log"
    "os/exec"
    "fmt"
)

func main() {

    // prepares a "cmd" struct
    cmd := exec.Command("ls")

    out, err := cmd.Run()

    if err != nil {
        log.Fatal(err)
    } else {
	    fmt.Println(out)
    }
}

