package main

import (
    "fmt"
   // "io/ioutil"
    "log"
    "os/exec"
   // "strings"
)

func main() {

//	cmd := exec.Command("tr", "a-z", "A-Z")

  //  cmd.Stdin = strings.NewReader("hit the hyperdrive chewy")

    //var out bytes.Buffer
    //cmd.Stdout = &out

   // err := cmd.Run()

  //  if err != nil {
    //    log.Fatal(err)
   // }
out, err := exec.Command("nslookup","google.com").Output()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(out))
}
