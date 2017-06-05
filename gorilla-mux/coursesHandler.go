package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

// CoursesList is here
func CoursesList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is List of Courses")
}

// CreateCourse is here
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	courseID := uuid.NewV4()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Creating a course with UUID: %v\n", courseID)
}

// CourseDetail is here
func CourseDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID := vars["courseId"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Viewing course with ID: %v\n", courseID)
}

// UpdateCourse is here
func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID := vars["courseId"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated course with ID: %v\n", courseID)
}

// DeleteCourse is here
func DeleteCourse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	courseID := vars["courseId"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted course with ID: %v\n", courseID)
}
