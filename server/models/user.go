package models

import (
	"context"
	"os"

	db "github.com/abhinavkale-dev/golang-auth/prisma/generated"
	"github.com/joho/godotenv"
)

type User struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var PrismaClient *db.PrismaClient

func InitPrisma() error {
	_ = godotenv.Load()

	if os.Getenv("DATABASE_URL") == "" {
		return os.ErrInvalid
	}

	PrismaClient = db.NewClient()
	return PrismaClient.Prisma.Connect()
}

func CreateUser(user *User) error {
	ctx := context.Background()
	_, err := PrismaClient.User.CreateOne(
		db.User.Email.Set(user.Email),
		db.User.Password.Set(user.Password),
	).Exec(ctx)
	return err
}

func GetUserByEmail(email string) (*User, error) {
	ctx := context.Background()
	u, err := PrismaClient.User.FindUnique(
		db.User.Email.Equals(email),
	).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}
