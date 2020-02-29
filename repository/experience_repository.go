package repository

import "context"

// ExperienceRepository interface that related with experience repository
type ExperienceRepository interface {
}

// ExperienceRepositoryImpl implementation interface
type ExperienceRepositoryImpl struct {
}

// NewExperienceRepositoryImpl constructor
func NewExperienceRepositoryImpl() ExperienceRepository {
	return ExperienceRepositoryImpl{}
}

// FindExperience function to find experience by talent id
func (e ExperienceRepositoryImpl) FindExperience(ctx context.Context, talentID int64) error {
	return nil
}
