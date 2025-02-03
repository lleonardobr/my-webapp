package handlers

import (
	"fmt"
	"my-webapp/models"
	"my-webapp/utils"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/login.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		user, err := models.AuthenticateUser(email, password)
		if err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Create session or token
		token, err := utils.CreateToken(user.ID)
		if err != nil {
			http.Error(w, "Failed to create token", http.StatusInternalServerError)
			return
		}

		// Set token in cookie or header
		http.SetCookie(w, &http.Cookie{
			Name:  "token",
			Value: token,
			Path:  "/",
		})

		http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "templates/login.html")
	}
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		err := models.CreateUser(email, password)
		if err != nil {
			http.Error(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "templates/register.html")
	}
}

func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		email := r.FormValue("email")

		// Implement reset logic
		fmt.Println("Reset password for:", email)
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	} else {
		http.ServeFile(w, r, "templates/reset.html")
	}
}

func GoogleAuthHandler(w http.ResponseWriter, r *http.Request) {
	// Implement Google OAuth2 logic
	fmt.Println("Google Auth")
	http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
}
