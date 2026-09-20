package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {

	// database
	database, errDB := Connect()
	if errDB != nil {
		log.Fatalf("Connection Failed: %v", errDB)
	}
	defer database.Close()

	handler := &DBHandler{
		db: database,
	}

	// routing
	router := httprouter.New()
	router.GET("/employees", handler.getEmployees) // done
	router.POST("/employees", handler.addEmployee) // done

	router.GET("/employees/:id", handler.getEmployee)    // done
	router.PUT("/employees/:id", handler.updateEmployee) // done
	router.DELETE("/employees/:id", handler.getEmployees)

	// server
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
