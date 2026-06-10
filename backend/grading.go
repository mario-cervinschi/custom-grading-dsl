package main

import (
	"ParserGen/evaluator"
	"ParserGen/parser"
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

func runCustomEvaluateLogic(data evaluator.Variables, operations []string) (evaluator.Variables, []evaluator.ExplanationData, error) {
	eval := evaluator.NewEvaluator(data)
	for _, op := range operations {
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
	return operations, scanner.Err()
}

func readData(filename string) (map[string]evaluator.Variables, error) {
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
