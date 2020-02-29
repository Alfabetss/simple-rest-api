package service

import (
	"context"

	"github.com/Alfabetss/simple-rest-api/config"
	"github.com/Alfabetss/simple-rest-api/entity"
	"github.com/Alfabetss/simple-rest-api/repository"
)

// TalentService service to handle business logic for talent
type TalentService interface {
	CreateTalent(ctx context.Context) error
}

// TalentServiceImpl implementation
type TalentServiceImpl struct {
}

// NewTalentServiceImpl constructor
func NewTalentServiceImpl() TalentService {
	return TalentServiceImpl{}
}

// CreateTalent function for create talent data
func (t TalentServiceImpl) CreateTalent(ctx context.Context) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}

	// begin transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// insert into talent table
	talentRepo := repository.NewTalentRepositoryImpl(tx)
	talentID, err := talentRepo.Create(ctx, entity.Talent{Name: "robert"})
	if err != nil {
		return err
	}

	// insert into experience table
	expRepo := repository.NewExperienceRepositoryImpl(tx)
	_, err = expRepo.Create(ctx, entity.Experience{
		TalentID: talentID,
		Company:  "abece",
	})

	if err != nil {
		return err
	}

	err = tx.Commit()
	return err
}
