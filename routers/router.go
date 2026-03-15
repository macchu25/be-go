package routers

import (	"go-crud-api/handlers"
"github.com/gorilla/mux")

func SetUpRouter() *mux.Router{
router := mux.NewRouter()




router.HandleFunc("/student",handlers.GetStudents).Methods("GET")
router.HandleFunc("/student/{id}",handlers.GetStudent).Methods("GET")
router.HandleFunc("/student",handlers.CreateStudent).Methods("POST")
router.HandleFunc("/student/{id}",handlers.DeleteStudent).Methods("DELETE")
router.HandleFunc("/student/{id}",handlers.UpdateStudent).Methods("PUT")

router.HandleFunc("/teacher",handlers.GetTeachers).Methods("GET")
router.HandleFunc("/teacher/{id}",handlers.GetTeacher).Methods("GET")
router.HandleFunc("/teacher",handlers.CreateTeacher).Methods("POST")
router.HandleFunc("/teacher/{id}",handlers.DeleteTeacher).Methods("DELETE")
router.HandleFunc("/teacher/{id}",handlers.UpdateTeacher).Methods("PUT")

router.HandleFunc("/class",handlers.GetClasses).Methods("GET")
router.HandleFunc("/class/{id}",handlers.GetClass).Methods("GET")
router.HandleFunc("/class",handlers.CreateClass).Methods("POST")
router.HandleFunc("/class/{id}",handlers.DeleteClass).Methods("DELETE")
router.HandleFunc("/class/{id}",handlers.UpdateClass).Methods("PUT")

return router
}
