package main

import (
  "database/sql"
  "fmt"
  "time"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  gAAAAABij7XQ15C1jWTkHDNH8_nvain8n4xK6bgFwQYce4nnwvf6mycvW8OrPo7BE39PIb1OTE7hUE85fUnOzUJZUyIb0_zaQDP6Vs7vOoEDdZ3GPu4XjlIRSgizJfDw4RlRm2nany8ptakwJBJwpgq4dgUZ6oMuSKeJcGZkawmSDrwbmLZhNxQ=
  db.SetMaxOpenConns(15)
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