package repository

import (
	"database/sql"
	"net/http"

	"github.com/vtdthang/goapi/entities"
	"github.com/vtdthang/goapi/lib/enums"
	httperror "github.com/vtdthang/goapi/lib/errors"
)

// IUserRepository represent user contract
type IUserRepository interface {
	FindByEmail(email string) (*entities.User, error)
	InsertOne(userEntity entities.User) error
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
	sqlStatement := `SELECT * FROM users WHERE email=$1;`
	var user entities.User
	row := db.DBCon.QueryRow(sqlStatement, email)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

	if err != nil {
		if err != sql.ErrNoRows {
			err := httperror.NewHTTPError(http.StatusInternalServerError, enums.ServerErrCode, enums.ServerErrMsg)
			return nil, err
		}

		//err := httperror.NewHTTPError(http.StatusNotFound, enums.UserNotFoundErrCode, enums.UserNotFoundErrMsg)
		return nil, nil
	}

	return &user, nil
}

// InsertOne insert new user
func (db *pgDB) InsertOne(userEntity entities.User) error {
	sqlStatement := `
	INSERT INTO users (id, email, first_name, last_name)
	VALUES ($1, $2, $3, $4)`
	_, err := db.DBCon.Exec(sqlStatement, userEntity.ID, userEntity.Email, userEntity.FirstName, userEntity.LastName)

	if err != nil {
		err = httperror.NewHTTPError(http.StatusInternalServerError, enums.ServerErrCode, enums.ServerErrMsg)
		return err
	}

	return nil
}
