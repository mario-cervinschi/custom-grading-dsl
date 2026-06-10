package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

func enableCors(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

func evaluateAllHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req EvaluateAllRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid JSON body"})
		return
	}

	if req.OperationsFile == "" || req.DataFile == "" ||
		strings.Contains(req.OperationsFile, "..") || strings.Contains(req.DataFile, "..") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid request parameters"})
		return
	}

	opFileName := req.OperationsFile
	if !strings.HasSuffix(opFileName, ".txt") {
		opFileName += ".txt"
	}
	opFilePath := filepath.Join("resources/operations", opFileName)

	dataFileName := req.DataFile
	if !strings.HasSuffix(dataFileName, ".txt") {
		dataFileName += ".txt"
	}
	dataFilePath := filepath.Join("resources/data", dataFileName)

	customOperations, err := readOperations(opFilePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Error: "Operations file not found"})
		return
	}

	customDB, err := readData(dataFilePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Error: "Data file not found"})
		return
	}

	startTime := time.Now()
	fmt.Printf("Calculations started..")

	var resultsArray []DataResult

	if req.Sequential {
		for ID, data := range customDB {
			memory, explanations, err := runCustomEvaluateLogic(data, customOperations)
			if err != nil {
				continue
			}

			var dataMap = make(map[string]Data)

			for _, expl := range explanations {
				if expl.Variable != "" {
					variable := expl.Variable
					data_cell := Data{Result: memory[variable]}

					step := &ExplanationStep{}

					if expl.OriginalExpression != "" && expl.SubstitutedExpression != "" {
						step = &ExplanationStep{
							Original:    expl.OriginalExpression,
							Substituted: expl.SubstitutedExpression,
							Result:      expl.Result,
						}
					}
					if expl.Description != "" {
						step.Description = expl.Description
					}
					data_cell.Explanation = step
					dataMap[variable] = data_cell
				}
			}

			for memKey, memValue := range memory {
				if _, exists := dataMap[memKey]; !exists {
					dataMap[memKey] = Data{Result: memValue}
				}
			}

			resultsArray = append(resultsArray, DataResult{
				ID:   ID,
				Data: dataMap,
			})
		}
	} else {
		var mu sync.Mutex
		var wg sync.WaitGroup

		for ID, data := range customDB {
			wg.Add(1)
			go func(studentID string, studentData map[string]interface{}) {
				defer wg.Done()

				memory, explanations, err := runCustomEvaluateLogic(studentData, customOperations)
				if err != nil {
					return
				}

				var dataMap = make(map[string]Data)

				for _, expl := range explanations {
					if expl.Variable != "" {
						variable := expl.Variable
						data_cell := Data{Result: memory[variable]}

						step := &ExplanationStep{}

						if expl.OriginalExpression != "" && expl.SubstitutedExpression != "" {
							step = &ExplanationStep{
								Original:    expl.OriginalExpression,
								Substituted: expl.SubstitutedExpression,
								Result:      expl.Result,
							}
						}
						if expl.Description != "" {
							step.Description = expl.Description
						}
						data_cell.Explanation = step
						dataMap[variable] = data_cell
					}
				}

				for memKey, memValue := range memory {
					if _, exists := dataMap[memKey]; !exists {
						dataMap[memKey] = Data{Result: memValue}
					}
				}

				mu.Lock()
				resultsArray = append(resultsArray, DataResult{
					ID:   studentID,
					Data: dataMap,
				})
				mu.Unlock()
			}(ID, data)
		}
		wg.Wait()
	}

	finishTime := time.Now()
	elapsedTime := finishTime.Sub(startTime)

	fmt.Printf("Calculations started at: %s | Finished: %s | Elapsed: %s\n",
		startTime.Format("15:04:05.000"),
		finishTime.Format("15:04:05.000"),
		elapsedTime)

	response := EvaluateAllResponse{
		Success: true,
		Results: resultsArray,
	}

	json.NewEncoder(w).Encode(response)
}

func parseINItoTable(content string) ([]ColumnDef, []map[string]interface{}) {
	lines := strings.Split(content, "\n")

	var data []map[string]interface{}
	var currentRow map[string]interface{}

	columnsMap := make(map[string]bool)
	var orderedColumns []string

	columnsMap["id"] = true
	orderedColumns = append(orderedColumns, "id")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			if currentRow != nil {
				data = append(data, currentRow)
			}

			currentRow = make(map[string]interface{})
			idVal := strings.Trim(line, "[]")
			currentRow["id"] = idVal
			continue
		}

		if currentRow != nil && strings.Contains(line, "=") {
			parts := strings.SplitN(line, "=", 2)
			if len(parts) == 2 {
				key := strings.TrimSpace(parts[0])
				val := strings.TrimSpace(parts[1])

				currentRow[key] = val

				if !columnsMap[key] {
					columnsMap[key] = true
					orderedColumns = append(orderedColumns, key)
				}
			}
		}
	}

	if currentRow != nil {
		data = append(data, currentRow)
	}

	var columns []ColumnDef
	for _, colName := range orderedColumns {
		header := strings.ToUpper(colName[:1]) + colName[1:]
		if colName == "id" {
			header = "ID"
		}
		columns = append(columns, ColumnDef{
			Field:  colName,
			Header: header,
		})
	}

	return columns, data
}

func readDataFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req ReadFileSpecificRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid JSON body"})
		return
	}

	fileName := req.Name
	if !strings.HasSuffix(fileName, ".txt") {
		fileName += ".txt"
	}
	filePath := filepath.Join("resources/data", fileName)

	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Error{Error: "File not found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Error: "Failed to read file"})
		return
	}

	cols, tableData := parseINItoTable(string(contentBytes))

	json.NewEncoder(w).Encode(TableDataResponse{
		Success: true,
		Columns: cols,
		Data:    tableData,
	})
}

func readOperationsFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req ReadFileSpecificRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid JSON body"})
		return
	}

	fileName := req.Name
	if !strings.HasSuffix(fileName, ".txt") {
		fileName += ".txt"
	}
	filePath := filepath.Join("resources/operations", fileName)

	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Error{Error: "File not found"})
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Error: "Failed to read file"})
		return
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"content": string(contentBytes),
	})
}

func saveFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req SaveFileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid JSON body"})
		return
	}

	var targetDirectory string
	switch req.Type {
	case "data":
		targetDirectory = "resources/data"
	case "operations":
		targetDirectory = "resources/operations"
	default:
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid type. Must be 'data' or 'operations'"})
		return
	}

	if req.Name == "" || strings.Contains(req.Name, "..") || strings.Contains(req.Name, "/") || strings.Contains(req.Name, "\\") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid file name"})
		return
	}

	fileName := req.Name
	if !strings.HasSuffix(fileName, ".txt") {
		fileName += ".txt"
	}

	filePath := filepath.Join(targetDirectory, fileName)

	if err := os.WriteFile(filePath, []byte(req.Content), 0644); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Error: "Failed to save file"})
		return
	}

	fmt.Printf("[LOG] - Saved changes to file: %s\n", filePath)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "File saved successfully",
	})
}

func createDataFileHandler(w http.ResponseWriter, r *http.Request) {
	processCreateFile(w, r, "resources/data")
}

func createOperationsFileHandler(w http.ResponseWriter, r *http.Request) {
	processCreateFile(w, r, "resources/operations")
}

func listDataFilesHandler(w http.ResponseWriter, r *http.Request) {
	processListFiles(w, r, "resources/data")
}

func listOperationsFilesHandler(w http.ResponseWriter, r *http.Request) {
	processListFiles(w, r, "resources/operations")
}

func evaluateBenchmarkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req BenchmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid JSON body"})
		return
	}

	if req.OperationsFile == "" || req.DataFile == "" ||
		strings.Contains(req.OperationsFile, "..") || strings.Contains(req.DataFile, "..") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid request parameters"})
		return
	}

	opFileName := req.OperationsFile
	if !strings.HasSuffix(opFileName, ".txt") {
		opFileName += ".txt"
	}
	opFilePath := filepath.Join("resources/operations", opFileName)

	dataFileName := req.DataFile
	if !strings.HasSuffix(dataFileName, ".txt") {
		dataFileName += ".txt"
	}
	dataFilePath := filepath.Join("resources/data", dataFileName)

	customOperations, err := readOperations(opFilePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Error: "Operations file not found"})
		return
	}

	customDB, err := readData(dataFilePath)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Error: "Data file not found"})
		return
	}

	iterations := req.Iterations
	if iterations <= 0 {
		iterations = 1
	}

	var timesMs []float64
	var sumMs float64

	fmt.Printf("\n[BENCHMARK] Executing %d iterations for %s with %s...\n", iterations, opFileName, dataFileName)

	for i := 0; i < iterations; i++ {
		start := time.Now()

		if req.Sequential {
			for _, data := range customDB {
				runCustomEvaluateLogic(data, customOperations)
			}
		} else {
			var wg sync.WaitGroup
			for _, data := range customDB {
				wg.Add(1)
				go func(studentVars map[string]interface{}) {
					defer wg.Done()
					runCustomEvaluateLogic(studentVars, customOperations)
				}(data)
			}
			wg.Wait()
		}

		elapsed := time.Since(start)
		ms := float64(elapsed.Microseconds()) / 1000.0
		timesMs = append(timesMs, ms)
		sumMs += ms

		fmt.Printf("Run %d/%d: %.3f ms\n", i+1, iterations, ms)
	}

	avgMs := sumMs / float64(iterations)
	fmt.Printf("[BENCHMARK FINISHED] Average: %.3f ms\n", avgMs)

	response := BenchmarkResponse{
		Success:       true,
		Iterations:    iterations,
		TimesMs:       timesMs,
		AverageTimeMs: avgMs,
	}

	json.NewEncoder(w).Encode(response)
}
