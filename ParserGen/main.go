package main

import (
	"ParserGen/evaluator"
	"ParserGen/parser"
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"github.com/joho/godotenv"
)

var cachedOperations []string
var studentDB map[string]evaluator.Variables

type SituatuionRequest struct {
	StudentID string `json:"student_id"`
	Group     string `json:"group"`
}

type ExplanationStep struct {
	Original    string      `json:"original,omitempty"`
	Substituted string      `json:"substituted,omitempty"`
	Result      interface{} `json:"result,omitempty"`
	Description string      `json:"description,omitempty"`
}

type Data struct {
	Result      interface{}      `json:"result"`
	Explanation *ExplanationStep `json:"explanation,omitempty"`
}

type Error struct {
	Error string `json:"error"`
}

type SituationResponse struct {
	Success bool            `json:"success"`
	Data    map[string]Data `json:"data,omitempty"`
	Error   *Error          `json:"error,omitempty"`
}

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

func runGradingLogic(studentData evaluator.Variables) (evaluator.Variables, []evaluator.ExplanationData, error) {
	eval := evaluator.NewEvaluator(studentData)
	for _, op := range cachedOperations {
		is := antlr.NewInputStream(op)
		lexer := parser.NewExprLexer(is)
		stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
		p := parser.NewExprParser(stream)
		p.RemoveErrorListeners()
		tree := p.Algorithm()
		eval.Visit(tree)
	}
	return eval.Memory, eval.AllExplanations, nil
}

func gradeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req SituatuionRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid JSON body"})
		return
	}

	fmt.Printf("[LOG] - Searching for student with ID %s\n", req.StudentID)

	studentData, found := fetchStudentData(req.StudentID)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(Error{Error: "Student ID not found in database"})
		return
	}

	if req.Group != "" {
		studentData["group"] = req.Group
	}

	memory, explanations, err := runGradingLogic(studentData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Error: "Calculation error"})
		return
	}

	response := SituationResponse{}
	var data = make(map[string]Data)

	for _, expl := range explanations {
		if expl.Variable != "" {
			variable := expl.Variable

			data_cell := Data{
				Result: memory[variable],
			}

			if expl.OriginalExpression != "" && expl.SubstitutedExpression != "" {
				step := ExplanationStep{}

				step.Original = expl.OriginalExpression
				step.Substituted = expl.SubstitutedExpression
				step.Result = expl.Result

				if expl.Description != "" {
					step.Description = expl.Description
				}

				data_cell.Explanation = &step
			}
			data[variable] = data_cell
		}
	}

	// ADD HERE EXTRA VARIABLE - VALUE PAIRS FROM memory

	data["esh"] = Data{Result: memory["esh"]}
	data["rsh"] = Data{Result: memory["rsh"]}
	data["epr"] = Data{Result: memory["epr"]}
	data["rpr"] = Data{Result: memory["rpr"]}
	data["eth"] = Data{Result: memory["esh"]}
	data["rth"] = Data{Result: memory["esh"]}

	//

	response.Success = true
	response.Data = data

	json.NewEncoder(w).Encode(response)
}

func fetchStudentData(studentID string) (evaluator.Variables, bool) {
	originalData, exists := studentDB[studentID]
	if !exists {
		return nil, false
	}

	clone := make(evaluator.Variables)
	for k, v := range originalData {
		clone[k] = v
	}
	return clone, true
}

// reads all valid operations from the file
// format:
// the file contains multiple operations
// an operation can be on multiple lines
// an operation is valid if it ends in ;

func readOperations(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var operations []string
	var currentOp strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		currentOp.WriteString(line)
		currentOp.WriteString("\n")

		if strings.Contains(line, ";") {
			operations = append(operations, currentOp.String())
			currentOp.Reset()
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return operations, nil
}

// reads all data related to a student
func readStudentData(filename string) (map[string]evaluator.Variables, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := make(map[string]evaluator.Variables)
	var currentID string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			currentID = strings.Trim(line, "[]")
			data[currentID] = make(evaluator.Variables)
			continue
		}

		if currentID != "" {
			parts := strings.Split(line, "=")
			if len(parts) != 2 {
				continue
			}
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])

			if strings.Contains(value, ".") {
				var floatVal float64
				if _, err := fmt.Sscanf(value, "%f", &floatVal); err == nil {
					data[currentID][key] = floatVal
					continue
				}
			} else {
				var intVal int
				if _, err := fmt.Sscanf(value, "%d", &intVal); err == nil {
					data[currentID][key] = float64(intVal)
					continue
				}
			}
			data[currentID][key] = value
		}
	}

	return data, scanner.Err()
}

func main() {
	fmt.Println("Main")

	err := godotenv.Load()
	if err != nil {
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
	cachedOperations, err = readOperations(operationsPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read operations: %v", err))
	}
	fmt.Printf("Success: Loaded %d operations.\n", len(cachedOperations))

	fmt.Printf("... Loading student DB from: %s\n", studentDataPath)
	studentDB, err = readStudentData(studentDataPath)
	if err != nil {
		panic(fmt.Sprintf("Failed to read student DB: %v", err))
	}
	fmt.Printf("Success: Loaded %d students.\n", len(studentDB))

	http.HandleFunc("/grade", enableCors(gradeHandler))

	port := ":8080"
	fmt.Printf("\nServer is running on http://localhost%s\n", port)
	fmt.Println("Waiting for POST requests on /grade endpoint...")

	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}
