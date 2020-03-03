package user

import (
	"database/sql"
	"fmt"

	"github.com/vtdthang/goapi/entities"
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
	return &pgDBConn{DBCon: db}
}

func (db *pgDB) FindByEmail(email string) (*entities.User, error) {
	sqlStatement := `SELECT * FROM users WHERE email=$1;`
	var user *entities.User
	row := db.DBCon.QueryRow(sqlStatement, email)
	err := row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)

	if err != nil {
		fmt.Println("No rows were returned!")
		return nil, err
	}

	fmt.Println(user)
	return &user, nil
}
