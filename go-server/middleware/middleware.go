package middleware

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"test-sharing-vision/go-server/models"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

// Constraints variable
const minUserName, minPassword, minName int = 3, 7, 3

func loadTheEnv() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func dbConn() (*sql.DB, error) {
	loadTheEnv()

	// Database Username
	dbUser := os.Getenv("DB_USER")

	// Database Password
	dbPass := os.Getenv("DB_PASSWORD")

	// Database Name
	dbName := os.Getenv("DB_NAME")

	dbDriver := "mysql"

	// connect to MySQL
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	fmt.Println("Connected to MySQL!")
	return db, nil
}

// CreateUser create user route
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "POST" {
		// header check
		if r.Header.Get("Content-Type") != "application/json" {
			http.Error(w, "Use content type application / json", http.StatusBadRequest)
			return
		}

		var user models.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Data Validation Process
		isValid, message := DataValidation(user)
		if !isValid {
			http.Error(w, message, http.StatusBadRequest)
			return
		}

		db, err := dbConn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer db.Close()

		// hashing password
		h := sha256.New()
		h.Write([]byte(user.Password))
		passwordHash := hex.EncodeToString(h.Sum(nil))

		query, err := db.Prepare("INSERT INTO users (username, password, name) VALUES(?,?,?)")
		if err != nil {
			panic(err.Error())
		}
		query.Exec(user.UserName, passwordHash, user.Name)
		//log.Println("INSERT: UserName: " + user.UserName + " | Password: " + user.Password + " | Name: " + user.Name)
	}
}

// GetAllUserPagination get all user with pagination route
func GetAllUserPagination(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "GET" {
		var users []models.User
		params := mux.Vars(r)
		limit := params["limit"]
		offset := params["offset"]

		db, err := dbConn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer db.Close()

		query, err := db.Query("SELECT id, username, password, name FROM users LIMIT ? OFFSET ?", limit, offset)
		if err != nil {
			panic(err.Error())
		}

		for query.Next() {
			var each = models.User{}
			var err = query.Scan(&each.ID, &each.UserName, &each.Password, &each.Name)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			users = append(users, each)
		}

		jsonData, err := json.Marshal(users)
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}

// GetUserByID get user by ID route
func GetUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == "GET" {
		var user models.UserResponse

		params := mux.Vars(r)
		userID := params["id"]

		db, err := dbConn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer db.Close()

		err = db.QueryRow("SELECT username, password, name FROM users WHERE id = ?", userID).Scan(&user.UserName, &user.Password, &user.Name)
		if err != nil {
			panic(err.Error())
		}

		jsonData, err := json.Marshal(user)
		if err != nil {
			return
		}
		fmt.Println(string(jsonData))

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}

// UpdateUser update one user route
func UpdateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "PUT" {
		var user models.User

		params := mux.Vars(r)
		user.ID, _ = strconv.Atoi(params["id"])

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Data Validation Process
		isValid, message := DataValidation(user)
		if !isValid {
			http.Error(w, message, http.StatusBadRequest)
			return
		}

		db, err := dbConn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer db.Close()

		// hashing password
		h := sha256.New()
		h.Write([]byte(user.Password))
		passwordHash := hex.EncodeToString(h.Sum(nil))

		query, err := db.Prepare("UPDATE users SET username = ?, password = ?, name = ? WHERE id = ?")
		if err != nil {
			panic(err.Error())
		}
		query.Exec(user.UserName, passwordHash, user.Name, user.ID)
		//log.Println("UPDATE: UserID: " + userID + " | UserName: " + user.UserName + " | Password: " + user.Password + " | Name: " + user.Name)
	}
}

// DeleteUser delete one user route
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "DELETE" {
		params := mux.Vars(r)
		userID := params["id"]

		db, err := dbConn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer db.Close()

		query, err := db.Prepare("DELETE FROM users WHERE id = ?")
		if err != nil {
			panic(err.Error())
		}
		query.Exec(userID)
	}
}

// GetAllUser get all user route
func GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "GET" {
		var users []models.User

		db, err := dbConn()
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		defer db.Close()

		query, err := db.Query("SELECT id, username, password, name FROM users")
		if err != nil {
			panic(err.Error())
		}

		for query.Next() {
			var each = models.User{}
			var err = query.Scan(&each.ID, &each.UserName, &each.Password, &each.Name)

			if err != nil {
				fmt.Println(err.Error())
				return
			}

			users = append(users, each)
		}

		jsonData, err := json.Marshal(users)
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}
}

// DataValidation . . .
func DataValidation(user models.User) (bool, string) {
	message := ""
	// username length check
	if len(user.UserName) < minUserName {
		message += "Username must consist " + strconv.Itoa(minUserName) + " characters or more\n"
		return false, message
	}

	// password length check
	if len(user.Password) < minPassword {
		message += "Password must consist " + strconv.Itoa(minPassword) + " characters or more\n"
		return false, message
	}

	// name length check
	if len(user.Name) < minName {
		message += "Name must consist " + strconv.Itoa(minName) + " characters or more\n"
		return false, message
	}
	return true, message
}
