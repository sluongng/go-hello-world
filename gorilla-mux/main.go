package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// HomeHandler Simple Testing home function
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is HOME")
}

// CoursesHandler is here
func CoursesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is Many Courses")
}

// SingleCourseHandler is here
func SingleCourseHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is SingleCourseHandler")
}

// EmployeesHandler is here
func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is EmployeesHandler")
}

// SingleEmployeeHandler is here
func SingleEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is SingleEmployeeHandler")
}

func main() {
	//fmt.Println("HelloWorld")

	r := mux.NewRouter()

	//HEALTH CHECK API
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	api := r.PathPrefix("/api/v1").Subrouter()

	//HOME DIRECTORY API
	api.HandleFunc("/", HomeHandler).
		Methods("GET").
		Name("API V1 Home")

	//COURSES API
	api.HandleFunc("/courses", CoursesHandler).
		Methods("GET", "POST"). // TODO: POST
		Name("List of courses")
	api.HandleFunc("/courses/{courseId}", SingleCourseHandler).
		Methods("GET", "PUT", "DELETE"). // TODO: PUT and DELETE
		Name("Course's detail")

	//EMPLOYEES API
	api.HandleFunc("/employees", EmployeesHandler).
		Methods("GET", "POST"). // TODO: POST
		Name("List of employees")
	api.HandleFunc("/employees/{employeeId}", SingleEmployeeHandler).
		Methods("GET", "PUT", "DELETE"). // TODO: PUT and DELETE
		Name("Employee's detail")

	http.Handle("/", r)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
