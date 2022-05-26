package main

import (
  "database/sql"
  "fmt"
  "time"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  gAAAAABij7Bl33hlNTkYtHnwUgUWpXeQ3mNgNoSv20NN-ywZ5do1Ct8aK9gz274uiBkBqLO4g54HZ41A_zDDJ-Pr3mplB6tFYFuNIEYQVjFeBvttP8F7PqlZ1LBunPkCBFxn9UYoWU4e7JNbVa-9BHUx6aeWnR2EMEUZ7-zaeUJaCeyblNT1Cgc=
  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(5)
  //Connection database query
  for i := 0; i < 100; i++ {
    go func(i int) {
      mSql := "select * from user"
      rows, _ := db.Query(mSql)
      rows.Close () // here, if you do not release the connection to the pool, other concurrencies will block after five runs
      fmt.Println (I)
    }(i)

  }

  for {
    time.Sleep(time.Second)
  }
}