package repo

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `json:"id" db:"id"`
	FirstName   string `json:"first_name" db:"first_name"`
	LastName    string `json:"last_name" db:"last_name"`
	Email       string `json:"email" db:"email"`
	Password    string `json:"password" db:"password"`
	IsShopOwner bool   `json:"is_shop_owner" db:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User,error)
	// List() ([]*User, error)
	// Delete(userID int) error
	// Update(user User) (*User, error)
}

type userRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}

func (u userRepo) Create(user User) (*User, error)  {
	query := `INSERT INTO users (first_name, last_name, email, password, is_shop_owner) 
	VALUES ($1, $2, $3, $4, $5) RETURNING id`

	var id int
	err := u.db.QueryRow(query, user.FirstName, user.LastName, user.Email, user.Password, user.IsShopOwner).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	user.ID = id
	return &user, nil
}

func (u userRepo) Find(email, pass string) (*User,error) {
	var user User
	fmt.Println("repo email:", email, " pass:", pass)
	query := `SELECT id, first_name, last_name, email, password, is_shop_owner FROM users WHERE email=$1 AND password=$2`
	err := u.db.Get(&user, query, email, pass)
	fmt.Println(err)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
