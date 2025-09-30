
# React + Vite & Go Server

This project includes a React + Vite frontend and a Go backend to save user info to a MySQL database.

## Go Server Setup
1. Update the DSN in `server.go` with your MySQL username and password.
2. Make sure MySQL is running and the `userInfo` database exists.
3. Install dependencies:
	- `go get github.com/go-sql-driver/mysql`
4. Run the server:
	- `go run server.go`

## API
- POST `/submit` with form fields: `name`, `age`, `email`, `contact`

## Example curl
```
curl -X POST -d "name=John&age=30&email=john@example.com&contact=1234567890" http://localhost:8080/submit
```

## Push to GitHub
1. Initialize git if not done:
	- `git init`
2. Add remote:
	- `git remote add origin https://github.com/<your-username>/userInfo.git`
3. Add, commit, and push:
	- `git add server.go README.md`
	- `git commit -m "Add Go server for userInfo"`
	- `git push -u origin main`
