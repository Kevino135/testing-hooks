package main

import (
  "database/sql"
  "fmt"
  "time"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  gAAAAABij8Ie0fsjCq6E6DcEXCG6jPXA8XEX3FJDqtIIMPDaVHxfcG1TBFW1mcA8w9jhDLIaiM1zAT7hM5x-WoBTxqL3cf5uKHpS46x0cOIpklASYf79qE0aQU-TelcoKz9oi8Xcnn0IcrFIsX45W2dhLVD_bR1YQlk6H3PsP4uoGtweXXfU7G8=
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