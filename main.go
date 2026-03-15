// package main

// import (
// 	"database/sql"
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"strconv"

// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/gorilla/mux"
// )

// type User struct {
// 	ID    int    `json:"id"`
// 	NAME  string `json:"name"`
// 	EMAIL string `json:"email"`
// 	PHONE string `json:"phone"`
// }

// var db *sql.DB
// var err error
// func enableCORS(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

// 		if r.Method == "OPTIONS" {
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
// func main(){
//     db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mydb")
// 	if err !=nil{
// 		log.Fatal(err)
// 	}

// 	fmt.Println("Kết nối database thành công!")
// router := mux.NewRouter()
// router.HandleFunc("/users",getUsers).Methods("GET")
// router.HandleFunc("/users/{id}",getUser).Methods("GET")
// router.HandleFunc("/users",createUser).Methods("POST")
// router.HandleFunc("/users/{id}",deleteUser).Methods("DELETE")
// router.HandleFunc("/users/{id}",updateUser).Methods("PUT")
// log.Println("SERVER OK 8000")
// log.Fatal(http.ListenAndServe(":8000", enableCORS(router)))
// }

// // hàm lấy user trong db
// func getUsers(w http.ResponseWriter, r *http.Request){
// 	rows, _ := db.Query("SELECT * FROM users")
// 	var  users []User
// 	for rows.Next(){
// 		var u User
// rows.Scan(&u.ID,&u.NAME,&u.EMAIL,&u.PHONE)
// 		users=append(users,u)

// 	}
// 	fmt.Println(users)
// 	json.NewEncoder(w).Encode(users)
// }
// // hàm lấy user theo id
// func getUser(w http.ResponseWriter, r *http.Request){
// 	params := mux.Vars(r)
// 	id := params["id"]
// 	rows := db.QueryRow("SELECT * FROM users WHERE id = ?", id )

// 		var u User
// rows.Scan(&u.ID,&u.NAME,&u.EMAIL,&u.PHONE)

// if err != nil {
// 		http.Error(w, "User not found", 404)
// 		return
// 	}

// fmt.Printf("User ID: %d Name: %s Email: %s Phone: %s\n",
// 	u.ID, u.NAME, u.EMAIL, u.PHONE)
// 	json.NewEncoder(w).Encode(u)
// }
// // hàm thêm user
// func createUser(w http.ResponseWriter, r *http.Request){
// 	var u User
// 	json.NewDecoder(r.Body).Decode(&u) // đọc dữ liệu
// 	res, _ := db.Exec("INSERT INTO users(name,email,phone) VALUES(?,?,?)",u.NAME,u.EMAIL,u.PHONE)
// 	id, _ := res.LastInsertId() // lấy id vừa tạo
// 	u.ID = int(id)
// 	json.NewEncoder(w).Encode(u)
// }
// func deleteUser(w http.ResponseWriter, r *http.Request) {
//     params := mux.Vars(r)
//     id := params["id"]
//     db.Exec("DELETE FROM users WHERE id=?", id)
//     json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
// }
// func updateUser(w http.ResponseWriter, r *http.Request) {
//     params := mux.Vars(r)
//     id := params["id"]
//     var u User
//     json.NewDecoder(r.Body).Decode(&u)
//     db.Exec("UPDATE users SET name=?, email=?, phone=? WHERE id=?", u.NAME, u.EMAIL, u.PHONE, id)
//     u.ID, _ = strconv.Atoi(id)
//     json.NewEncoder(w).Encode(u)
// }
                                                                         
package main

import (
	"fmt"
	"go-crud-api/db"
	"go-crud-api/routers"
	"net/http"


)
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "https://crud-horizon-lab.vercel.app")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}
func main() {
	db.Connect()

	router := routers.SetUpRouter()



	fmt.Println("Server running on port 8080")

http.ListenAndServe(":8080", enableCORS(router))
}

