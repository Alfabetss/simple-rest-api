package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Alfabetss/simple-rest-api/entity"
	"github.com/Masterminds/squirrel"
)

// TalentRepository that related with talent repository
type TalentRepository interface {
	Create(ctx context.Context, talent entity.Talent) (int64, error)
	FindTalent(ctx context.Context, ID int64) (entity.Talent, error)
	Delete(ctx context.Context, ID int64) error
}

// TalentRepositoryImpl implementation interface
type TalentRepositoryImpl struct {
	db *sql.Tx
}

// NewTalentRepositoryImpl constructor
func NewTalentRepositoryImpl(db *sql.Tx) TalentRepository {
	return TalentRepositoryImpl{
		db: db,
	}
}

// Create function for insert to talent table
func (t TalentRepositoryImpl) Create(ctx context.Context, talent entity.Talent) (int64, error) {
	query, args, err := squirrel.Insert("talent").Columns(
		"name",
	).Values(talent.Name).ToSql()
	if err != nil {
		return 0, err
	}

	res, err := t.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	log.Printf("success insert talent : %s, with id %d", talent.Name, id)

	return id, nil
}

// FindTalent function to find talent by talent id
func (t TalentRepositoryImpl) FindTalent(ctx context.Context, ID int64) (entity.Talent, error) {
	var talent entity.Talent
	query, args, err := squirrel.Select("id", "name").
		From("talent").
		Where(squirrel.Eq{"id": ID}).ToSql()
	if err != nil {
		return talent, nil
	}

	row := t.db.QueryRow(query, args...)
	err = row.Scan(
		&talent.ID,
		&talent.Name,
	)
	if err != nil {
		return talent, err
	}

	return talent, nil
}

// Delete function for delete talent row by id
func (t TalentRepositoryImpl) Delete(ctx context.Context, ID int64) error {
	query, args, err := squirrel.Delete("talent").
		Where(squirrel.Eq{"id": ID}).ToSql()
	if err != nil {
		return err
	}

	_, err = t.db.Exec(query, args...)
	if err != nil {
		return err
	}

	log.Printf("success delete talent with id : %d", ID)
	return nil
}
