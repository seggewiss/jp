# jp — JSON Path Parser (Go)

A lightweight command-line tool written in Go that parses JSON files and extracts values using **dot-notation paths**.

This project is built as a learning exercise to deepen my understanding of Go, type handling, error design, and CLI tooling.

---

## 🚧 Status

> ⚠️ This project is **still under active development**  
> Expect breaking changes, missing edge cases, and ongoing refactoring.

Current focus areas:
- API stabilization
- Performance improvements
- Expanded JSON path features
- Better CLI UX

---

## ✨ Features

- Parse JSON files from disk
- Access nested values using dot notation
- Supports:
    - Objects
    - Arrays (by index)
    - Scalars (string, number, boolean, null)
- Custom error handling for:
    - Invalid keys
    - Array out-of-bounds access
- Pretty printing result with `--pretty`

---

## 📦 Installation

Get the pkg

```bash
go get https://github.com/seggewiss/jp.git
```

## 🚀 Usage
```shell
jp [--pretty] <file> <dotNotationPath>
```

## Example
Given `data.json`:
```json
{
  "user": {
    "name": "John",
    "age": 30,
    "tags": ["dev", "go"]
  }
}
```

Run:
```shell
jp data.json user.name
```

Output:
```json
{
  "user.name": "John"
}
```

Array access:
```shell
jp data.json user.tags.0
```

Output:
```json
{
  "user.tags.0": "dev"
}
```

## 🧠 Why this project exists

This project is part of my effort to:

* Get my hands into Go again
* Build production-style project structure (cmd/, pkg/)
* Practice error handling patterns
* Work with encoding/json internals
* Improve testing discipline

## ⚠️ Limitations (current)
* No JSONPath support (only simple dot notation)
* No wildcards or filters
* No streaming for large JSON files
* Limited CLI features
* Error messages still evolving

## 🔮 Future improvements
* JSONPath-like expressions
* Performance optimizations for large files
* Optional JSON streaming decoder
* More expressive error reporting
* Benchmark suite

## 📌 Notes
This project is intentionally kept minimal and experimental.
The goal is not just functionality, but understanding how Go handles real-world data traversal and tooling design.