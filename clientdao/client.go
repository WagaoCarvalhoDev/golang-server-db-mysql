package client

import (
	"db"
	"encoding/json"
	"err"
	"fmt"
	"models"
	"net/http"
	"strconv"
	"strings"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	sid := strings.TrimPrefix(r.URL.Path, "/users/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userId(w, r, id)
	case r.Method == "GET":
		usersAll(w, r)
	case r.Method == "POST":
		createUser(w, r, "Wagner")
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Sorry... :(")
	}
}

func userId(w http.ResponseWriter, r *http.Request, id int) {
	db, erro := db.OpenConn()
	err.Err(erro)
	defer db.Close()

	var u models.User
	db.QueryRow("SELECT id, name FROM users WHERE id = ?", id).Scan(&u.Id, &u.Name)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application-json")
	fmt.Fprint(w, string(json))

}

func usersAll(w http.ResponseWriter, r *http.Request) {
	db, erro := db.OpenConn()
	err.Err(erro)
	defer db.Close()

	rows, _ := db.Query("SELECT id, name FROM users")
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-Type", "application-json")
	fmt.Fprint(w, string(json))
}

func createUser(w http.ResponseWriter, r *http.Request, u string) {
	db, erro := db.OpenConn()
	err.Err(erro)
	defer db.Close()

	stmt, erro := db.Prepare("INSERT INTO users(name) VALUES(?)")
	err.Err(erro)
	defer stmt.Close()

	res, erro := stmt.Exec(u)
	err.Err(erro)

	lastID, erro := res.LastInsertId()
	err.Err(erro)

	rows, erro := res.RowsAffected()
	err.Err(erro)

	fmt.Printf("Inseriu %d linha(s) com o ID %d.\n", rows, lastID)
}
