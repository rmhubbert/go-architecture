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

func (rr *RoleRepository) initDatabase() {
	var err error
	rr.db, err = sql.Open("sqlite", rr.dbPath)
	if err != nil {
		log.Fatalf("failed to open db: %v", err.Error())
	}

	_, err = rr.db.Exec(`CREATE TABLE IF NOT EXISTS roles(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
		)`)
	if err != nil {
		log.Fatalf("failed to create roles table: %v", err.Error())
	}
}

func (rr *RoleRepository) GetOne(ctx context.Context, id int) (*Role, error) {
	role := Role{}
	row := rr.db.QueryRowContext(ctx, `SELECT * FROM roles WHERE id = ?`, id)
	err := row.Scan(&role.Id, &role.Name)
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (rr *RoleRepository) GetMany(ctx context.Context) ([]*Role, error) {
	rows, err := rr.db.QueryContext(ctx, `SELECT * FROM roles`)
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

func (rr *RoleRepository) GetManyById(ctx context.Context, ids []int) ([]*Role, error) {
	var roles []*Role
	if len(ids) == 0 {
		return roles, nil
	}

	cids := make([]any, len(ids))
	for i, v := range ids {
		cids[i] = v
	}

	query := "SELECT * FROM roles WHERE"
	for i := 0; i < len(ids); i++ {
		if i > 0 {
			query += " OR"
		}
		query += " id = ?"
	}

	rows, err := rr.db.QueryContext(ctx, query, cids...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		role := &Role{}
		if err := rows.Scan(&role.Id, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (rr *RoleRepository) Create(ctx context.Context, role *Role) (*Role, error) {
	res, err := rr.db.ExecContext(ctx,
		`INSERT INTO roles (name) VALUES (?);`,
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

func (rr *RoleRepository) Update(ctx context.Context, role *Role) (*Role, error) {
	_, err := rr.db.ExecContext(
		ctx,
		`UPDATE roles SET name = ? WHERE id = ?`,
		role.Name,
		role.Id,
	)
	if err != nil {
		return nil, err
	}

	return rr.GetOne(ctx, role.Id)
}

func (rr *RoleRepository) Delete(ctx context.Context, id int) error {
	_, err := rr.db.ExecContext(ctx, `DELETE FROM roles WHERE id = ?`, id)
	return err
}
