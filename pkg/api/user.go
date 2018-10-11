package api

import (
	"database/sql"
	"encoding/base64"
	"errors"
	"log"

	"crypto/rand"

	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id             string
	Username       string
	Token          string
	Role           int
	HashedPassword string
	Verified       bool
}

func (user User) checkPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	correct := err == nil
	return correct
}

func UsernameInUse(username string) (bool, error) {
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE lower(username) = lower($1)", username)
	var count int
	switch err := row.Scan(&count); err {
	case sql.ErrNoRows:
		return false, nil
	case nil:
		return count > 0, nil
	default:
		log.Print(err)
		return false, errors.New("Internal server error :^(")
	}
}

func ValidateUsername(username string) error {
	usernameLength := len(username)
	if usernameLength == 0 {
		return errors.New("Username cannot be 0 characters")
	}

	if usernameLength > 16 {
		return errors.New("Username cannot be longer than 16 characters")
	}

	return nil
}

func ValidatePassword(password string) error {
	passwordLength := len(password)
	if passwordLength < 5 {
		return errors.New("Password cannot be less than 5 characters")
	}

	if passwordLength > 64 {
		return errors.New("Password cannot be longer than 64 characters")
	}

	return nil
}

func HashPasswordBcrypt(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Print(err)
		return "", errors.New("Internal server error :^(")
	}
	return string(bytes), nil
}

func GetNewToken() (string, error) {
	cryptoBytes := make([]byte, 64)
	_, err := rand.Read(cryptoBytes)
	if err != nil {
		log.Print(err)
		return "", errors.New("Internal server error :^(")
	}

	token := base64.URLEncoding.EncodeToString(cryptoBytes)
	return token, nil
}

// returns token, error
func CreateUser(username string, password string) (string, error) {
	err := ValidateUsername(username)
	if err != nil {
		return "", err
	}

	err = ValidatePassword(password)
	if err != nil {
		return "", err
	}

	usernameInUse, err := UsernameInUse(username)
	if err != nil {
		return "", err
	}

	if usernameInUse == true {
		return "", errors.New("Username in use")
	}

	hashedPassword, err := HashPasswordBcrypt(password)
	if err != nil {
		return "", err
	}

	id := xid.New().String()
	token, err := GetNewToken()
	if err != nil {
		return "", err
	}

	sql := "INSERT INTO users (id, username, token, password) VALUES ($1, $2, $3, $4)"
	_, err = db.Exec(sql, id, username, token, hashedPassword)
	if err != nil {
		log.Print(err)
		return "", errors.New("Internal server error :^(")
	}

	return token, nil
}

func getUserByKeyValue(key string, value string) (User, error) { // stops code duplication for GetUserById and GetUserByUsername
	user := User{}
	row := db.QueryRow("SELECT id, username, token, role, password, verified FROM users WHERE "+key+" = $1", value)
	switch err := row.Scan(&user.Id, &user.Username, &user.Token, &user.Role, &user.HashedPassword, &user.Verified); err {
	case sql.ErrNoRows:
		return user, errors.New("User not found")
	case nil:
		return user, nil
	default:
		log.Print(err)
		return user, errors.New("Internal server error :^(")
	}
}

func GetUserById(id string) (User, error) {
	user, err := getUserByKeyValue("id", id)
	return user, err
}

func GetUserByUsername(username string) (User, error) {
	user, err := getUserByKeyValue("username", username)
	return user, err
}

func GetUserByToken(token string) (User, error) {
	user, err := getUserByKeyValue("token", token)
	return user, err
}
