package app

import (
	"html/template"
	"net/http"
	"time"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Record struct {
	ID        ID        `json:"id"`
	PaperID   ID        `json:"paper_id"`
	UserID    ID        `json:"user_id"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ServeLoginForm serves a login form to the user
func ServeLoginForm(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Login</title>
		</head>
		<body>
			<h1>Login</h1>
			<form method="POST" action="/login">
				<label for="username">Username:</label>
				<input type="text" id="username" name="username" required><br>
				<label for="password">Password:</label>
				<input type="password" id="password" name="password" required><br>
				<button type="submit">Login</button>
			</form>
		</body>
		</html>
		`
		t, err := template.New("login").Parse(tmpl)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		t.Execute(w, nil)
	} else if r.Method == http.MethodPost {
		// Handle login logic here
		username := r.FormValue("username")
		password := r.FormValue("password")
		// For now, just echo the username and password
		w.Write([]byte("Username: " + username + ", Password: " + password))
	} else {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
