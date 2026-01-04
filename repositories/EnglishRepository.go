package repositories

import (
	"context"

	"github.com/Nguyen-Tan-Dat/Vocabularies-Learning-API/graph/model"
	"gorm.io/gorm"
)

type EnglishRepository struct {
	DB *gorm.DB
}

const table = "english"

func (r *EnglishRepository) GetAll(ctx context.Context) ([]*model.English, error) {
	var list []*model.English
	if err := r.DB.Table(table).
		WithContext(ctx).
		Find(&list).
		Error; err != nil {
		return nil, err
	}
	return list, nil
}
func (r *EnglishRepository) Search(ctx context.Context, word string) ([]*model.English, error) {
	var list []*model.English
	if err := r.DB.Table(table).
		WithContext(ctx).
		Where(
			"word LIKE ?",
			"%"+word+"%",
		).
		Find(&list).
		Error; err != nil {
		return nil, err
	}
	return list, nil
}
