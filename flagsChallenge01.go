
package main

import (
    "bufio"
    "flag"
    "fmt"
    "io"
    "os"
)

func main() {

    var count int
    var fileName string
    flag.IntVar(&count, "n", 5, "number of lines to read from the file")
    flag.StringVar(&fileName, "fileName", "", "file to read")
    flag.Parse()

    var in io.Reader
    if  fileName != "" {
        f, err := os.Open(fileName)
        
        // error handling
        if err != nil {
            fmt.Println("error opening file: err:", err)
            os.Exit(1)
        }
        defer f.Close()   // still close the file if an error occurs

        in = f
    } else {
	    panic("No filename passed")
    }


    // read lines from the io.Reader var in
    buf := bufio.NewScanner(in)

    // iterate up to the value of count user a "for loop" 
    for i := 0; i < count; i++ {
        if !buf.Scan() {            // if scanning produces a false
            break                   // because we were asked to scan more lines
        }                           // than actually exist 
        fmt.Println(buf.Text())
    }

    if err := buf.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error reading: err:", err)
    }
}
