package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Printf("Main - Using %d CPU cores\n", runtime.NumCPU())

	http.HandleFunc("/data/new", enableCors(createDataFileHandler))
	http.HandleFunc("/operations/new", enableCors(createOperationsFileHandler))

	http.HandleFunc("/data/list", enableCors(listDataFilesHandler))
	http.HandleFunc("/operations/list", enableCors(listOperationsFilesHandler))

	http.HandleFunc("/data/read", enableCors(readDataFileHandler))
	http.HandleFunc("/operations/read", enableCors(readOperationsFileHandler))

	http.HandleFunc("/file/save", enableCors(saveFileHandler))

	http.HandleFunc("/evaluate/all", enableCors(evaluateAllHandler))
	http.HandleFunc("/evaluate/benchmark", enableCors(evaluateBenchmarkHandler))

	port := ":8080"
	fmt.Printf("\nServer is running on http://localhost%s\n", port)
	fmt.Println("Waiting for requests...")

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
