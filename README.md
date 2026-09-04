# Custom DSL Evaluation Engine
A modern web application and evaluation engine for a custom, domain-specific language (DSL) designed for grading and academic evaluations. The system features a responsive Angular web interface, a multi-threaded evaluation engine in Go, and real-time IDE features (diagnostics, syntax highlighting, autocomplete) powered by a custom Language Server Protocol (LSP) over WebSockets.

---

## 📂 Project Structure

```text
app/
├── backend/                       # Go evaluation backend
│   ├── evaluator/                 # Core visitor evaluation logic & builders
│   ├── parser/                    # ANTLR-generated DSL lexer & parser
│   ├── resources/                 # Storage for files (data & operations)
│   │   ├── data/                  # INI files containing datasets
│   │   └── operations/            # Text files containing DSL formulas
│   ├── grading.go                 # IO functions & evaluation orchestrator
│   ├── handlers.go                # HTTP handlers & concurrency endpoints
│   ├── main.go                    # Go server entrypoint
│   └── models.go                  # Go request/response structs
│
├── frontend/                      # Angular web application
│   ├── src/
│   │   ├── app/
│   │   │   ├── core/              # Shared client logic (services, LSP client)
│   │   │   ├── pages/             # App pages
│   │   │   └── shared/            # Shared UI components (header, modals)
│   │   └── environments/          # App configurations
│   └── package.json
│
└── suggestions/                   # Custom Language Server Protocol (LSP)
    ├── parser/                    # Tree-sitter grammar rules for DSL
    └── extension/
        ├── server/                # LSP Server containing auto-completer and diagnostic rules
        └── ws-server/             # WebSocket server wrapper for JSON-RPC communication
```

---

## 📝 Custom DSL Syntax Reference

The evaluation engine interprets statements line-by-line. Standard features include:

### 1. Variables & Assignments
Variables are dynamically typed (supporting floats, booleans, and strings).
```c
laboratory_grade = 8.5;
bonus = 1.0;
```

### 2. Explainable Statements (`EXPLAIN`)
Using `EXPLAIN` instructs the engine to capture evaluation metadata (formula details, substituted values, results) for UI rendering:
```c
EXPLAIN final_grade = lab_grade * 0.3 + exam * 0.7, "Weighted Grade Calculation";
```

### 3. Conditional Guards (`WHEN ... EXISTS`)
Executes an operation only if a list of variables is fully defined in the current dataset record:
```c
WHEN bonus EXISTS EXPLAIN final_grade = MIN(10, final_grade + bonus), "Bonus added";
```

### 4. Vector Operations (`APPLY`)
Applies a lambda transformation to multiple variables at once:
```c
APPLY x => x + 1 TO laboratory_1, laboratory_2, laboratory_3;
```

### 5. Advanced Functions & Operators
*   **Ternary Operator**: `condition ? trueExpression : falseExpression`
*   **Built-in Functions**: `MAX(a, b)`, `MIN(a, b)`, `ROUND(x)`
*   **Regex Operations**: `variable ~ "pattern"` (match) and `variable !~ "pattern"` (no match)
*   **Predefined Status Constants**: `absent`, `present`, `excused`, `nothing`, `fraud`, `cancelled`, `invalid`, `alert`, `conflict`, `ungraded`, `obscured`, `toolow`.

---

## 🛠️ Installation & Setup

Ensure you have **Go (1.20+)**, **Node.js (18+)**, and **npm** installed on your system.

### 1. Backend Evaluation Engine (Go)
1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```
2. Build the Go application:
   ```bash
   go build .
   ```
3. Run the executable generated in the directory.

The backend server will launch and listen for HTTP connections on `http://localhost:8080`.

---

### 2. Frontend (Angular)
1. Navigate to the `frontend` directory:
   ```bash
   cd frontend
   ```
2. Install npm dependencies:
   ```bash
   npm install
   ```
3. Start the Angular development server:
   ```bash
   ng serve
   ```

Open your browser and navigate to `http://localhost:4200` to access the playground.

---

### 3. LSP Environment (TypeScript)
The suggestion and diagnostic subsystem requires running a WebSocket server wrapper that executes the LSP server.

1. Install dependencies for the **parser**, **server**, and **ws-server** modules:
   ```bash
   # Install Tree-Sitter parser dependencies
   cd suggestions/parser
   npm install

   # Install LSP server dependencies
   cd ../extension/server
   npm install

   # Install WebSocket server dependencies
   cd ../ws-server
   npm install
   ```

2. Compile the TypeScript source code for the **server** and **ws-server** components:
   ```bash
   # Compile LSP server
   cd ../server
   npm run compile

   # Compile WebSocket server
   cd ../ws-server
   npm run compile
   ```

3. Launch the final WebSocket service from the `ws-server` module:
   ```bash
   npm run start
   ```

The LSP suggestion engine will start running on `ws://localhost:3000` to serve real-time autocompletions and live syntax validations directly to the web editor.

---

## 📊 API & Endpoint Reference

### Go Backend Routes (`http://localhost:8080`)
*   `POST /evaluate/all`: Evaluates an operations file against a dataset file. Accepts `sequential: true/false` parameter to toggle multi-threaded parallel execution.
*   `POST /evaluate/benchmark`: Performs multiple evaluation iterations on a dataset to log average response times, comparing sequential vs. concurrent execution modes.
*   `POST /file/save`: Saves updated files back to disk (data or operations).
*   `POST /data/read` / `POST /operations/read`: Reads files from local resources to display in the UI.
*   `GET /data/list` / `GET /operations/list`: Lists saved files.
*   `POST /data/new` / `POST /operations/new`: Creates empty resource files.

### LSP WebSocket Connection (`ws://localhost:3000`)
*   Accepts standard LSP JSON-RPC payloads.
*   Supports standard methods: `initialize`, `textDocument/didOpen`, `textDocument/didChange`, `textDocument/completion`.
*   Supports custom request: `custom/validate` for AST syntax validation.
