package service

import (
	"context"

	"github.com/Alfabetss/simple-rest-api/config"
	"github.com/Alfabetss/simple-rest-api/entity"
	"github.com/Alfabetss/simple-rest-api/repository"
)

// CreateTalentRequest request for create talent
type CreateTalentRequest struct {
	Name      string              `json:"name"`
	Companies []ExperienceRequest `json:"experience"`
}

// ExperienceRequest experience request
type ExperienceRequest struct {
	CompanyName string `json:"companyName"`
}

// TalentService service to handle business logic for talent
type TalentService interface {
	CreateTalent(ctx context.Context, req *CreateTalentRequest) error
}

// TalentServiceImpl implementation
type TalentServiceImpl struct {
}

// NewTalentServiceImpl constructor
func NewTalentServiceImpl() TalentService {
	return TalentServiceImpl{}
}

// CreateTalent function for create talent data
func (t TalentServiceImpl) CreateTalent(ctx context.Context, req *CreateTalentRequest) error {
	db, err := config.Connect()
	if err != nil {
		return err
	}

	// begin transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		db.Close()
		tx.Rollback()
	}()

	// insert into talent table
	talentRepo := repository.NewTalentRepositoryImpl(tx)
	talentID, err := talentRepo.Create(ctx, entity.Talent{Name: req.Name})
	if err != nil {
		return err
	}

	// insert into experience table
	expRepo := repository.NewExperienceRepositoryImpl(tx)
	for _, exp := range req.Companies {
		_, err = expRepo.Create(ctx, entity.Experience{
			TalentID: talentID,
			Company:  exp.CompanyName,
		})

		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	return err
}
