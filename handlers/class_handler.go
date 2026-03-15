package handlers

import (
	"encoding/json"
	"go-crud-api/db"
	"go-crud-api/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)


// GET ALL CLASSES

func GetClasses(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT id,name,teacher_id FROM class")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var classes []models.Class

	for rows.Next() {
		var c models.Class
		err := rows.Scan(&c.ID, &c.NAME, &c.TeacherID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		classes = append(classes, c)
	}

	json.NewEncoder(w).Encode(classes)
}


// GET CLASS BY ID

func GetClass(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	row := db.DB.QueryRow("SELECT id,name,teacher_id FROM class WHERE id=?", id)

	var c models.Class

	err := row.Scan(&c.ID, &c.NAME, &c.TeacherID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(c)
}


// CREATE CLASS

func CreateClass(w http.ResponseWriter, r *http.Request) {

	var c models.Class

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, err := db.DB.Exec(
		"INSERT INTO class(name,teacher_id) VALUES(?,?)",
		c.NAME,
		c.TeacherID,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	c.ID = int(id)

	json.NewEncoder(w).Encode(c)
}


// UPDATE CLASS

func UpdateClass(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	var c models.Class

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.DB.Exec(
		"UPDATE class SET name=?, teacher_id=? WHERE id=?",
		c.NAME,
		c.TeacherID,
		id,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	c.ID, _ = strconv.Atoi(id)

	json.NewEncoder(w).Encode(c)
}

// DELETE CLASS

func DeleteClass(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id := params["id"]

	_, err := db.DB.Exec("DELETE FROM class WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Deleted successfully",
	})
}