package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func processCreateFile(w http.ResponseWriter, r *http.Request, targetDirectory string) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var req CreateFileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid JSON body"})
		return
	}

	if req.FileName == "" || strings.Contains(req.FileName, "..") || strings.Contains(req.FileName, "/") || strings.Contains(req.FileName, "\\") {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Error{Error: "Invalid file name"})
		return
	}

	fileName := req.FileName
	if !strings.HasSuffix(fileName, ".txt") {
		fileName += ".txt"
	}

	if err := os.MkdirAll(targetDirectory, os.ModePerm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Error: "Failed to create directory"})
		return
	}

	filePath := filepath.Join(targetDirectory, fileName)

	if _, err := os.Stat(filePath); err == nil {
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(Error{Error: "A file with this name already exists"})
		return
	}

	if err := os.WriteFile(filePath, []byte(req.Content), 0644); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Error{Error: "Failed to create file on disk"})
		return
	}

	fmt.Printf("[LOG] - Created file: %s\n", filePath)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success":   true,
		"file_path": filePath,
	})
}

func processListFiles(w http.ResponseWriter, r *http.Request, targetDirectory string) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	files := []string{}
	entries, err := os.ReadDir(targetDirectory)
	if err == nil {
		for _, entry := range entries {
			if !entry.IsDir() {
				files = append(files, entry.Name())
			}
		}
	}

	json.NewEncoder(w).Encode(files)
}
