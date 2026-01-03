package services

import (
	"context"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/model"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/repositories"
)

type EnglishService struct {
	Repo repositories.EnglishRepository
}

func (s *EnglishService) CreateTopic(ctx context.Context, name string, userId int) (*model.Topic, error) {
	return s.Repo.Create(ctx, name, userId)
}

func (s *EnglishService) GetAll(ctx context.Context) ([]*model.English, error) {
	return s.Repo.GetAll(ctx)
}
