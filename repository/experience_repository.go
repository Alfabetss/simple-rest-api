package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Alfabetss/simple-rest-api/entity"
	"github.com/Masterminds/squirrel"
)

// ExperienceRepository interface that related with experience repository
type ExperienceRepository interface {
	Create(ctx context.Context, exp entity.Experience) (int64, error)
	FindExperience(ctx context.Context, talentID int64) error
}

// ExperienceRepositoryImpl implementation interface
type ExperienceRepositoryImpl struct {
	db *sql.Tx
}

// NewExperienceRepositoryImpl constructor
func NewExperienceRepositoryImpl(db *sql.Tx) ExperienceRepository {
	return ExperienceRepositoryImpl{
		db: db,
	}
}

// FindExperience function to find experience by talent id
func (e ExperienceRepositoryImpl) FindExperience(ctx context.Context, talentID int64) error {
	return nil
}

// Create function to insert to experience table
func (e ExperienceRepositoryImpl) Create(ctx context.Context, exp entity.Experience) (int64, error) {
	query, args, err := squirrel.Insert("experience").Columns("company", "talent_id").
		Values(exp.Company, exp.TalentID).ToSql()
	if err != nil {
		return 0, err
	}

	res, err := e.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Printf("success insert experience for talent id : %d, company name : %s", id, exp.Company)
	return id, nil
}
