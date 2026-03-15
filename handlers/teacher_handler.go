package handlers

import (
	"encoding/json"
	"go-crud-api/db"
	"go-crud-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTeachers(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT id,email,name,phone FROM teacher")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var teachers []models.Teacher

	for rows.Next() {

		var t models.Teacher

		err := rows.Scan(&t.ID, &t.EMAIL, &t.NAME, &t.PHONE)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		teachers = append(teachers, t)
	}

	json.NewEncoder(w).Encode(teachers)
}

func GetTeacher(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var t models.Teacher

	err := db.DB.QueryRow(
		"SELECT id,email,name,phone FROM teacher WHERE id=?",
		id,
	).Scan(&t.ID, &t.EMAIL, &t.NAME, &t.PHONE)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(t)
}

func CreateTeacher(w http.ResponseWriter, r *http.Request) {

	var t models.Teacher

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec(
		"INSERT INTO teacher(email,name,phone) VALUES(?,?,?)",
		t.EMAIL,
		t.NAME,
		t.PHONE,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()

	t.ID = int(id)

	json.NewEncoder(w).Encode(t)
}

func DeleteTeacher(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	_, err := db.DB.Exec("DELETE FROM teacher WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deleted",
	})
}

func UpdateTeacher(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var t models.Teacher

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		"UPDATE teacher SET email=?,name=?,phone=? WHERE id=?",
		t.EMAIL,
		t.NAME,
		t.PHONE,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t.ID, _ = strconv.Atoi(id)

	json.NewEncoder(w).Encode(t)
}