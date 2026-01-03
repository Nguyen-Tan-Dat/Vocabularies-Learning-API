package repositories

import (
	"context"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/model"
	"gorm.io/gorm"
)

type TopicRepository struct {
	DB *gorm.DB
}

func (r *TopicRepository) Create(ctx context.Context, name string, userId int32) (*model.Topic, error) {
	topic := &model.Topic{
		Name:   name,
		OfUser: userId,
	}
	if err := r.DB.WithContext(ctx).Create(topic).Error; err != nil {
		return nil, err
	}

	// topic.ID đã được DB sinh
	return topic, nil
}
func (r *TopicRepository) Update(ctx context.Context, userId int32, input model.UpdateTopicInput) (*model.Topic, error) {
	topic := &model.Topic{
		ID:     input.ID,
		Name:   input.Name,
		OfUser: userId,
	}
	if err := r.DB.WithContext(ctx).Create(topic).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *TopicRepository) GetByUserID(ctx context.Context, userId int32) ([]*model.Topic, error) {
	var topics []*model.Topic
	if err := r.DB.WithContext(ctx).Where("of_user = ?", userId).Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func (r *TopicRepository) Search(ctx context.Context, userID int32, name string) ([]*model.Topic, error) {
	var topics []*model.Topic
	if err := r.DB.WithContext(ctx).
		Where(
			"of_user = ? AND name LIKE ?",
			userID,
			"%"+name+"%",
		).
		Find(&topics).
		Error; err != nil {
		return nil, err
	}
	return topics, nil
}

func (r *TopicRepository) Get(ctx context.Context, userID int32, id int32) (*model.Topic, error) {
	var topics *model.Topic
	if err := r.DB.WithContext(ctx).
		Where(
			"of_user = ? AND id = ?",
			userID,
			id,
		).
		Find(&topics).
		Error; err != nil {
		return nil, err
	}
	return topics, nil
}
