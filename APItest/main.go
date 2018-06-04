package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// var (
// 	ID       int
// 	name     string
// 	birthday string
// 	gender   string
// )

// func main() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", Index)
// 	router.HandleFunc("/todos", TodoIndex)
// 	router.HandleFunc("/todos/{todoId}", TodoShow)

// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

// func Index(w http.ResponseWriter, r *http.Request) {
// 	db, err := sql.Open("mysql", "michael:michael@tcp(10.91.4.96:3306)/YiChun?charset=utf8")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//fmt.Fprintln(w, "Welcome")
// 	rows, err := db.Query("select * from student")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&ID, &name, &birthday, &gender)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Fprintln(w, ID, name, birthday, gender)
// 	}
// }

// func TodoIndex(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Todo Index!")
// }

// func TodoShow(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	todoId := vars["todoId"]
// 	fmt.Fprintln(w, "Todo show: ", todoId)
// }

func main() {
	a := App{}
	a.Initialize("michael", "michael", "YiChun")
	a.Run(":8080")
}
