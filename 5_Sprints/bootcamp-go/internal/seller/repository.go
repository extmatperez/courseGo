package seller

import (
	"context"
	"database/sql"
	"errors"

	"github.com/extmatperez/meli_bootcamp10_sprints/internal/domain"
)

// Repository encapsulates the storage of a Seller.
type Repository interface {
	GetAll(ctx context.Context) ([]domain.Seller, error)
	GetByCid(ctx context.Context, cid int) (domain.Seller, error)
	GetById(ctx context.Context, id int) (domain.Seller, error)
	ExistsCid(ctx context.Context, cid int) bool
	ExistsId(ctx context.Context, id int) bool
	Save(ctx context.Context, s domain.SellerToSave) (int, error)
	Update(ctx context.Context, s domain.Seller) error
	Delete(ctx context.Context, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Seller, error) {

	rows, err := r.db.Query(`SELECT * FROM "main"."sellers"`)
	if err != nil {
		return nil, err
	}

	var sellers []domain.Seller

	for rows.Next() {
		s := domain.Seller{}
		_ = rows.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
		sellers = append(sellers, s)
	}

	return sellers, nil
}

/// Busca en la base de datos el seller por cid
func (r *repository) GetByCid(ctx context.Context, cid int) (domain.Seller, error) {

	sqlStatement := `SELECT * FROM "main"."sellers" WHERE cid=$1;`
	row := r.db.QueryRow(sqlStatement, cid)
	s := domain.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
	if err != nil {
		return domain.Seller{}, err
	}

	return s, nil
}

/// Busca en la base de datos el seller por id
func (r *repository) GetById(ctx context.Context, id int) (domain.Seller, error) {

	sqlStatement := `SELECT * FROM "main"."sellers" WHERE id=$1;`
	row := r.db.QueryRow(sqlStatement, id)
	s := domain.Seller{}
	err := row.Scan(&s.ID, &s.CID, &s.CompanyName, &s.Address, &s.Telephone, &s.LocalityID)
	if err != nil {
		return domain.Seller{}, err
	}

	return s, nil
}

/// Busca si existe en la base de datos el cid
func (r *repository) ExistsCid(ctx context.Context, cid int) bool {
	sqlStatement := `SELECT cid FROM "main"."sellers" WHERE cid=$1;`
	row := r.db.QueryRow(sqlStatement, cid)
	err := row.Scan(&cid)
	if err != nil {
		return false
	}
	return true
}

/// Busca si existe en la base de datos el id
func (r *repository) ExistsId(ctx context.Context, id int) bool {
	sqlStatement := `SELECT id FROM "main"."sellers" WHERE id=$1;`
	row := r.db.QueryRow(sqlStatement, id)
	err := row.Scan(&id)
	if err != nil {
		return false
	}
	return true
}

func (r *repository) Save(ctx context.Context, s domain.SellerToSave) (int, error) {

	stmt, err := r.db.Prepare(`INSERT INTO "main"."sellers"("cid","company_name","address","telephone","locality_id") VALUES (?,?,?,?,?)`)
	if err != nil {
		return 0, err
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, s domain.Seller) error {
	stmt, err := r.db.Prepare(`UPDATE "main"."sellers" SET "cid"=?, "company_name"=?, "address"=?, "telephone"=?, "locality_id"=?  WHERE "id"=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(s.CID, s.CompanyName, s.Address, s.Telephone, s.LocalityID, s.ID)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affect < 1 {
		return errors.New("seller not found")
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stmt, err := r.db.Prepare(`DELETE FROM "main"."sellers" WHERE id=?`)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	affect, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affect < 1 {
		return errors.New("seller not found")
	}

	return nil
}
