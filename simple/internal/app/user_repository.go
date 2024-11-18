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
		password TEXT NOT NULL
		)`)
	if err != nil {
		log.Fatalf("failed to create user table: %v", err.Error())
	}

	_, err = ur.db.Exec(`CREATE TABLE IF NOT EXISTS user_role(
		user_id INTEGER PRIMARY KEY,
		role_id INTEGER
		)`)
	if err != nil {
		log.Fatalf("failed to create user_role table: %v", err.Error())
	}
}

func (ur *UserRepository) GetOne(ctx context.Context, id int) (*User, error) {
	user := User{}
	row := ur.db.QueryRowContext(ctx, `SELECT * FROM users WHERE id = ?`, id)
	err := row.Scan(&user.Id, &user.Email, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (ur *UserRepository) GetMany(ctx context.Context) ([]*User, error) {
	rows, err := ur.db.QueryContext(ctx, `SELECT * FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (ur *UserRepository) Create(ctx context.Context, user *User) (*User, error) {
	res, err := ur.db.ExecContext(ctx,
		`INSERT INTO USERS (name, email, password) VALUES (?, ?, ?);`,
		user.Name,
		user.Email,
		user.Password,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	user.Id = int(id)

	if err := ur.syncRoles(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *UserRepository) Update(ctx context.Context, user *User) (*User, error) {
	_, err := ur.db.ExecContext(
		ctx,
		`UPDATE users SET name = ?, email = ? WHERE id = ?`,
		user.Name,
		user.Email,
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

func (ur *UserRepository) Roles(ctx context.Context, userId int) ([]*Role, error) {
	rows, err := ur.db.QueryContext(ctx, `SELECT FROM user_role WHERE user_id = ?`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	roles := []*Role{}
	for rows.Next() {
		role := &Role{}
		if err := rows.Scan(role.Id, role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	return roles, nil
}

func (ur *UserRepository) syncRoles(ctx context.Context, user *User) error {
	_, err := ur.db.ExecContext(ctx,
		`DELETE FROM user_role WHERE user_id = ?`,
		user.Id,
	)
	if err != nil {
		return err
	}

	for _, role := range user.Roles {
		_, err := ur.db.ExecContext(ctx,
			`INSERT INTO user_role (user_id, role_id) VALUES (?, ?);`,
			user.Id,
			role.Id,
		)
		if err != nil {
			return err
		}
	}

	return nil
}
