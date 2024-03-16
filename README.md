# Why
The major objective of this tool was to parse a JSON text without loading all of it into memory. However, a syntax tree is still constructed for any later inferences.

# Build
Clone the repo
```bash
git clone github.com/jnafolayan/json-parser
```

Build the jp cli
```bash
go build -o jp ./cmd/jp
```

# Usage
If using the compiled cli
```bash
# Parse a JSON file
jp file.json
# Supply a raw JSON text
jp -raw '["hello world"]'
```

Otherwise, run the program
```bash
# Parse a JSON file
go run cmd/jp/main.go file.json
# Supply a raw JSON text
go run cmd/jp/main.go -raw '["hello world"]'
```