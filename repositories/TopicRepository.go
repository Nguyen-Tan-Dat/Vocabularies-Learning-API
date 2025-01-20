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
	topic := &model.Topic{Name: name, UserID: userId}
	if err := r.DB.WithContext(ctx).Create(topic).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *TopicRepository) GetByUserID(ctx context.Context, userId int32) ([]*model.Topic, error) {
	var topics []*model.Topic
	if err := r.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}
