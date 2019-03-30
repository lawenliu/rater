package models

import (
	"crypto/rand"
	"fmt"
	"io"
	"time"
	"errors"

	"database/sql"
	"rater/backend-server/models/mymysql"

	"golang.org/x/crypto/scrypt"
	"github.com/go-sql-driver/mysql"
)

// User model definiton.
type User struct {
	ID       int64    `json:"id,omitempty"`
	Phone    string    `json:"phone,omitempty"`
	Name     string    `json:"name,omitempty"`
	Password string    `json:"password,omitempty"`
	Salt     string    `json:"salt,omitempty"`
	RegDate  time.Time `json:"reg_date,omitempty"`
}

const pwHashBytes = 64

func generateSalt() (salt string, err error) {
	buf := make([]byte, pwHashBytes)
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", buf), nil
}

func generatePassHash(password string, salt string) (hash string, err error) {
	h, err := scrypt.Key([]byte(password), []byte(salt), 16384, 8, 1, pwHashBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h), nil
}

// NewUser alloc and initialize a user.
func NewUser(r *RegisterForm, t time.Time) (u *User, err error) {
	salt, err := generateSalt()
	if err != nil {
		return nil, err
	}
	hash, err := generatePassHash(r.Password, salt)
	if err != nil {
		return nil, err
	}

	user := User{
		Name:     r.Name,
		Password: hash,
		Salt:     salt,
		RegDate:  t,
	}

	return &user, nil
}

// Insert insert a document to collection.
func (u *User) Insert() (code int, err error) {
	db := mymysql.Conn()
	//defer db.Close()

	st, err := db.Prepare("INSERT INTO users(name, password, salt, reg_date) VALUES(?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	//if result, err := st.Exec(
	if _, err := st.Exec(u.Name, u.Password, u.Salt, u.RegDate); err != nil {
		if e, ok := err.(*mysql.MySQLError); ok {
			//Duplicate key
			if e.Number == 1062 {
				return ErrDupRows, err
			}

			return ErrDatabase, err
		}

		return ErrDatabase, err
	}

	//r.ID, _ = result.LastInsertId()

	return 0, nil
}

// FindByID query a document according to input id.
func (u *User) FindByID(id string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, phone, name, password, reg_date FROM users WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(id)

	var tmpID sql.NullInt64
	var tmpPhone sql.NullString
	var tmpName sql.NullString
	var tmpPassword sql.NullString
	var tmpRegDate mysql.NullTime
	if err := row.Scan(&tmpID, &tmpPhone, &tmpName, &tmpPassword, &tmpRegDate); err != nil {
		// Not found.
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		}

		return ErrDatabase, err
	}

	if tmpID.Valid {
		u.ID = tmpID.Int64
	}
	if tmpPhone.Valid {
		u.Phone = tmpPhone.String
	}
	if tmpName.Valid {
		u.Name = tmpName.String
	}
	if tmpPassword.Valid {
		u.Password = tmpPassword.String
	}
	if tmpRegDate.Valid {
		u.RegDate = tmpRegDate.Time
	}

	return 0, nil
}

// CheckPass compare input password.
func (u *User) CheckPass(pass string) (ok bool, err error) {
	hash, err := generatePassHash(pass, u.Salt)
	if err != nil {
		return false, err
	}

	return u.Password == hash, nil
}

// ClearPass clear password information.
func (u *User) ClearPass() {
	u.Password = ""
	u.Salt = ""
}

// ChangePass update password and salt information according to input id.
func ChangePass(id, oldPass, newPass string) (code int, err error) {
	db := mymysql.Conn()
	defer db.Close()

	st, err := db.Prepare("SELECT salt, password FROM users WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()
	row := st.QueryRow(id)
	var tmpSalt sql.NullString
	var tmpPassword sql.NullString
	if err := row.Scan(&tmpSalt, &tmpPassword); err != nil {
		// Not found.
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		}

		return ErrDatabase, err
	}
	if !tmpSalt.Valid {
		return -1, errors.New("can't user, please make sure you registered before.")
	}
	
	oldHash, err := generatePassHash(oldPass, tmpSalt.String)
	if err != nil {
		return ErrSystem, err
	}
	if tmpPassword.String != oldHash {
		return -1, errors.New("can't user, please make sure you registered before.")
	}
	newSalt, err := generateSalt()
	if err != nil {
		return ErrSystem, err
	}
	newHash, err := generatePassHash(newPass, newSalt)
	if err != nil {
		return ErrSystem, err
	}

	// update database
	st, err = db.Prepare("UPDATE users SET password = ?, salt = ? WHERE id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	result, err := st.Exec(newHash, newHash, id)
	if err != nil {
		return ErrDatabase, err
	}

	num, _ := result.RowsAffected()
	if num > 0 {
		return 0, nil
	}

	return ErrNotFound, nil
}

