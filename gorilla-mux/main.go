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
	r.Path("/status").
			HandlerFunc(HealthCheck).
			Methods("GET")

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
		Name("Create new course")
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
		Name("Create new employees")
	api.Path("/employees/{employeeId}").
		HandlerFunc(EmployeeDetail).
		Methods("GET").
		Name("Employee's detail")
	api.Path("/employees/{employeeId}").
		HandlerFunc(UpdateEmployee).
		Methods("PUT").
		Name("Update Employee's detail")
	api.Path("/employees/{employeeId}").
		HandlerFunc(DeleteEmployee).
		Methods("DELETE").
		Name("Delete Employee")

	http.Handle("/", r)

	err := http.ListenAndServe(":9090", nil)

	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
