package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/satori/uuid"
)

// Course struct: Exported fields must capitalize first letter
type Course struct {
	ID        uuid.UUID	`json:"Id"`
	Name      string	`json:"Name"`
	StartTime time.Time	`json:"StartTime"`
	EndTime   time.Time	`json:"ExpireTime"`
}

// Courses struct
type Courses map[string]*Course

// CoursesList is here
func CoursesList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is List of Courses")
}

// CreateCourse is here
func CreateCourse(w http.ResponseWriter, r *http.Request) {
	newCourse := Course{
		uuid.NewV4(),
		"",
		time.Now().UTC(),
		time.Now().Add(24 * time.Hour).UTC(),
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "applicaton/json")

	if err := json.NewEncoder(w).Encode(newCourse); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Fprintf(w, "Creating a course with UUID: %v\n", courseID)
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
