package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//fmt.Println("HelloWorld")

	r := mux.NewRouter()

	//HEALTH CHECK API
	r.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	api := r.PathPrefix("/api/v1").Subrouter()

	//HOME DIRECTORY API
	api.Path("/").
		HandlerFunc(HomeHandler).
		Methods("GET").
		Name("API V1 Home")

	//COURSES API
	api.Path("/courses").
		HandlerFunc(CoursesList).
		Methods("GET").
		Name("List of courses")
	api.Path("/courses").
		HandlerFunc(CreateCourse).
		Methods("POST").
		Name("Create a course")
	api.Path("/courses/{courseId}").
		HandlerFunc(CourseDetail).
		Methods("GET").
		Name("Course's detail")
	api.Path("/courses/{courseId}").
		HandlerFunc(UpdateCourse).
		Methods("PUT").
		Name("Update Course's detail")
	api.Path("/courses/{courseId}").
		HandlerFunc(DeleteCourse).
		Methods("DELETE").
		Name("Delete Course")

	//EMPLOYEES API
	api.Path("/employees").
		HandlerFunc(EmployeesList).
		Methods("GET").
		Name("List of employees")
	api.Path("/employees").
		HandlerFunc(CreateEmployee).
		Methods("POST").
		Name("List of employees")
	api.Path("/employees/{employeeId}").
		HandlerFunc(EmployeeDetail).
		Methods("GET").
		Name("Employee's detail")
	api.Path("/employees/{employeeId}").
		HandlerFunc(UpdateEmployee).
		Methods("PUT").
		Name("Employee's detail")
	api.Path("/employees/{employeeId}").
		HandlerFunc(DeleteEmployee).
		Methods("DELETE").
		Name("Employee's detail")

	http.Handle("/", r)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
