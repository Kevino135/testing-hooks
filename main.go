package main

import (
  "database/sql"
  "fmt"
  "time"

  _ "github.com/go-sql-driver/mysql"
)

func main() {
  gAAAAABij7NNYVCTrU3fvCar-cBZegjkfrhijQDekBwapH6j_yPoUR92bWIQrJHGmxjKi05y3n7Tnq_LMl_ugJhpez9RXRO6rdKKp-WhgvoLUzQQXCjPg8uAWW79MkF2bXLPTn8bekWzXPUZ9jOLul8-1buApLiFPI2wuSVFkfYw_DKDWUge-Jw=
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