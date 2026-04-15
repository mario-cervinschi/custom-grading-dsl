package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Main")

	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}

	studentDataPath := os.Getenv("STUDENT_DATA_FILE")
	operationsPath := os.Getenv("OPERATIONS_FILE")

	if studentDataPath == "" {
		studentDataPath = "student_data.txt"
	}
	if operationsPath == "" {
		operationsPath = "operations.txt"
	}

	fmt.Printf("... Loading operations from: %s\n", operationsPath)
	var err error
	cachedOperations, err = readOperations(operationsPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read operations: %v", err))
	}
	fmt.Printf("Success: Loaded %d operations.\n", len(cachedOperations))

	fmt.Printf("... Loading student DB from: %s\n", studentDataPath)
	studentDB, err = readData(studentDataPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read student DB: %v", err))
	}
	fmt.Printf("Success: Loaded %d students.\n", len(studentDB))

	http.HandleFunc("/grade", enableCors(gradeHandler))

	http.HandleFunc("/data/new", enableCors(createDataFileHandler))
	http.HandleFunc("/operations/new", enableCors(createOperationsFileHandler))

	http.HandleFunc("/data/list", enableCors(listDataFilesHandler))
	http.HandleFunc("/operations/list", enableCors(listOperationsFilesHandler))

	http.HandleFunc("/data/read", enableCors(readDataFileHandler))
	http.HandleFunc("/operations/read", enableCors(readOperationsFileHandler))

	http.HandleFunc("/file/save", enableCors(saveFileHandler))

	http.HandleFunc("/evaluate/all", enableCors(evaluateAllHandler))

	port := ":8080"
	fmt.Printf("\nServer is running on http://localhost%s\n", port)
	fmt.Println("Waiting for requests...")

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
