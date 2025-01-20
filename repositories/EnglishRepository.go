package repositories

import (
	"context"
	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/model"
	"gorm.io/gorm"
)

type EnglishRepository struct {
	DB *gorm.DB
}

func (r *EnglishRepository) Create(ctx context.Context, name string, userId int) (*model.Topic, error) {
	topic := &model.Topic{Name: name, UserID: int32(userId)}
	if err := r.DB.WithContext(ctx).Create(topic).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *EnglishRepository) GetByUserID(ctx context.Context, userId int) ([]*model.Topic, error) {
	var topics []*model.Topic
	if err := r.DB.WithContext(ctx).Where("user_id = ?", userId).Find(&topics).Error; err != nil {
		return nil, err
	}
	return topics, nil
}
