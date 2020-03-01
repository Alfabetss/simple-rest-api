package service

import (
	"context"
	"log"

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

// FindTalentResponse response for find talent
type FindTalentResponse struct {
	Talent     entity.Talent       `json:"talent"`
	Experience []entity.Experience `json:"experience"`
}

// TalentService service to handle business logic for talent
type TalentService interface {
	CreateTalent(ctx context.Context, req *CreateTalentRequest) error
	FindTalent(ctx context.Context, ID int64) (FindTalentResponse, error)
	Delete(ctx context.Context, ID int64) (err error)
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

// FindTalent function for find talent data
func (t TalentServiceImpl) FindTalent(ctx context.Context, ID int64) (response FindTalentResponse, err error) {
	db, err := config.Connect()
	if err != nil {
		return
	}

	// begin transaction
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		db.Close()
		tx.Rollback()
	}()

	talentRepository := repository.NewTalentRepositoryImpl(tx)
	talent, err := talentRepository.FindTalent(ctx, ID)
	if err != nil {
		return
	}

	experienceRepository := repository.NewExperienceRepositoryImpl(tx)
	experience, err := experienceRepository.FindExperience(ctx, ID)
	if err != nil {
		return
	}

	err = tx.Commit()
	return FindTalentResponse{
		Talent:     talent,
		Experience: experience,
	}, err
}

// Delete function to delete row in experience & talent table
func (t TalentServiceImpl) Delete(ctx context.Context, ID int64) (err error) {
	db, err := config.Connect()
	if err != nil {
		return
	}

	// begin transaction
	tx, err := db.Begin()
	if err != nil {
		return
	}
	defer func() {
		db.Close()
		tx.Rollback()
	}()

	expRepository := repository.NewExperienceRepositoryImpl(tx)
	err = expRepository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	talentRepo := repository.NewTalentRepositoryImpl(tx)
	err = talentRepo.Delete(ctx, ID)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("failed to commit error : %s", err.Error())
	}

	return err
}
