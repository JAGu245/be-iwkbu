package handlers

import (
	"database/sql"
	"encoding/json"
	"jm-CICO/config"
	"jm-CICO/models"
	"jm-CICO/utils"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	var hashedPassword string
	err := config.DB.QueryRow("SELECT id, username, fullname, password, role FROM users WHERE username = ?", req.Username).
		Scan(&user.ID, &user.Username, &user.Fullname, &hashedPassword, &user.Role)

	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	if !utils.CheckPasswordHash(req.Password, hashedPassword) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateJWT(user.Username, user.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.LoginResponse{
		Token: token,
		User:  user,
	})
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req models.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Get username from token claims (via context)
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized: invalid claims", http.StatusUnauthorized)
		return
	}

	usernameFromToken, _ := claims["username"].(string)
	if usernameFromToken == "" {
		http.Error(w, "Unauthorized: missing username", http.StatusUnauthorized)
		return
	}

	var storedPassword string
	var userID int
	err := config.DB.QueryRow("SELECT id, password FROM users WHERE username = ?", usernameFromToken).
		Scan(&userID, &storedPassword)

	if err == sql.ErrNoRows {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	// Verify current password
	if !utils.CheckPasswordHash(req.CurrentPassword, storedPassword) {
		http.Error(w, "Kata sandi saat ini salah", http.StatusUnauthorized)
		return
	}

	// Perform update
	// Perform update
	query := "UPDATE users SET "
	var params []interface{}

	if req.NewUsername != "" {
		query += "username = ?, "
		params = append(params, req.NewUsername)
	}
	if req.NewFullname != "" {
		query += "fullname = ?, "
		params = append(params, req.NewFullname)
	}
	if req.NewPassword != "" {
		newHashedPassword, _ := utils.HashPassword(req.NewPassword)
		query += "password = ?, "
		params = append(params, newHashedPassword)
	}

	// Remove trailing comma and space
	if len(params) == 0 {
		http.Error(w, "No changes provided", http.StatusBadRequest)
		return
	}
	query = query[:len(query)-2]
	query += " WHERE id = ?"
	params = append(params, userID)

	_, err = config.DB.Exec(query, params...)
	if err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Profil berhasil diperbarui"})
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Verify that the requester is an admin
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized: invalid claims", http.StatusUnauthorized)
		return
	}

	requesterRole, _ := claims["role"].(string)
	if requesterRole != "admin" {
		http.Error(w, "Forbidden: priority access required", http.StatusForbidden)
		return
	}

	var req struct {
		Username string `json:"username"`
		Fullname string `json:"fullname"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	if req.Username == "" || req.Password == "" {
		http.Error(w, "Username and Password are required", http.StatusBadRequest)
		return
	}

	if req.Fullname == "" {
		req.Fullname = req.Username
	}

	if req.Role == "" {
		req.Role = "user"
	}

	hashedPassword, _ := utils.HashPassword(req.Password)
	_, err := config.DB.Exec("INSERT INTO users (username, fullname, password, role) VALUES (?, ?, ?, ?)", req.Username, req.Fullname, hashedPassword, req.Role)
	if err != nil {
		http.Error(w, "Failed to create user (possible duplicate username)", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User berhasil ditambahkan"})
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Verify that the requester is an admin
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized: invalid claims", http.StatusUnauthorized)
		return
	}

	requesterRole, _ := claims["role"].(string)
	if requesterRole != "admin" {
		http.Error(w, "Forbidden: priority access required", http.StatusForbidden)
		return
	}

	rows, err := config.DB.Query("SELECT id, username, fullname, role, created_at FROM users ORDER BY created_at DESC")
	if err != nil {
		http.Error(w, "Server error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Fullname, &user.Role, &user.CreatedAt); err != nil {
			log.Println("Scan error:", err)
			continue
		}
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(map[string]interface{}{
		"data": users,
	})
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	utils.EnableCORS(w)
	if r.Method == "OPTIONS" {
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Verify that the requester is an admin
	claims, ok := r.Context().Value(utils.UserContextKey).(jwt.MapClaims)
	if !ok {
		http.Error(w, "Unauthorized: invalid claims", http.StatusUnauthorized)
		return
	}

	requesterRole, _ := claims["role"].(string)
	requesterUsername, _ := claims["username"].(string)

	if requesterRole != "admin" {
		http.Error(w, "Forbidden: priority access required", http.StatusForbidden)
		return
	}

	var req struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Prevent self-deletion
	if req.Username == requesterUsername {
		http.Error(w, "Cannot delete your own account", http.StatusForbidden)
		return
	}

	_, err := config.DB.Exec("DELETE FROM users WHERE id = ?", req.ID)
	if err != nil {
		http.Error(w, "Failed to delete user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "User berhasil dihapus"})
}

// SeedUser creates an admin user if none exists
func SeedUser() {
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Println("Seed error:", err)
		return
	}

	if count == 0 {
		hashedPassword, _ := utils.HashPassword("admin123")
		_, err := config.DB.Exec("INSERT INTO users (username, fullname, password, role) VALUES (?, ?, ?, ?)", "admin", "Administrator Jasa Raharja", hashedPassword, "admin")
		if err != nil {
			log.Println("Failed to seed admin:", err)
			return
		}
		log.Println("Admin user seeded: admin / admin123")
	}
}
