[private]
default: dev

# Run the program
dev:
    go run main.go

# Get pending tasks from doc
p:
    cv docs/design.txt 0

