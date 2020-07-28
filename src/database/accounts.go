package database

import "time"

type Account struct {
	Id       int
	Name     string
	Expire   *time.Time
	Os       string
	Password string
}

func (db *DB) GetAccounts() ([]Account, error) {
	var res []Account
	query := "SELECT id, name, expire, os FROM accounts"
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

func (db *DB) CreateAccount(name, password, os string, expire *time.Time) error {
	query, err := db.Conn.Prepare("INSERT INTO accounts (name, password, os, expire) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	if _, err = query.Exec(name, password, os, expire); err != nil {
		return err
	}
	return nil
}

func (db *DB) DeleteAccounts(name string) error {
	return nil
}
