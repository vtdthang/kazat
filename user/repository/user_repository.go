package repository

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/vtdthang/goapi/entities"
	"github.com/vtdthang/goapi/lib/enums"
	httperror "github.com/vtdthang/goapi/lib/errors"
)

// IUserRepository represent user contract
type IUserRepository interface {
	FindByEmail(email string) (*entities.User, error)
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
		err := httperror.NewHTTPError(http.StatusInternalServerError, enums.ServerErrCode, enums.ServerErrMsg)

		fmt.Println("REPO ", err)
		return nil, err
	}

	return &user, nil
}
