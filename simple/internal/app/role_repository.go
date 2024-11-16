package app

import (
	"context"
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

type RoleRepository struct {
	db     *sql.DB
	dbPath string
}

func NewRoleRepository(dbPath string) *RoleRepository {
	ur := &RoleRepository{
		db:     &sql.DB{},
		dbPath: dbPath,
	}

	ur.initDatabase()
	return ur
}

func (ur *RoleRepository) initDatabase() {
	var err error
	ur.db, err = sql.Open("sqlite", ur.dbPath)
	if err != nil {
		log.Fatalf("failed to open db: %v", err.Error())
	}

	_, err = ur.db.Exec(`CREATE TABLE IF NOT EXISTS roles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
		)`)
	if err != nil {
		log.Fatalf("failed to create roles table: %v", err.Error())
	}
}

func (ur *RoleRepository) GetOne(ctx context.Context, id int) (*Role, error) {
	role := Role{}
	row := ur.db.QueryRowContext(ctx, `SELECT * FROM roles WHERE id = ?`, id)
	err := row.Scan(&role.Id, &role.Name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (ur *RoleRepository) GetMany(ctx context.Context) ([]*Role, error) {
	rows, err := ur.db.QueryContext(ctx, `SELECT * FROM roles`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*Role
	for rows.Next() {
		role := &Role{}
		if err := rows.Scan(&role.Id, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (ur *RoleRepository) Create(ctx context.Context, role *Role) (*Role, error) {
	res, err := ur.db.ExecContext(ctx,
		`INSERT INTO USERS (name) VALUES (?);`,
		role.Name,
	)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	role.Id = int(id)
	return role, nil
}

func (ur *RoleRepository) Update(ctx context.Context, role *Role) (*Role, error) {
	_, err := ur.db.ExecContext(
		ctx,
		`UPDATE roles SET name = ? WHERE id = ?`,
		role.Name,
		role.Id,
	)
	if err != nil {
		return nil, err
	}

	return ur.GetOne(ctx, role.Id)
}

func (ur *RoleRepository) Delete(ctx context.Context, id int) error {
	_, err := ur.db.ExecContext(ctx, `DELETE FROM roles WHERE id = ?`, id)
	return err
}
