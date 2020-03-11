package repository

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/vtdthang/goapi/entities"
	"github.com/vtdthang/goapi/lib/enums"
	httperror "github.com/vtdthang/goapi/lib/errors"
	"github.com/vtdthang/goapi/lib/helpers"
)

// IUserRepository represent user contract
type IUserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	InsertOne(userEntity entities.User) error
	CreateUserAndAuthData(userEntity entities.User, authEntity entities.Auth) error
}

type pgDB struct {
	DBCon *sql.DB
}

//NewUserRepository will create an object which represent for IUserRepository
func NewUserRepository(db *sql.DB) IUserRepository {
	return &pgDB{DBCon: db}
}

// FindByEmail find a user by email
func (db *pgDB) FindByEmail(email string) (*entities.User, error) {
	defer helpers.TimeTrack(time.Now(), "FindByEmailRepository")

	sqlStatement := `
		SELECT id, first_name, last_name, email, password 
		FROM users 
		WHERE email=$1;
	`
	var pID string
	var pFirstName string
	var pLastName string
	var pEmail string
	var pPassword sql.NullString

	row := db.DBCon.QueryRow(sqlStatement, email)
	err := row.Scan(&pID, &pFirstName, &pLastName, &pEmail, &pPassword)

	if err != nil {
		fmt.Println("FindByEmail ", err)
		if err != sql.ErrNoRows {
			err := httperror.NewHTTPError(http.StatusInternalServerError, enums.ServerErrCode, enums.ServerErrMsg)
			return nil, err
		}

		//err := httperror.NewHTTPError(http.StatusNotFound, enums.UserNotFoundErrCode, enums.UserNotFoundErrMsg)
		return nil, nil
	}

	userEntity := &entities.User{
		ID:        pID,
		FirstName: pFirstName,
		LastName:  pLastName,
		Email:     pEmail,
		Password:  pPassword.String,
	}

	return userEntity, nil
}

// InsertOne insert new user
func (db *pgDB) InsertOne(userEntity entities.User) error {
	sqlStatement := `
	INSERT INTO users (id, email, first_name, last_name, password)
	VALUES ($1, $2, $3, $4, $5)`
	_, err := db.DBCon.Exec(sqlStatement, userEntity.ID, userEntity.Email, userEntity.FirstName, userEntity.LastName, userEntity.Password)

	if err != nil {
		err = httperror.NewHTTPError(http.StatusInternalServerError, enums.ServerErrCode, enums.ServerErrMsg)
		return err
	}

	return nil
}

func (db *pgDB) CreateUserAndAuthData(userEntity entities.User, authEntity entities.Auth) error {
	tx, err := db.DBCon.Begin()
	if err != nil {
		return err
	}

	{
		stmt, err := tx.Prepare(`INSERT INTO users (id, email, first_name, last_name, password)
		VALUES ($1, $2, $3, $4, $5);`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(userEntity.ID, userEntity.Email, userEntity.FirstName, userEntity.LastName, userEntity.Password); err != nil {
			tx.Rollback()
			return err
		}
	}

	{
		stmt, err := tx.Prepare(`INSERT INTO auths (id, user_id, refresh_token, expires_at)
                     VALUES($1, $2, $3, $4);`)
		if err != nil {
			tx.Rollback()
			return err
		}
		defer stmt.Close()

		if _, err := stmt.Exec(authEntity.ID, authEntity.UserID, authEntity.RefreshToken, authEntity.ExpiresAt); err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
