package main

import (
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	a := App{}
	a.Initialize("michael", "michael", "YiChun")
	a.Run(":8080")
}
