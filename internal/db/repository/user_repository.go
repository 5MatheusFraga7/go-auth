package repository

import (
	"database/sql"
	"fmt"
	"go-auth/internal/db"
	"go-auth/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	DbAdapter db.Database
}

// Construtor do repositório de usuários, recebendo qual banco deve conectar

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

func (ur *UserRepository) Create(user models.UserParams) []models.User {
	ur.DbAdapter.Connect()
	defer ur.DbAdapter.Close()
	_, err := ur.DbAdapter.Exec(getCreateUserSql(),
		user.Name,
		user.Email,
		hashPassword(user.Password),
	)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return nil
}

func getFindByEmailSql() string {
	query := `SELECT * FROM users WHERE email = $1`
	return query
}

func getCreateUserSql() string {
	query := `INSERT INTO users(name, email, encrypted_password, created_at) VALUES ($1, $2, $3, now())`
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

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(hashedPassword)
}
