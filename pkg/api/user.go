package api

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	Id             string
	Username       string
	Role           int
	HashedPassword string
	Verified       bool
}

func getUserByKeyValue(key string, value string) (User, error) { // stops code duplication for GetUserById and GetUserByUsername
	user := User{}
	row := db.QueryRow("SELECT id, username, role, password, verified FROM users WHERE "+key+" = $1", value)
	switch err := row.Scan(&user.Id, &user.Username, &user.Role, &user.HashedPassword, &user.Verified); err {
	case sql.ErrNoRows:
		return user, errors.New("USER_NOT_FOUND")
	case nil:
		return user, nil
	default:
		log.Print(err)
		return user, errors.New("INTERNAL_SERVER_ERROR")
	}
}

func GetUserById(id string) (User, error) {
	user, err := getUserByKeyValue("id", id)
	return user, err
}
