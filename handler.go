package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type DBHandler struct {
	db *sql.DB
}

type Employee struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type newEmployee struct {
	Name  string
	Email string
	Phone string
}

// GET /employees
func (h *DBHandler) getEmployees(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	query := `SELECT id, name, email, phone, created_at, updated_at FROM employees`
	rows, err := h.db.Query(query)
	if err != nil {
		http.Error(res, "Failed to get employees", http.StatusBadRequest)
		return
	}
	defer rows.Close()

	employees := []Employee{}

	for rows.Next() {
		var employee Employee
		err := rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.Email,
			&employee.Phone,
			&employee.CreatedAt,
			&employee.UpdatedAt,
		)
		if err != nil {
			http.Error(res, "Failed to scan employee", http.StatusBadRequest)
			return
		}

		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		http.Error(res, "Failed to read employee", http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(employees)
}

// POST /employees
func (h *DBHandler) addEmployee(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var data newEmployee
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		http.Error(res, "invalid request body", http.StatusBadRequest)
		return
	}

	query := `
		INSERT INTO employees (name, email, phone) VALUES (?, ?, ?)
	`
	name := &data.Name
	email := &data.Email
	phone := &data.Phone

	_, err = h.db.Exec(query, name, email, phone)
	if err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			http.Error(res, "Email itu gak bisa dipake bang", http.StatusConflict)
			return
		}
		http.Error(res, "Failed to input new employee...", http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)

	json.NewEncoder(res).Encode(map[string]string{
		"message": "employee berhasil ditambahkan",
	})
}

// GET /employees/:id
func (h *DBHandler) getEmployee(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	query := `SELECT id, name, email, phone, created_at, updated_at FROM employees WHERE id=?`

	var employee Employee
	err := h.db.QueryRow(query, id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Email,
		&employee.Phone,
		&employee.CreatedAt,
		&employee.UpdatedAt,
	)

	if errors.Is(err, sql.ErrNoRows) {
		http.Error(res, "Employee not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(res, "Failed to scan employee", http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(employee)
}

// PUT /employees/:id
func (h *DBHandler) updateEmployee(res http.ResponseWriter, req *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	var data newEmployee
	err := json.NewDecoder(req.Body).Decode(&data)
	if err != nil {
		http.Error(res, "Failed to get request body....", http.StatusBadRequest)
		return
	}
	name := &data.Name
	email := &data.Email
	phone := &data.Phone

	query := `UPDATE employees SET name=?, email=?, phone=? WHERE id=?`
	_, err = h.db.Exec(query, name, email, phone, id)
	if err != nil {
		http.Error(res, "Failed to update the employee data....", http.StatusBadRequest)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)

	json.NewEncoder(res).Encode(map[string]string{
		"message": "update employee success",
	})
}

func (h *DBHandler) deleteEmployee(res http.ResponseWriter, _ *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	query := `DELETE FROM employees WHERE id=?`
	result, err := h.db.Exec(query, id)
	if err != nil {
		http.Error(res, "Failed to delete the employee", http.StatusBadRequest)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		http.Error(res, "Failed to check deleted employee", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(res, "Employee not found", http.StatusNotFound)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(map[string]string{
		"message": "success delete the employee",
	})
}
