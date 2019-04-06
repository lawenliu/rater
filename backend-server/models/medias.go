package models

import (
	"time"

	"database/sql"
	"rater/backend-server/models/mymysql"

	"github.com/go-sql-driver/mysql"
)

// Poster model definiton.
type Media struct {
	UserName     string    `json:"user_name,omitempty"`
	FilePath     string    `json:"file_path,omitempty"`
	Dtype        string    `json:"dtype,omitempty"`
	UploadDate   time.Time `json:"upload_date,omitempty"`
}

// NewUser alloc and initialize a user.
func NewMedia(userName, filePath, dtype string, t time.Time) (u *Media, err error) {
	media := Media{
		UserName:   userName,
		FilePath:   filePath,
		Dtype:      dtype,
		UploadDate: t,
	}

	return &media, nil
}

// Insert insert a document to collection.
func (u *Media) Insert() (code int, err error) {
	db := mymysql.Conn()
	//defer db.Close()

	st, err := db.Prepare("INSERT INTO medias(user_name, file_path, dtype, upload_date) VALUES(?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	if _, err := st.Exec(u.UserName, u.FilePath, u.Dtype, u.UploadDate); err != nil {
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

// FindByTitle query a document according to input id.
func (u *Media) FindByUser(userName string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT user_name, file_path, dtype, upload_date FROM medias WHERE user_id = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(userName)

	var tmpUserName sql.NullString
	var tmpFilePath sql.NullString
	var tmpDtype sql.NullString
	var tmpUploadDate mysql.NullTime
	if err := row.Scan(&tmpUserName, &tmpFilePath, &tmpDtype, &tmpUploadDate); err != nil {
		// Not found.
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		}

		return ErrDatabase, err
	}

	if tmpUserName.Valid {
		u.UserName = tmpUserName.String
	}
	if tmpFilePath.Valid {
		u.FilePath = tmpFilePath.String
	}
	if tmpDtype.Valid {
		u.Dtype = tmpDtype.String
	}
	if tmpUploadDate.Valid {
		u.UploadDate = tmpUploadDate.Time
	}

	return 0, nil
}
