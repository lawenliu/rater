package models

import (
	"time"

	"database/sql"
	"rater/backend-server/models/mymysql"

	"github.com/go-sql-driver/mysql"
)

// Poster model definiton.
type Poster struct {
	ID          int64     `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Type        string    `json:"type,omitempty"`
	Category    string    `json:"cate,omitempty"`
	Content     string    `json:"content,omitempty"`
	Dtype       string    `json:"dtype,omitempty"`
	ReferUrl    string    `json:"refer_url,omitempty"`
	UpdateDate  time.Time `json:"update_date,omitempty"`
	Status      string    `json:"status,omitempty"`
}

// NewPoster alloc and initialize a user.
func NewPoster(r *PostersForm, t time.Time, s string) (u *Poster, err error) {
	poster := Poster{
		Title:      r.Title,
		Type:       r.Type,
		Category:   r.Category,
		Content:    r.Content,
		Dtype:      r.Dtype,
		ReferUrl:   r.ReferUrl,
		UpdateDate: t,
		Status:     s,
	}

	return &poster, nil
}

// Insert insert a document to collection.
func (u *Poster) Insert() (code int, err error) {
	db := mymysql.Conn()
	//defer db.Close()

	st, err := db.Prepare("INSERT INTO posters(title, type, category, content, dtype, refer_url, update_date, status) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	if _, err := st.Exec(u.Title, u.Type, u.Category, u.Content, u.Dtype, u.ReferUrl, u.UpdateDate, u.Status); err != nil {
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
func (u *Poster) FindByTitle(title string) (code int, err error) {
	db := mymysql.Conn()

	st, err := db.Prepare("SELECT id, title, type, category, content, dtype, refer_url, update_date, status FROM posters WHERE title = ?")
	if err != nil {
		return ErrDatabase, err
	}
	defer st.Close()

	row := st.QueryRow(title)

	var tmpID sql.NullInt64
	var tmpTitle sql.NullString
	var tmpType sql.NullString
	var tmpCategory sql.NullString
	var tmpContent sql.NullString
	var tmpDtype sql.NullString
	var tmpReferUrl sql.NullString
	var tmpUpdateDate mysql.NullTime
	var tmpStatus sql.NullString

	if err := row.Scan(&tmpID, &tmpTitle, &tmpType, &tmpCategory, &tmpContent, &tmpDtype, &tmpReferUrl, &tmpUpdateDate, &tmpStatus); err != nil {
		// Not found.
		if err == sql.ErrNoRows {
			return ErrNotFound, err
		}

		return ErrDatabase, err
	}

	if tmpID.Valid {
		u.ID = tmpID.Int64
	}
	if tmpTitle.Valid {
		u.Title = tmpTitle.String
	}
	if tmpType.Valid {
		u.Type = tmpType.String
	}
	if tmpCategory.Valid {
		u.Category = tmpCategory.String
	}
	if tmpContent.Valid {
		u.Content = tmpContent.String
	}
	if tmpDtype.Valid {
		u.Dtype = tmpDtype.String
	}
	if tmpReferUrl.Valid {
		u.ReferUrl = tmpReferUrl.String
	}
	if tmpUpdateDate.Valid {
		u.UpdateDate = tmpUpdateDate.Time
	}
	if tmpStatus.Valid {
		u.Status = tmpStatus.String
	}
	return 0, nil
}
