build:
	go build main.go
run:
	./main start
createDB:
	psql -c "CREATE DATABASE todo_app"
