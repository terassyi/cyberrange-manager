package account

import "time"

type Account struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Expire   time.Time `json:"expire"`
	Os       string    `json:"os"`
	Status   bool      `json:"status"`
}
