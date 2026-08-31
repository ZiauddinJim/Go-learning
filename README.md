# 🐹 Go Fundamentals & Learning Journey

Welcome to the **Go-learning** repository! This project serves as a practical, code-based guide covering core Golang fundamentals and syntax basics.

---

## 📑 Topics Covered in `main.go`

This repository demonstrates key Go language concepts through structured functions inside [`main.go`](file:///G:/Golang/main.go):

| Topic | Function | Description |
| :--- | :--- | :--- |
| **Go Syntax & Variables** | `syntax()`, `variable()` | Standard `var` keyword vs. short declaration operator (`:=`). |
| **Initial / Zero Values** | `initialValue()` | Declaring variables without explicit initial values (`string`, `int`, `bool`). |
| **Multiple Values** | `MultipleValue()` | Declaring multiple variables of single or mixed types in a single line. |
| **Block Declarations** | `BlockDeclaration()` | Grouping variable declarations using `var (...)`. |
| **Constants** | `Constants()`, `multiConst()` | Single and block constant declarations using the `const` keyword. |
| **Output Functions** | `outputFunc()` | Formatting and printing output using `fmt.Print()`, `fmt.Println()`, and `fmt.Printf()`. |

---

## 🚀 How to Run the Code

Make sure you have [Go](https://go.dev/dl/) installed on your machine.

1. **Clone the repository:**
   ```bash
   git clone https://github.com/ZiauddinJim/Go-learning.git
   cd Go-learning
   ```

2. **Run the application:**
   ```bash
   go run main.go
   ```

---

## 💡 Quick Code Overview

Here is a quick look at the core concepts in `main.go`:

### 1. Variable Declarations
```go
var name string = "jim" // Explicit type declaration
age := 25               // Short declaration (type inferred)

var (
    a int
    b int    = 1
    c string = "mim"
)
```

### 2. Constants
```go
const A string = "jim"
const (
    a        = 1
    b string = "cat"
)
```

### 3. Print & Formatting Verbs
```go
fmt.Print("No automatic newline or spacing")
fmt.Println("Automatic spacing and newline")
fmt.Printf("Name: %v, Type: %T\n", name, name)
```

---

## 📝 License
This project is open source and available under the terms in [LICENSE](file:///G:/Golang/LICENSE).
