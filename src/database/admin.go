package database

type AdminAccount struct {
	Id       int
	Name     string
	Password string
}

func (db *DB) GetAdminAccountByName(name string) (*AdminAccount, error) {
	admin := &AdminAccount{}
	query := "SELECT * FROM admin_accounts WHERE name = ?"
	if err := db.Conn.QueryRow(query, name).Scan(&admin.Id, &admin.Name, &admin.Password); err != nil {
		return nil, err
	}
	return admin, nil
}

func (db *DB) CreateAdminAccount(name, password string) error {
	query, err := db.Conn.Prepare("INSERT INTO admin_accounts (name, password) VALUES (?, ?)")
	if err != nil {
		return err
	}
	if _, err = query.Exec(name, password); err != nil {
		return err
	}
	return nil
}
