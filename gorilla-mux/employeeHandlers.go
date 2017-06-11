package main

import (
	"net/http"
	"fmt"
	
	"github.com/satori/uuid"
	"github.com/gorilla/mux"
)

// EmployeesList is here
func EmployeesList(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "This is a list of Employees")
}

// CreateEmployee is here
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employeeID := uuid.NewV4()

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Creating an employee with UUID: %v\n", employeeID)
}

// EmployeeDetail is here
func EmployeeDetail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeID := vars["employeeId"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Viewing employee with ID: %v\n", employeeID)
}

// UpdateEmployee is here
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeID := vars["employeeId"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated employee with ID: %v\n", employeeID)
}

// DeleteEmployee is here
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	employeeID := vars["employeeId"]

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted employee with ID: %v\n", employeeID)
}