package database

import "time"

type Account struct {
	Id       int
	Name     string
	Expire   time.Time
	Os       string
	Password string
}

func (db *DB) GetAccounts() ([]Account, error) {
	var res []Account
	query := "SELECT id, name, expire, os FROM accounts;"
	rows, err := db.Conn.Query(query)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		account := Account{}
		if err := rows.Scan(&account.Id, &account.Name, &account.Expire, &account.Os); err != nil {
			return nil, err
		}
		res = append(res, account)
	}
	return res, nil
}
