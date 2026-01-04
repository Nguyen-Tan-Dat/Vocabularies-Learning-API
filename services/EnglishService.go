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
func (s *EnglishService) Search(ctx context.Context, word string) ([]*model.English, error) {
	return s.Repo.Search(ctx, word)
}
