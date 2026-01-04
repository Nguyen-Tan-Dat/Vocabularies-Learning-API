package services

import (
	"context"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/model"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/repositories"
)

type TopicService struct {
	Repo repositories.TopicRepository
}

func (s *TopicService) Create(ctx context.Context, input model.NewTopicInput, userId int32) (*model.Topic, error) {
	return s.Repo.Create(ctx, input.Name, userId)
}
func (s *TopicService) Search(ctx context.Context, userId int32, name string) ([]*model.Topic, error) {
	return s.Repo.Search(ctx, userId, name)
}

func (s *TopicService) Update(ctx context.Context, userId int32, input model.UpdateTopicInput) (*model.Topic, error) {
	return s.Repo.Update(ctx, userId, input)
}
func (s *TopicService) GetTopics(ctx context.Context, userId int32) ([]*model.Topic, error) {
	return s.Repo.GetByUserID(ctx, userId)
}

func (s *TopicService) Get(ctx context.Context, userID int32, id int32) (*model.Topic, error) {
	return s.Repo.Get(ctx, userID, id)
}
func (s *TopicService) Delete(ctx context.Context, userID int32, id int32) bool {
	return s.Repo.Delete(ctx, userID, id)
}
