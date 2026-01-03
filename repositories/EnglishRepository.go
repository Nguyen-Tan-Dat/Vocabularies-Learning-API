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
	topic := &model.Topic{Name: name, OfUser: int32(userId)}
	if err := r.DB.WithContext(ctx).Create(topic).Error; err != nil {
		return nil, err
	}
	return topic, nil
}

func (r *EnglishRepository) GetAll(ctx context.Context) ([]*model.English, error) {
	var model []*model.English
	if err := r.DB.Table("english").WithContext(ctx).Find(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}
