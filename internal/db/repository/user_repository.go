package repository

import (
	"database/sql"
	"fmt"
	"go-auth/internal/db"
	"go-auth/internal/models"
)

type UserRepository struct {
	DbAdapter db.Database
}

func NewUserRepository(dbAdapter db.Database) *UserRepository {
	return &UserRepository{
		DbAdapter: dbAdapter,
	}
}

func (ur *UserRepository) FindByEmail(email string) []models.User {
	ur.DbAdapter.Connect()
	defer ur.DbAdapter.Close()
	rows, err := ur.DbAdapter.Query(getFindByEmailSql(), email)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer rows.Close()

	return mapRowsToUsers(rows)
}

func getFindByEmailSql() string {
	query := `SELECT * FROM users WHERE email = $1`
	return query
}

func mapRowsToUsers(rows *sql.Rows) []models.User {
	var users []models.User

	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.EncryptedPassword, &user.CreatedAt)
		if err != nil {
			fmt.Println("Erro ao escanear linha:", err)
		} else {
			users = append(users, user)
		}
	}

	return users
}
