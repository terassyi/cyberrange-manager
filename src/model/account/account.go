package account

import (
	"github.com/terassyi/cyberrange-manager/src/database"
	"github.com/terassyi/cyberrange-manager/src/util"
	"time"
)

type Account struct {
	Id       int
	Name     string
	Password string
	Os       string
	Address  string
	Expire   *time.Time
	Status   string
}

func New(name, password string) (*Account, error) {
	hash, err := util.CalcHash(password)
	if err != nil {
		return nil, err
	}
	return &Account{
		Id:       -1,
		Name:     name,
		Password: hash,
		Os:       "",
		Expire:   nil,
	}, nil
}

func GetAccounts(db *database.DB) ([]Account, error) {
	var accounts []Account
	data, err := db.GetAccounts()
	if err != nil {
		return nil, err
	}
	for _, a := range data {
		account := Account{
			Id:       a.Id,
			Name:     a.Name,
			Password: "",
			Os:       a.Os,
			Expire:   a.Expire,
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
