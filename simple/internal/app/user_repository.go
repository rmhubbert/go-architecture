package app

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type UserRepository struct {
	db     *sql.DB
	dbPath string
}

func NewUserRepository(dbPath string) *UserRepository {
	ur := &UserRepository{
		db:     &sql.DB{},
		dbPath: dbPath,
	}

	ur.initDatabase()
	return ur
}

func (ur *UserRepository) initDatabase() {
	var err error
	ur.db, err = sql.Open("sqlite", ur.dbPath)
	if err != nil {
		log.Fatalf("failed to open db: %v", err.Error())
	}

	_, err = ur.db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL,
		password TEXT NOT NULL,
		role_id INTEGER NOT NULL
		)`)
	if err != nil {
		log.Fatalf("failed to create user table: %v", err.Error())
	}
}

func (ur *UserRepository) GetOne(ctx context.Context, id int) (*User, error) {
	user := User{
		Role: &Role{},
	}
	row := ur.db.QueryRowContext(
		ctx,
		`SELECT users.id, users.name, users.email, users.password, roles.id, roles.name FROM users LEFT JOIN roles ON users.role_id = roles.id WHERE users.id = ?`,
		id,
	)
	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role.Id,
		&user.Role.Name,
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (ur *UserRepository) GetMany(ctx context.Context) ([]*User, error) {
	rows, err := ur.db.QueryContext(
		ctx,
		`SELECT users.id, users.name, users.email, users.password, roles.id, roles.name FROM users LEFT JOIN roles ON users.role_id = roles.id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{
			Role: &Role{},
		}
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Role.Id, &user.Role.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) Create(ctx context.Context, user *User) (*User, error) {
	res, err := ur.db.ExecContext(ctx,
		`INSERT INTO USERS (name, email, password, role_id) VALUES (?, ?, ?, ?);`,
		user.Name,
		user.Email,
		user.Password,
		user.Role.Id,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.Id = int(id)

	return user, nil
}

func (ur *UserRepository) Update(ctx context.Context, user *User) (*User, error) {
	_, err := ur.db.ExecContext(
		ctx,
		`UPDATE users SET name = ?, email = ?, role_id = ? WHERE id = ?`,
		user.Name,
		user.Email,
		user.Role.Id,
		user.Id,
	)
	if err != nil {
		return nil, err
	}

	return ur.GetOne(ctx, user.Id)
}

func (ur *UserRepository) Delete(ctx context.Context, id int) error {
	_, err := ur.db.ExecContext(ctx, `DELETE FROM users WHERE id = ?`, id)
	return err
}

func (ur *UserRepository) UpdatePassword(ctx context.Context, user *User) (*User, error) {
	_, err := ur.db.ExecContext(
		ctx,
		`UPDATE users SET password = ? WHERE id = ?`,
		user.Password,
		user.Id,
	)
	if err != nil {
		return nil, err
	}

	return ur.GetOne(ctx, user.Id)
}
