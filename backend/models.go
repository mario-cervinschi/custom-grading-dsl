package main

type ReadFileRequest struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type CreateFileRequest struct {
	FileName string `json:"file_name"`
	Content  string `json:"content,omitempty"`
}

type EvaluateAllRequest struct {
	OperationsFile string `json:"operations_file"`
	DataFile       string `json:"data_file"`
	Sequential     bool   `json:"sequential,omitempty"`
}

type DataResult struct {
	ID   string          `json:"id"`
	Data map[string]Data `json:"data"`
}

type EvaluateAllResponse struct {
	Success bool         `json:"success"`
	Results []DataResult `json:"results"`
	Error   *Error       `json:"error,omitempty"`
}

type SaveFileRequest struct {
	Type    string `json:"type"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

type ReadFileSpecificRequest struct {
	Name string `json:"name"`
}

type ColumnDef struct {
	Field  string `json:"field"`
	Header string `json:"header"`
}

type TableDataResponse struct {
	Success bool                     `json:"success"`
	Columns []ColumnDef              `json:"columns"`
	Data    []map[string]interface{} `json:"data"`
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

type BenchmarkRequest struct {
	OperationsFile string `json:"operations_file"`
	DataFile       string `json:"data_file"`
	Iterations     int    `json:"iterations"`
	Sequential     bool   `json:"sequential,omitempty"`
}

type BenchmarkResponse struct {
	Success       bool      `json:"success"`
	Iterations    int       `json:"iterations"`
	TimesMs       []float64 `json:"times_ms"`
	AverageTimeMs float64   `json:"average_time_ms"`
	Error         *Error    `json:"error,omitempty"`
}
