package services

import (
	"context"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/model"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/repositories"
)

type EnglishService struct {
	Repo repositories.EnglishRepository
}

func (s *EnglishService) GetAll(ctx context.Context) ([]*model.English, error) {
	return s.Repo.GetAll(ctx)
}

// READ - Get by ID
func (s *EnglishService) GetByID(
	ctx context.Context,
	id int,
) (*model.English, error) {

	return s.Repo.GetByID(ctx, id)
}

// CREATE
func (s *EnglishService) Create(
	ctx context.Context,
	input *model.NewEnglishInput,
) (*model.English, error) {

	return s.Repo.Create(
		ctx,
		input.Word,
		input.Phonetic,
		input.Audio,
	)
}

// UPDATE
func (s *EnglishService) Update(
	ctx context.Context,
	input *model.UpdateEnglishInput,
) (*model.English, error) {

	return s.Repo.Update(ctx, input)
}

// DELETE
func (s *EnglishService) Delete(
	ctx context.Context,
	id int,
) error {

	return s.Repo.Delete(ctx, id)
}
