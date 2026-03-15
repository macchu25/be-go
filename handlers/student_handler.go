package handlers

import (
	"encoding/json"

	"go-crud-api/db"
	"go-crud-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// hàm lấy user trong db
func GetStudents(w http.ResponseWriter, r *http.Request) {
	rows, _ := db.DB.Query("SELECT id,name,email,class_id FROM student")
	var student []models.Student
	for rows.Next() {
		var u models.Student
		rows.Scan(&u.ID, &u.NAME,  &u.EMAIL,&u.ClassID)
		student = append(student, u)

	}

	json.NewEncoder(w).Encode(student)
}

// hàm lấy teac theo id
func GetStudent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	rows := db.DB.QueryRow("SELECT  id,name,email,class_id FROM student WHERE id = ?", id)

	var u models.Student
	rows.Scan(&u.ID, &u.NAME,  &u.EMAIL,&u.ClassID)
}
// hàm thêm user 
func CreateStudent(w http.ResponseWriter, r *http.Request){
	var u models.Student
	json.NewDecoder(r.Body).Decode(&u) // đọc dữ liệu 
	res, _ := db.DB.Exec("INSERT INTO student(name,email,class_id) VALUES(?,?,?)",u.NAME,u.EMAIL,u.ClassID)
	id, _ := res.LastInsertId() // lấy id vừa tạo
	u.ID = int(id)
	json.NewEncoder(w).Encode(u)
}
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]
    db.DB.Exec("DELETE FROM student WHERE id=?", id)
    json.NewEncoder(w).Encode(map[string]string{"message": "Deleted"})
}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]
    var u models.Student
    json.NewDecoder(r.Body).Decode(&u)
    // Cập nhật đúng các cột: name, email, class_id
    db.DB.Exec("UPDATE student SET name=?, email=?, class_id=? WHERE id=?", u.NAME, u.EMAIL, u.ClassID, id)
    u.ID, _ = strconv.Atoi(id)
    json.NewEncoder(w).Encode(u)
}

