/* Alta3 Research | RZFeeser
   mySQL - Insert a row into a table within our database */
   
package main

import (
    "database/sql"
    "fmt"
    "log"
"time"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    db, err := sql.Open("mysql", "student:playstation5@tcp(127.0.0.1:3306)/testdb")
    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    // this is the data we wish to insert into our program
    sql := "INSERT INTO cities(name, population) VALUES ('Moscow', 12506000)"

    res, err := db.Exec(sql)
    if err != nil {
        panic(err.Error())
    }
    lastId1, err := res.LastInsertId()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("The last inserted row id: %d\n", lastId1)

    time.Sleep(2 * time.Minute)
    sql = "INSERT INTO cities(name, population) VALUES ('Hartford', 506000)"
    res2, err := db.Exec(sql)

    if err != nil {
        panic(err.Error())
    }

    // With LastInsertId, we get the last inserted id
    lastId, err := res2.LastInsertId()

    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("The last inserted row id: %d\n", lastId)
}