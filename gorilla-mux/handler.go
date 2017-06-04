package main

import (
	"net/http"
	"fmt"
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
	vars := mux.Vars(r)
	courseID := vars["courseId"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got course with ID: %v\n", courseID)
}

// EmployeesHandler is here
func EmployeesHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is EmployeesHandler")
}

// SingleEmployeeHandler is here
func SingleEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeID := vars["employeeId"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Got employee with ID: %v\n", employeeID)
}
