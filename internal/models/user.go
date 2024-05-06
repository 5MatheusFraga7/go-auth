package models

import "time"

type User struct {
	Id                int
	Name              string
	Email             string
	EncryptedPassword string
	CreatedAt         time.Time
}
